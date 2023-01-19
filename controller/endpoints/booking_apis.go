package endpoints

import (
	"api-testing-example/controller"
	"api-testing-example/model"
	"api-testing-example/utils"
	"net/http"
)

func GetBookingIds() *http.Response {
	response := controller.MakeGetRequest("/booking")
	return response
}

func CreateBooking(booking model.Booking, token string) *http.Response {
	bookingSample := utils.GetRequestBody(booking)
	response := controller.MakePostRequestWithAuthToken("/booking", token, bookingSample)
	return response
}
