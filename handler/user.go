package handler

import (
	"encoding/json"
	"fmt"
	"kaia/auth"
	"kaia/helper"
	"kaia/user"

	// "kaia/user"
	"net/http"

	// user "kaia/user"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
	authService auth.Service
}

func NewUserHandler(userService user.Service, authService auth.Service) *userHandler {
	return &userHandler{userService, authService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Register account failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := h.userService.RegisterUser(input)

	if err != nil {
		response := helper.APIResponse("Register account failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	token, err := h.authService.GenerateToken(newUser.UUID)
	if err != nil {
		response := helper.APIResponse("Register account failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(newUser, token)

	response := helper.APIResponse("Account has been registered", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) Login(c *gin.Context) {
	var input user.LoginInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loggedinUser, err := h.userService.Login(input)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		errorMessageString := errorMessage["errors"].(string)

		response := helper.APIResponseArray("Login failed", http.StatusUnprocessableEntity, "error", errorMessageString)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	token, err := h.authService.GenerateToken(loggedinUser.UUID)
	if err != nil {
		response := helper.APIResponse("Login failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(loggedinUser, token)

	response := helper.APIResponse("Successfuly loggedin", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)

}

func (h *userHandler) CheckEmailAvailability(c *gin.Context) {
	var input user.CheckEmailInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Email checking failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	isEmailAvailable, err := h.userService.IsEmailAvailable(input)
	if err != nil {
		errorMessage := gin.H{"errors": "Server error"}
		response := helper.APIResponse("Email checking failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data := gin.H{
		"is_available": isEmailAvailable,
	}

	metaMessage := "Email has been registered"

	if isEmailAvailable {
		metaMessage = "Email is available"
	}

	response := helper.APIResponse(metaMessage, http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) UploadAvatar(c *gin.Context) {
	file, err := c.FormFile("avatar")
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload avatar image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	userID := currentUser.UUID

	path := fmt.Sprintf("images/%d-%s", userID, file.Filename)

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload avatar image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.userService.SaveAvatar(userID, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload avatar image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{"is_uploaded": true}
	response := helper.APIResponse("Avatar successfuly uploaded", http.StatusOK, "success", data)

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) FetchUser(c *gin.Context) {

	currentUser := c.MustGet("currentUser").(user.User)

	formatter := user.FormatUser(currentUser, "")

	response := helper.APIResponse("Successfuly fetch user data", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)

}

type DataData struct {
	Data struct {
		AccessToken string
	}
}

func (h *userHandler) AuthToken(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(user.User)
	fmt.Println(currentUser)
	if currentUser.UserName == "rifanganteng" {
		var respGetInstList interface{}

		// publicToken := d.Data.AccessToken
		// fmt.Println(publicToken)

		clientGetInstList := http.Client{}
		reqGetInstList, err := http.NewRequest("GET", "https://api.onebrick.io/v1/institution/list", nil)
		if err != nil {
			//Handle Error
		}

		reqGetInstList.Header = http.Header{
			"Content-Type":  {"application/json"},
			"Authorization": {"Bearer public-production-ad98df55-fa5a-4664-8049-a5bfe4224887 "},
		}

		resGetInstList, err := clientGetInstList.Do(reqGetInstList)

		defer resGetInstList.Body.Close()

		err = json.NewDecoder(resGetInstList.Body).Decode(&respGetInstList)

		c.JSON(200, respGetInstList)
	} else {
		var respGetInstList interface{}

		// publicToken := d.Data.AccessToken
		// fmt.Println(publicToken)

		clientGetInstList := http.Client{}
		reqGetInstList, err := http.NewRequest("GET", "https://api.onebrick.io/v1/institution/list", nil)
		if err != nil {
			//Handle Error
		}

		reqGetInstList.Header = http.Header{
			"Content-Type":  {"application/json"},
			"Authorization": {"Bearer public-production-ad98df55-fa5a-4664-8049-a5bfe4224887 "},
		}

		resGetInstList, err := clientGetInstList.Do(reqGetInstList)

		defer resGetInstList.Body.Close()

		err = json.NewDecoder(resGetInstList.Body).Decode(&respGetInstList)

		c.JSON(200, respGetInstList)
	}
	// var responsetoreturn [string]interface{}
	// client := http.Client{}
	// req, err := http.NewRequest("GET", "https://api.onebrick.io/v1/auth/token", nil)
	// if err != nil {
	// 	//Handle Error
	// }

	// req.Header = http.Header{
	// 	"Content-Type":  {"application/json"},
	// 	"Authorization": {"Bearer ZjAzMzg1NzItZjQ2NC00NjQ2LTk1MjktNWMyZWE4MDA1MTRhOjJGVHd5N1FIdHNiRTZiRzdETnVjSk9TSnBrWTBuMw=="},
	// }

	// res, err := client.Do(req)

	// defer res.Body.Close()

	// err = json.NewDecoder(res.Body).Decode(&responsetoreturn)
	// body, _ := ioutil.ReadAll(res.Body)
	// d := DataData{}

	// json.Unmarshal([]byte(responsetoreturn), &d)

	// fmt.Println(d)
	// //get institution list

}

type RequestAPIV1AUTH struct {
	InstitutionId int    `json:"institution_id"`
	Username      string `json:"username"`
	Password      int    `json:"password"`
}

func (h *userHandler) AuthTokenToAccessToken(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(user.User)

	var input user.RequestAPIV1AUTH

	err := c.ShouldBindJSON(&input)
	if err != nil {
		return
	}
	response, err := h.userService.AuthTokenToAccessToken(input, currentUser)
	if err != nil {
		response := helper.APIResponse("Auth Bank Error Occured", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	c.JSON(200, response)
}

type DataSession struct {
	Data Data `json:"data"`
}
type Data struct {
	Username  string `json:"username"`
	UniqueID  string `json:"uniqueId"`
	SessionID string `json:"sessionId"`
	OtpToken  string `json:"otpToken"`
}

func (h *userHandler) AuthTokenToAccessTokenGopay(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(user.User)
	var input user.RequestAPIV1AUTH

	err := c.ShouldBindJSON(&input)
	if err != nil {
		return
	}
	response, err := h.userService.AuthTokenToAccessTokenGopay(input, currentUser)
	if err != nil {
		response := helper.APIResponse("Auth Bank Error Occured", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	returnResponse := helper.APIResponse("AuthToken 1st Step Succeded", http.StatusOK, "success", response)

	c.JSON(200, returnResponse)

}

func (h *userHandler) OTPSessionToToken(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(user.User)
	fmt.Println(currentUser)

	var input user.PayloadOTP

	c.ShouldBindJSON(&input)

	response, err := h.userService.OTPSessionToToken(input, currentUser)
	if err != nil {
		response := helper.APIResponse("Auth Bank Error Occured", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	c.JSON(200, response)

}

func (h *userHandler) AccountListTransactions(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(user.User)
	fmt.Println(currentUser)
	var startDateString string
	var endDateString string

	startDate := c.Param("start")
	endDate := c.Param("end")

	startDateString = startDate
	endDateString = endDate

	accessTokens, err := h.userService.GetAccessTokensPerUser(currentUser)

	mergedAccountListAndTransactions, err := h.userService.GetAccountListTransactions(currentUser, accessTokens, startDateString, endDateString)

	fmt.Println(accessTokens)

	if err != nil {
		response := helper.APIResponse("Get Account List Transactions Error Occured", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	c.JSON(200, mergedAccountListAndTransactions)

}

// func (h *userHandler) AuthTokenToAccessToken(c *gin.Context) {
// 	currentUser := c.MustGet("currentUser").(user.User)
// 	fmt.Println(currentUser)
// 	var brickAuthResponse brickAuthEntity.BrickAuthResponse
// 	var dataToAccessTokenToBeSaved user.AccessToken
// 	var requestHandler user.RequestAPIV1AUTH
// 	if err := c.BindJSON(&requestHandler); err != nil {
// 		panic(err)
// 	}
// 	if currentUser.UserName == "rifanganteng" {
// 		// var respGetInstList interface{}

// 		// publicToken := d.Data.AccessToken
// 		// fmt.Println(publicToken)
// 		// var jsonStr = []byte(`
// 		// 	{
// 		// 		"institution_id" : 2,
// 		// 		"username" : "AMMADRIF0122",
// 		// 		"password": 252354
// 		// 	}
// 		// `)

// 		clientGetInstList := http.Client{}
// 		reqGetAccessToken, err := http.NewRequest("POST", "https://api.onebrick.io/v1/auth", c.Request.Body)
// 		if err != nil {
// 			//Handle Error
// 		}

// 		reqGetAccessToken.Header = http.Header{
// 			"Content-Type":  {"application/json"},
// 			"Authorization": {"Bearer public-production-ad98df55-fa5a-4664-8049-a5bfe4224887"},
// 		}

// 		resGetAccessToken, err := clientGetInstList.Do(reqGetAccessToken)
// 		if err != nil {
// 			fmt.Println(err.Error())
// 			return
// 		}

// 		fmt.Println(resGetAccessToken)

// 		defer resGetAccessToken.Body.Close()

// 		var testingDebug interface{}
// 		json.NewDecoder(resGetAccessToken.Body).Decode(&testingDebug)
// 		// if err != nil {
// 		// 	panic(err)
// 		// 	c.JSON(500, brickAuthResponse)
// 		// }
// 		fmt.Println(testingDebug)
// 		fmt.Println(resGetAccessToken.StatusCode)
// 		dataToAccessTokenToBeSaved.AccessToken = brickAuthResponse.Data
// 		dataToAccessTokenToBeSaved.InstitutionID = requestHandler.InstitutionId
// 		dataToAccessTokenToBeSaved.UserEmail = currentUser.Email
// 		dataToAccessTokenToBeSaved.UserID = currentUser.UUID
// 		fmt.Println("this is dataToAccessTokenToBeSaved  : ", dataToAccessTokenToBeSaved)
// 		// if resGetAccessToken.StatusCode == 200 {

// 		// }

// 		c.JSON(200, brickAuthResponse)
// 	}

// }

//backup

// func (h *userHandler) AuthTokenToAccessTokenGopay(c *gin.Context) {
// 	currentUser := c.MustGet("currentUser").(user.User)
// 	fmt.Println(currentUser)

// 	if currentUser.UserName == "rifanganteng" {
// 		var respDataSession DataSession
// 		var respDataSessionError interface{}

// 		clientGetInstList := http.Client{}
// 		reqGetInstList, err := http.NewRequest("POST", "https://api.onebrick.io/v1/auth", c.Request.Body)
// 		if err != nil {
// 			//Handle Error
// 		}

// 		reqGetInstList.Header = http.Header{
// 			"Content-Type":  {"application/json"},
// 			"Authorization": {"Bearer public-production-ad98df55-fa5a-4664-8049-a5bfe4224887"},
// 		}

// 		resGetInstList, err := clientGetInstList.Do(reqGetInstList)

// 		defer resGetInstList.Body.Close()

// 		if resGetInstList.StatusCode == 500 {
// 			err = json.NewDecoder(resGetInstList.Body).Decode(&respDataSessionError)
// 		}
// 		err = json.NewDecoder(resGetInstList.Body).Decode(&respDataSession)

// 		fmt.Println("testing : ", resGetInstList.StatusCode)
// 		fmt.Println("testing 2 : ", respDataSessionError)

// 		sessionToken, err := h.userService.SaveGopayData(user.DataSession(respDataSession.Data))
// 		fmt.Println(sessionToken)
// 		fmt.Println(respDataSession)
// 		if resGetInstList.StatusCode == 500 {
// 			c.JSON(200, respDataSessionError)
// 		} else {
// 			c.JSON(200, respDataSession)
// 		}

// 	}

// }

// BACKUP :

// func (h *userHandler) OTPSessionToToken(c *gin.Context) {
// 	currentUser := c.MustGet("currentUser").(user.User)
// 	fmt.Println(currentUser)

// 	var input user.PayloadOTP

// 	c.ShouldBindJSON(&input)

// 	if currentUser.UserName == "rifanganteng" {
// 		var respDataSessionError interface{}
// 		var respGetInstList interface{}

// 		sessionToken, _ := h.userService.FindOtpSession(input)
// 		fmt.Println(sessionToken)
// 		var sessPayload user.SessionPayload

// 		sessPayload.Username = sessionToken.Username
// 		sessPayload.UniqueID = sessionToken.UniqueID
// 		sessPayload.SessionID = sessionToken.SessionID
// 		sessPayload.OtpToken = sessionToken.OtpToken
// 		sessPayload.Otp = input.OTP
// 		fmt.Println(sessPayload)
// 		// data := url.Values{}
// 		// data.Set("username", sessionToken.Username)
// 		// data.Set("uniqueId", sessionToken.UniqueID)
// 		// data.Set("sessionId", sessionToken.SessionID)
// 		// data.Set("otpToken", sessionToken.OtpToken)
// 		// data.Set("otp", input.OTP)
// 		personJSON, err := json.Marshal(sessPayload)

// 		clientGetInstList := http.Client{}
// 		reqGetInstList, err := http.NewRequest("POST", "https://api.onebrick.io/v1/auth/gopay", bytes.NewBuffer(personJSON))
// 		if err != nil {
// 			//Handle Error
// 		}

// 		reqGetInstList.Header = http.Header{
// 			"Content-Type":  {"application/json"},
// 			"Authorization": {"Bearer public-production-ad98df55-fa5a-4664-8049-a5bfe4224887"},
// 		}
// 		resGetInstList, err := clientGetInstList.Do(reqGetInstList)
// 		if resGetInstList.StatusCode != 200 {
// 			err = json.NewDecoder(resGetInstList.Body).Decode(&respDataSessionError)
// 		} else {
// 			err = json.NewDecoder(resGetInstList.Body).Decode(&respGetInstList)
// 		}
// 		fmt.Println(resGetInstList)

// 		defer resGetInstList.Body.Close()
// 		if resGetInstList.StatusCode != 200 {
// 			c.JSON(400, respDataSessionError)
// 		} else {
// 			c.JSON(200, respGetInstList)
// 		}

// 	}

// }
