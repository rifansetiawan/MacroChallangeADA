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
