package controller

import (
	"io"
	"log"
	"net/http"
	"os"
)

func makeRequest(method, url string, body io.Reader) *http.Request {
	request, _ := http.NewRequest(method, os.Getenv("BASE_URL")+url, body)
	request.Header.Set("content-type", "application/json")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Connection", "keep-alive")
	return request
}

func makeRequestWithAuthToken(method, url string, body io.Reader, token string) *http.Request {
	request := makeRequest(method, url, body)
	request.Header.Set("Cookie", "token="+token)
	return request
}

func MakeGetRequest(url string) *http.Response {
	request := makeRequest("GET", url, nil)

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Fatalln("An error occurred while executing request", err)
	}

	return response
}

func MakePostRequest(url string, body io.Reader) *http.Response {
	request := makeRequest("POST", url, body)

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Fatalln("An error occurred while executing request", err)
	}

	return response
}

func MakePostRequestWithAuthToken(url, token string, body io.Reader) *http.Response {
	request := makeRequestWithAuthToken("POST", url, body, token)

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Fatalln("An error occurred while executing request", err)
	}

	return response
}
