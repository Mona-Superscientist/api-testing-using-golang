package endpoints

import (
	"api-testing-example/controller"
	"api-testing-example/model"
	"api-testing-example/utils"
	"encoding/json"
	"io"
	"net/http"
)

func CreateToken(username, password string) *http.Response {
	credentials := model.AuthenticationInput{
		Username: username,
		Password: password,
	}

	response := controller.MakePostRequest("/auth", utils.GetRequestBody(credentials))
	return response
}

func GetAuthToken(username, password string) string {
	response := CreateToken(username, password)
	resp, _ := io.ReadAll(response.Body)
	var responseBody model.AuthenticationToken
	json.Unmarshal(resp, &responseBody)
	return responseBody.Token
}
