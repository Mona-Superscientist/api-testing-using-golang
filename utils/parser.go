package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func GetRequestBody(body interface{}) io.Reader {
	requestBody, _ := json.Marshal(body)
	return bytes.NewBuffer(requestBody)
}

func GetResponseBody(response *http.Response) interface{} {
	resp, _ := io.ReadAll(response.Body)
	var responseBody interface{}
	json.Unmarshal(resp, &responseBody)
	return responseBody
}
