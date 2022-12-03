package user

import "gorm.io/gorm"

type Repository interface {
	Save(user User) (User, error)
	FindByEmail(email string) (User, error)
	FindByUUID(UUID string) (User, error)
	Update(user User) (User, error)
	UpdateRegistId(user User) (User, error)
	FindAll() ([]User, error)
	SaveSession(session Session) (Session, error)
	DeletePrevSession(phoneNumber string)
	FindOtpNumber(phoneNumber string) (Session, error)
	SaveAccessToken(access_tokens AccessToken) (AccessToken, error)
	GetAccessTokensPerUser(currentUser User) ([]AccessToken, error)
	DeleteExistingAccessTokensPerUser(currentUser User, institutionId int) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(user User) (User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) SaveSession(session Session) (Session, error) {
	err := r.db.Create(&session).Error
	if err != nil {
		return session, err
	}

	return session, nil
}
func (r *repository) DeletePrevSession(phoneNumber string) {
	r.db.Delete(&Session{}, "username LIKE ?", "%"+phoneNumber+"%")
}

func (r *repository) SaveAccessToken(access_tokens AccessToken) (AccessToken, error) {
	err := r.db.Create(&access_tokens).Error
	if err != nil {
		return access_tokens, err
	}

	return access_tokens, nil
}

func (r *repository) FindOtpNumber(phoneNumber string) (Session, error) {
	var session Session

	err := r.db.Where("username = ?", phoneNumber).Find(&session).Error
	if err != nil {
		return session, err
	}

	return session, nil
}

func (r *repository) FindByEmail(email string) (User, error) {
	var user User

	err := r.db.Where("email = ?", email).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindByUUID(UUID string) (User, error) {
	var user User

	err := r.db.Where("uuid = ?", UUID).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) Update(user User) (User, error) {
	// var user User
	err := r.db.Save(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) UpdateRegistId(user User) (User, error) {
	err := r.db.Exec("UPDATE users SET registration_id = ? where uuid = ?", user.RegistrationId, user.UUID).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindAll() ([]User, error) {
	var users []User

	err := r.db.Find(&users).Error
	if err != nil {
		return users, err
	}

	return users, nil
}
func (r *repository) GetAccessTokensPerUser(currenUser User) ([]AccessToken, error) {
	var access_tokens []AccessToken

	err := r.db.Where("user_email = ?", currenUser.Email).Find(&access_tokens).Error
	if err != nil {
		return access_tokens, err
	}

	return access_tokens, nil
}

func (r *repository) DeleteExistingAccessTokensPerUser(currenUser User, institutionID int) error {
	var access_token AccessToken
	r.db.Where("user_email = ? AND institution_id = ?", currenUser.Email, institutionID).Delete(&access_token)
	return nil
}
