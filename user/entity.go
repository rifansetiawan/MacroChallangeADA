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

type Session struct {
	Username  string
	UniqueID  string
	SessionID string
	OtpToken  string
}

type SessionPayload struct {
	Username  string `json:"username"`
	UniqueID  string `json:"uniqueId"`
	SessionID string `json:"sessionId"`
	OtpToken  string `json:"otpToken"`
	Otp       int    `json:"otp"`
}
