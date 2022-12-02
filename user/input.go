package user

type RegisterUserInput struct {
	UserName    string `json:"user_name" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password" binding:"required"`
	PhoneNumber string `json:"phone_number"`
}

type DataSession struct {
	Username  string `json:"username"`
	UniqueID  string `json:"uniqueId"`
	SessionID string `json:"sessionId"`
	OtpToken  string `json:"otpToken"`
}

type LoginInput struct {
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required"`
}

type CheckEmailInput struct {
	Email string `json:"email" binding:"required,email"`
}

type PayloadOTP struct {
	Username      string `json:"username" binding:"required"`
	InstitutionID int    `json:"institution_id" binding:"required"`
	OTP           int    `json:"otp" binding:"required"`
}

type FormCreateUserInput struct {
	Name       string `form:"name" binding:"required"`
	Email      string `form:"email" binding:"required,email"`
	Occupation string `form:"occupation" binding:"required"`
	Password   string `form:"password" binding:"required"`
	Error      error
}

type FormUpdateUserInput struct {
	UUID       string
	UserName   string `form:"user_name" binding:"required"`
	Name       string `form:"name" binding:"required"`
	Email      string `form:"email" binding:"required,email"`
	Occupation string `form:"occupation" binding:"required"`
	Error      error
}

type AccessToken struct {
	UserID        string `json:"user_id"`
	UserName      string `json:"user_name"`
	UserEmail     string `json:"user_email"`
	InstitutionID int    `json:"institution_id"`
	AccessToken   string `json:"access_token"`
}

type RequestAPIV1AUTH struct {
	InstitutionId int    `json:"institution_id"`
	Username      string `json:"username"`
	Password      string `json:"password"`
}

type RequestGopayOTP struct {
	OTP      int    `json:"otp"`
	Username string `json:"username"`
}
