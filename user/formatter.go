package user

type UserFormatter struct {
	UUID     string `json:"uuid"`
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	Token    string `json:"token"`
	ImageURL string `json:"image_url"`
}

func FormatUser(user User, token string) UserFormatter {
	formatter := UserFormatter{
		UUID:     user.UUID,
		UserName: user.UserName,
		Email:    user.Email,
		Token:    token,
		ImageURL: user.AvatarFileName,
	}

	return formatter
}

type UserFormatterDeviceToken struct {
	UUID           string `json:"uuid"`
	UserName       string `json:"user_name"`
	Email          string `json:"email"`
	RegistrationID string `json:"registration_id"`
}

func FormatUserDeviceToken(user User) UserFormatterDeviceToken {
	formatter := UserFormatterDeviceToken{
		UUID:           user.UUID,
		UserName:       user.UserName,
		Email:          user.Email,
		RegistrationID: user.RegistrationId,
	}

	return formatter
}
