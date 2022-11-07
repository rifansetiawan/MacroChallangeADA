package user

import "time"

type User struct {
	UUID           string
	UserName       string
	Email          string
	PasswordHash   string
	AvatarFileName string
	Role           string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
