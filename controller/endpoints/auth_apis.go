package endpoints

import (
	"api-testing-example/constants"
	"api-testing-example/controller"
	"api-testing-example/model"
	"api-testing-example/utils"
	"encoding/json"
	"io"
	"net/http"
)

// Login using /auth endpoint and return response
func CreateToken(username, password string) *http.Response {
	credentials := model.AuthenticationInput{
		Username: username,
		Password: password,
	}

	response := controller.MakePostRequest(
		constants.AuthURL, "", utils.GetRequestBody(credentials),
	)
	return response
}

// Extract token value from response body of /auth endpoint to be passed to
// other requests
func GetAuthToken(username, password string) string {
	response := CreateToken(username, password)
	resp, _ := io.ReadAll(response.Body)
	var responseBody model.AuthenticationToken
	json.Unmarshal(resp, &responseBody)
	return responseBody.Token
}
