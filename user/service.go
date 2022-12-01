package user

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	brickAuthEntity "kaia/brick_auth"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
	Login(input LoginInput) (User, error)
	IsEmailAvailable(input CheckEmailInput) (bool, error)
	SaveAvatar(UUID string, fileLocation string) (User, error)
	GetUserByUUID(UUID string) (User, error)
	GetAllUsers() ([]User, error)
	UpdateUser(input FormUpdateUserInput) (User, error)
	SaveGopayData(session DataSession) (Session, error)
	FindOtpSession(input PayloadOTP) (Session, error)
	AuthTokenToAccessToken(input RequestAPIV1AUTH, currentUser User) (interface{}, error)
	AuthTokenToAccessTokenGopay(input RequestAPIV1AUTH, currentUser User) (brickAuthEntity.BrickAuthResponseGopay, error)
	OTPSessionToToken(input PayloadOTP, currentUser User) (brickAuthEntity.BrickAuthResponse, error)
	// GetAccountListTransactions(currentUser User) ([]string, error)
	GetAccessTokensPerUser(currentUser User) ([]AccessToken, error)
	GetAccountListTransactions(currentUser User, accessTokens []AccessToken, startDate string, endDate string) (interface{}, error)

	// SaveAccessToken(userDetailandAccessToekn AccessToken) (AccessToken, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) AuthTokenToAccessToken(input RequestAPIV1AUTH, currentUser User) (interface{}, error) {
	fmt.Println("this is request : ", input)
	var respDataSessionError interface{}
	var dataToAccessTokenToBeSaved AccessToken
	var brickAuthResponse brickAuthEntity.BrickAuthResponse

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(input)
	clientGetInstList := http.Client{}
	reqGetInstList, err := http.NewRequest("POST", "https://api.onebrick.io/v1/auth", b)
	if err != nil {
		//Handle Error
	}

	reqGetInstList.Header = http.Header{
		"Content-Type":  {"application/json"},
		"Authorization": {"Bearer public-production-ad98df55-fa5a-4664-8049-a5bfe4224887"},
	}

	resGetInstList, err := clientGetInstList.Do(reqGetInstList)

	defer resGetInstList.Body.Close()

	if resGetInstList.StatusCode != 200 {
		err = json.NewDecoder(resGetInstList.Body).Decode(&respDataSessionError)
		return respDataSessionError, nil
	} else {
		err = json.NewDecoder(resGetInstList.Body).Decode(&brickAuthResponse)

		dataToAccessTokenToBeSaved.AccessToken = brickAuthResponse.Data
		dataToAccessTokenToBeSaved.InstitutionID = input.InstitutionId
		dataToAccessTokenToBeSaved.UserEmail = currentUser.Email
		dataToAccessTokenToBeSaved.UserName = currentUser.UserName
		dataToAccessTokenToBeSaved.UserID = currentUser.UUID
		s.repository.SaveAccessToken(dataToAccessTokenToBeSaved)

		fmt.Println("this is dataToAccessTokenToBeSaved : ", dataToAccessTokenToBeSaved)

	}
	fmt.Println("ini return : ", brickAuthResponse)
	return brickAuthResponse, nil
}

