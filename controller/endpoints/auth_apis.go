package endpoints

import (
	"api-testing-example/controller"
	"api-testing-example/model"
	"api-testing-example/utils"
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
