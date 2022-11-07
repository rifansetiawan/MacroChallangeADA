package user

import "time"

type User struct {
	UUID           string
	FirstName      string
	LastName       string
	Email          string
	PasswordHash   string
	AvatarFileName string
	Role           string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