func (s *service) AuthTokenToAccessTokenGopay(input RequestAPIV1AUTH, currentUser User) (brickAuthEntity.BrickAuthResponseGopay, error) {
	fmt.Println("this is request : ", input)
	var respDataSessionError interface{}
	var respDataSession brickAuthEntity.BrickAuthResponseGopay

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(input)
	clientGetInstList := http.Client{}
	reqGetInstList, err := http.NewRequest("POST", "https://api.onebrick.io/v1/auth", b)
	if err != nil {
		//Handle Error
	}

	reqGetInstList.Header = http.Header{
		"Content-Type":  {"application/json"},
		"Authorization": {"Bearer public-production-ad98df55-fa5a-4664-8049-a5bfe4224887"},
	}

	resGetInstList, err := clientGetInstList.Do(reqGetInstList)

	// defer resGetInstList.Body.Close()

	// if resGetInstList.StatusCode != 200 {
	// 	err = json.NewDecoder(resGetInstList.Body).Decode(&respDataSessionError)
	// 	return respDataSessionError, nil
	// }
	err = json.NewDecoder(resGetInstList.Body).Decode(&respDataSession)

	fmt.Println("testing : ", resGetInstList.StatusCode)
	fmt.Println("testing 2 : ", respDataSessionError)
	fmt.Println(" testing 2 : ", respDataSession)
	sessions := Session{}
	sessions.Username = respDataSession.Data.Username
	sessions.UniqueID = respDataSession.Data.UniqueID
	sessions.OtpToken = respDataSession.Data.OtpToken
	sessions.SessionID = respDataSession.Data.SessionID
	s.repository.DeletePrevSession(sessions.Username)
	s.repository.SaveSession(sessions)

	return respDataSession, nil
}

func (s *service) OTPSessionToToken(input PayloadOTP, currentUser User) (brickAuthEntity.BrickAuthResponse, error) {
	// var respDataSessionError interface{}
	var brickAuthResponse brickAuthEntity.BrickAuthResponse
	var dataToAccessTokenToBeSaved AccessToken
	phoneNumber := input.Username
	otpInput := input.OTP
	fmt.Println("this is phone number : ", phoneNumber)
	sessionByNumberPhone, err := s.repository.FindOtpNumber(phoneNumber)

	var sessPayload SessionPayload

	sessPayload.Username = sessionByNumberPhone.Username
	sessPayload.UniqueID = sessionByNumberPhone.UniqueID
	sessPayload.SessionID = sessionByNumberPhone.SessionID
	sessPayload.OtpToken = sessionByNumberPhone.OtpToken
	sessPayload.Otp = otpInput
	fmt.Println(sessPayload)

	personJSON, err := json.Marshal(sessPayload)

	clientGetInstList := http.Client{}
	reqGetInstList, err := http.NewRequest("POST", "https://api.onebrick.io/v1/auth/gopay", bytes.NewBuffer(personJSON))
	if err != nil {
		//Handle Error
	}

	reqGetInstList.Header = http.Header{
		"Content-Type":  {"application/json"},
		"Authorization": {"Bearer public-production-ad98df55-fa5a-4664-8049-a5bfe4224887"},
	}
	resGetInstList, err := clientGetInstList.Do(reqGetInstList)
	// if resGetInstList.StatusCode != 200 {
	// 	err = json.NewDecoder(resGetInstList.Body).Decode(&respDataSessionError)
	// } else {
	// 	err = json.NewDecoder(resGetInstList.Body).Decode(&brickAuthResponse)
	// }
	err = json.NewDecoder(resGetInstList.Body).Decode(&brickAuthResponse)
	dataToAccessTokenToBeSaved.AccessToken = brickAuthResponse.Data
	dataToAccessTokenToBeSaved.InstitutionID = input.InstitutionID
	dataToAccessTokenToBeSaved.UserEmail = currentUser.Email
	dataToAccessTokenToBeSaved.UserName = currentUser.UserName
	dataToAccessTokenToBeSaved.UserID = currentUser.UUID
	s.repository.SaveAccessToken(dataToAccessTokenToBeSaved)

	fmt.Println("this is dataToAccessTokenToBeSaved : ", dataToAccessTokenToBeSaved)

	// if resGetInstList.StatusCode != 200 {
	// 	return respDataSessionError, nil
	// }
	fmt.Println("ini return : ", brickAuthResponse)
	return brickAuthResponse, nil
}

func (s *service) FindOtpSession(input PayloadOTP) (Session, error) {
	phoneNumber := input.Username
	fmt.Println("this is phone number : ", phoneNumber)
	sessionByNumberPhone, err := s.repository.FindOtpNumber(phoneNumber)
	if err != nil {
		return sessionByNumberPhone, err
	}

	return sessionByNumberPhone, nil
}

