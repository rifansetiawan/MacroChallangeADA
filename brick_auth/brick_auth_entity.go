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
