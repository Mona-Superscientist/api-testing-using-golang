package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func GetRequestBody(body interface{}) io.Reader {
	requestBody, _ := json.Marshal(body)
	return bytes.NewBuffer(requestBody)
}

func GetResponseBody(response *http.Response) interface{} {
	resp, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalln("A Fatal error occurred", err)
	}
	var responseBody interface{}
	error := json.Unmarshal(resp, &responseBody)
	if err != nil {
		fmt.Println("error:", error)
	}
	return responseBody
}