func (s *service) SaveGopayData(session DataSession) (Session, error) {
	fmt.Println("session in service : ", session)
	sessions := Session{}
	sessions.Username = session.Username
	sessions.UniqueID = session.UniqueID
	sessions.OtpToken = session.OtpToken
	sessions.SessionID = session.SessionID

	newSession, err := s.repository.SaveSession(sessions)
	if err != nil {
		return newSession, err
	}

	return newSession, nil
}

// func (s *service) SaveAccessToken(userDetailandAccessToekn AccessToken) (AccessToken, error) {
// 	fmt.Println("session in service : ", userDetailandAccessToekn)
// 	userDetailandAccessToekn := AccessToken{}
// 	userDetailandAccessToekn.AccessToken = session.Username
// 	sessions.UniqueID = session.UniqueID
// 	sessions.OtpToken = session.OtpToken
// 	sessions.SessionID = session.SessionID

// 	newSession, err := s.repository.SaveSession(sessions)
// 	if err != nil {
// 		return newSession, err
// 	}

// 	return newSession, nil
// }

func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
	user := User{}
	user.UserName = input.UserName
	user.Email = input.Email

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}

	user.PasswordHash = string(passwordHash)
	user.Role = "user"
	user.UUID = uuid.New().String()

	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (s *service) Login(input LoginInput) (User, error) {
	email := input.Email
	password := input.Password

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}

	if user.UUID == "" {
		return user, errors.New("No user found on that email")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *service) IsEmailAvailable(input CheckEmailInput) (bool, error) {
	email := input.Email

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return false, err
	}

	if user.UUID == "" {
		return true, nil
	}

	return false, nil
}

func (s *service) SaveAvatar(UUID string, fileLocation string) (User, error) {
	user, err := s.repository.FindByUUID(UUID)
	if err != nil {
		return user, err
	}

	user.AvatarFileName = fileLocation

	updatedUser, err := s.repository.Update(user)
	if err != nil {
		return updatedUser, err
	}

	return updatedUser, nil
}

func (s *service) GetUserByUUID(UUID string) (User, error) {
	user, err := s.repository.FindByUUID(UUID)
	if err != nil {
		return user, err
	}

	if user.UUID == "" {
		return user, errors.New("No user found on with that ID")
	}

	return user, nil
}

func (s *service) GetAllUsers() ([]User, error) {
	users, err := s.repository.FindAll()
	if err != nil {
		return users, err
	}

	return users, nil
}

func (s *service) UpdateUser(input FormUpdateUserInput) (User, error) {
	user, err := s.repository.FindByUUID(input.UUID)
	if err != nil {
		return user, err
	}

	user.UserName = input.UserName
	user.Email = input.Email

	updatedUser, err := s.repository.Update(user)
	if err != nil {
		return updatedUser, err
	}

	return updatedUser, nil
}

func (s *service) GetAccessTokensPerUser(currentUser User) ([]AccessToken, error) {
	accessTokens, err := s.repository.GetAccessTokensPerUser(currentUser)
	if err != nil {
		return accessTokens, err
	}
	return accessTokens, nil
}

