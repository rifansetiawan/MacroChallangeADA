package handler

import (
	"kaia/auth"
	"kaia/user"
)

type brickAuthHandler struct {
	userService user.Service
	authService auth.Service
}

func NewBrickAuthHandler(userService user.Service, authService auth.Service) *userHandler {
	return &userHandler{userService, authService}
}

// func (h *userHandler) AuthToken(c *gin.Context) {

// 	// headers := map[string]string{
// 	// 	"Authorization": "Basic ZjAzMzg1NzItZjQ2NC00NjQ2LTk1MjktNWMyZWE4MDA1MTRhOjJGVHd5N1FIdHNiRTZiRzdETnVjSk9TSnBrWTBuMw==",
// 	// }

// 	// options := rest.Options{
// 	// 	Method:      "GET",
// 	// 	URL:         "https://api.onebrick.io/v1/auth/token",
// 	// 	Headers:     headers,
// 	// 	Timeout:     30 * time.Second,
// 	// 	ContentType: "application/json",
// 	// }

// 	// response := rest.GET(&options)
// 	// var result []map[string]string
// 	// json.Unmarshal(response.Body, &result)
// 	client := http.Client{}
// 	req, err := http.NewRequest("GET", "https://api.onebrick.io/v1/auth/token", nil)
// 	if err != nil {
// 		//Handle Error
// 	}

// 	req.Header = http.Header{
// 		"Host":          {"www.host.com"},
// 		"Content-Type":  {"application/json"},
// 		"Authorization": {"Bearer ZjAzMzg1NzItZjQ2NC00NjQ2LTk1MjktNWMyZWE4MDA1MTRhOjJGVHd5N1FIdHNiRTZiRzdETnVjSk9TSnBrWTBuMw=="},
// 	}

// 	res, err := client.Do(req)
// 	c.JSON(http.StatusOK, res)

// }
