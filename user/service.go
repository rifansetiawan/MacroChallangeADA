package user

import (
	"errors"
	"fmt"

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
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
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
