package helper

import (
	"github.com/go-playground/validator/v10"
)

type Response struct {
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Data    interface{} `json:"data"`
}

type ResponseCustom struct {
	Message string      `json:"message"`
	Status  int         `json:"status"`
	Data    interface{} `json:"data"`
}

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

func APIResponse(message string, code int, status string, data interface{}) Response {
	// meta := Meta{
	// 	Message: message,
	// 	Code:    code,
	// 	Status:  status,
	// 	Data:    data,
	// }

	jsonResponse := Response{
		// Meta: meta,
		// Data: data,
		Message: message,
		Code:    code,
		Status:  status,
		Data:    data,
	}

	return jsonResponse
}

func APIResponseCustom(message string, status int, data interface{}) ResponseCustom {
	// meta := Meta{
	// 	Message: message,
	// 	Code:    code,
	// 	Status:  status,
	// 	Data:    data,
	// }

	jsonResponse := ResponseCustom{
		// Meta: meta,
		// Data: data,
		Message: message,
		Status:  status,
		Data:    data,
	}

	return jsonResponse
}

func FormatValidationError(err error) []string {
	var errors []string

	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}

	return errors
}

func APIResponseArray(message string, code int, status string, data string) Response {
	// meta := Meta{
	// 	Message: message,
	// 	Code:    code,
	// 	Status:  status,
	// 	Data:    data,
	// }
	var errorsArray []string
	// var dataErrorArray

	errorsArray = append(errorsArray, data)
	myMap := map[string][]string{
		"errors": errorsArray,
	}

	jsonResponse := Response{
		// Meta: meta,
		// Data: data,
		Message: message,
		Code:    code,
		Status:  status,
		Data:    myMap,
	}

	return jsonResponse
}