func (s *service) GetAccountListTransactions(currentUser User, accessTokens []AccessToken, startDate string, endDate string) (interface{}, error) {
	var arrayAccuntListTransactionSuccess []AccountListTransactionResponsePerAccessToken
	var arrayErrorResponseAccountListTransaction []ErrorResponseAccountListTransaction
	var mergedAccountListTransaction MergedAccountListTransaction
	for i, oneAccessToken := range accessTokens {
		fmt.Println("access token ke : ", i)
		fmt.Println("access token ", oneAccessToken)

		var responseAccountListTransactionInterface AccountListTransactionResponsePerAccessToken
		var responseErrorAccountListTransactionInterface ErrorResponseAccountListTransaction

		clientAccountListTransactions := http.Client{}
		reqAccountListTransaction, err := http.NewRequest("GET", "https://api.onebrick.io/v1/account/list", nil)
		if err != nil {
			//Handle Error
		}

		reqAccountListTransaction.Header = http.Header{
			"Content-Type":  {"application/json"},
			"Authorization": {"Bearer " + oneAccessToken.AccessToken},
		}

		responseAccountListTransaction, err := clientAccountListTransactions.Do(reqAccountListTransaction)

		defer responseAccountListTransaction.Body.Close()

		// if responseAccountListTransaction.StatusCode = 200 {
		// 	err = json.NewDecoder(resGetInstList.Body).Decode(&respDataSessionError)
		// 	return respDataSessionError, nil
		// }
		if responseAccountListTransaction.StatusCode == 200 {
			json.NewDecoder(responseAccountListTransaction.Body).Decode(&responseAccountListTransactionInterface)
		} else {
			json.NewDecoder(responseAccountListTransaction.Body).Decode(&responseErrorAccountListTransactionInterface)
		}
		// return respDataSessionError, nil
		// fmt.Println("this is response accountlist transactions : ", responseAccountListTransactionInterface)
		if responseAccountListTransactionInterface.Status == 200 {
			for _, responseAccountLooping := range responseAccountListTransactionInterface.Data {
				if responseAccountLooping.Type != "PayLater" {
					responseAccountLooping.InstitutionID = oneAccessToken.InstitutionID

					transactions, _ := s.GetTransactionsOnly(oneAccessToken, startDate, endDate)
					fmt.Println("ini transactionsssss : ", transactions.Data)
					responseAccountLooping.Transactions = transactions.Data
					mergedAccountListTransaction.Data = append(mergedAccountListTransaction.Data, responseAccountLooping)

				}
			}

			arrayAccuntListTransactionSuccess = append(arrayAccuntListTransactionSuccess, responseAccountListTransactionInterface)
			fmt.Println("status 200 : ", responseAccountListTransactionInterface.Data)

		} else {
			arrayErrorResponseAccountListTransaction = append(arrayErrorResponseAccountListTransaction, responseErrorAccountListTransactionInterface)
			responseErrorAccountListTransactionInterface.InstitutionID = oneAccessToken.InstitutionID
			mergedAccountListTransaction.DataError = append(mergedAccountListTransaction.DataError, responseErrorAccountListTransactionInterface)
		}
	}
	mergedAccountListTransaction.Message = "RifanGanteng OK"
	mergedAccountListTransaction.Status = 200
	var jsonData, _ = json.Marshal(mergedAccountListTransaction)
	var jsonString = string(jsonData)
	fmt.Println("this is Merged response accountlist transactions : ", string(jsonString))

	return mergedAccountListTransaction, nil
}

func (s *service) GetTransactionsOnly(oneAccessToken AccessToken, startDate string, endDate string) (ListTransactionFromTo, error) {
	var responseListTransactionMother ListTransactionFromTo

	clientListTransactionsMother := http.Client{}
	requestTransactionMother, err := http.NewRequest("GET", fmt.Sprintf("https://api.onebrick.io/v1/transaction/list?from=%s&to=%s", startDate, endDate), nil)
	if err != nil {
		//Handle Error
	}

	requestTransactionMother.Header = http.Header{
		"Content-Type":  {"application/json"},
		"Authorization": {"Bearer " + oneAccessToken.AccessToken},
	}

	responseListTransactionM, err := clientListTransactionsMother.Do(requestTransactionMother)

	defer responseListTransactionM.Body.Close()

	// if responseAccountListTransaction.StatusCode = 200 {
	// 	err = json.NewDecoder(resGetInstList.Body).Decode(&respDataSessionError)
	// 	return respDataSessionError, nil
	// }
	json.NewDecoder(responseListTransactionM.Body).Decode(&responseListTransactionMother)

	responseListTransactionMother.Message = "RifanGanteng List Transactions From To OK"
	responseListTransactionMother.Status = 200
	var jsonData, _ = json.Marshal(responseListTransactionMother)
	var jsonString = string(jsonData)
	fmt.Println("this is List Transaction Mother : ", string(jsonString))

	return responseListTransactionMother, nil
}
