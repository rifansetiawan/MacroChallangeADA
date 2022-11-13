package rest

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
)

func GetHttp(url string, header map[string]string) map[string]interface{} {
	options := Options{
		Method:      "GET",
		URL:         url,
		Headers:     header,
		Timeout:     30 * time.Second,
		ContentType: "application/json",
	}

	response := GET(&options)
	result := make(map[string]interface{})
	json.Unmarshal(response.Body, &result)

	return result
}

func PostHttp(url string, header map[string]string, body map[string]interface{}) map[string]interface{} {
	b, err := json.Marshal(body)
	if err != nil {
		fmt.Println("error marshal json")
		fmt.Println(err)
	}

	options := Options{
		Method:      "POST",
		URL:         url,
		Headers:     header,
		Body:        b,
		Timeout:     30 * time.Second,
		ContentType: "application/json",
	}

	logrus.Debug("making POST request")
	logrus.Debug(options)
	response := POST(&options)
	logrus.Debug("response")
	logrus.Debug(response)
	shortenLink := make(map[string]interface{})
	json.Unmarshal(response.Body, &shortenLink)

	return shortenLink
}
