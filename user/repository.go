package user

import "gorm.io/gorm"

type Repository interface {
	Save(user User) (User, error)
	FindByEmail(email string) (User, error)
	FindByUUID(UUID string) (User, error)
	Update(user User) (User, error)
	FindAll() ([]User, error)
	SaveSession(session Session) (Session, error)
	FindOtpNumber(phoneNumber string) (Session, error)
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
	err := r.db.Save(&user).Error

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
