package controller

import (
	"api-testing-example/utils"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func makeRequest(method, path, token string, body io.Reader) *http.Request {
	utils.LoadEnv()
	request, _ := http.NewRequest(method, os.Getenv("BASE_URL")+path, body)
	request.Header.Set("content-type", "application/json")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Connection", "keep-alive")

	if token != "" {
		request.Header.Set("Cookie", "token="+token)
	}

	return request
}

func MakeGetRequest(path string, token string) *http.Response {
	request := makeRequest("GET", path, token, nil)

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Fatalln("An error occurred while executing request", err)
	}

	return response
}

func MakePostRequest(path, token string, body io.Reader) *http.Response {
	fmt.Println("test")
	fmt.Println(path)
	request := makeRequest("POST", path, token, body)

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Fatalln("An error occurred while executing request", err)
	}

	return response
}

func MakeDeleteRequest(path, token string) *http.Response {
	request := makeRequest("DELETE", path, token, nil)

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Fatalln("An error occurred while executing request", err)
	}

	return response
}
