package brick_auth

type Response struct {
	Status  int
	Message string
	Data    DataResponse
}

type DataResponse struct {
	AccessToken  string
	PrimaryColor string
}

type BrickAuthResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

type BrickAuthResponseError struct {
	Status       int    `json:"status"`
	Message      string `json:"message"`
	Data         string `json:"data"`
	ErrorCode    string `json:"error_code"`
	ErrorMessage string `json:"error_string"`
}

type BrickAuthResponseGopay struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    DataSession `json:"data"`
}

type DataSession struct {
	Username  string `json:"username"`
	UniqueID  string `json:"uniqueId"`
	SessionID string `json:"sessionId"`
	OtpToken  string `json:"otpToken"`
}
