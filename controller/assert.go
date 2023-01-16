package controller

import (
	"api-testing-example/utils"
	"fmt"
	"net/http"

	. "github.com/onsi/gomega"
)

func assertOnStatusCode(response *http.Response, statusCode int) {
	Expect(response.StatusCode).Should(Equal(statusCode))
	Expect(response.Status).Should(Equal(fmt.Sprintf("%d", statusCode) + " " + http.StatusText(statusCode)))
}

func Assert_StatusOK(response *http.Response) {
	assertOnStatusCode(response, http.StatusOK)
}

func Assert_StatusBadRequest(response *http.Response) {
	assertOnStatusCode(response, http.StatusBadRequest)
}

func Assert_StatusCreated(response *http.Response) {
	assertOnStatusCode(response, http.StatusCreated)
}

func Assert_ResponseBody_Contain_Key(response *http.Response, key string) {
	responseBody := utils.GetResponseBody(response)
	Expect(responseBody).Should(HaveKey(key))
}

func Assert_ResponseBody_Has_Value(response *http.Response, key string, value string) {
	responseBody := utils.GetResponseBody(response)
	Expect(responseBody).Should(HaveKeyWithValue(key, value))
}
