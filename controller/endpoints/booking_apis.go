package endpoints

import (
	"api-testing-example/controller"
	"api-testing-example/model"
	"api-testing-example/utils"
	"fmt"
	"net/http"
)

func GetBookingIds() *http.Response {
	response := controller.MakeGetRequest("/booking", "")
	return response
}

func CreateBooking(booking model.Booking, token string) *http.Response {
	bookingSample := utils.GetRequestBody(booking)
	response := controller.MakePostRequest("/booking", token, bookingSample)
	return response
}

func DeleteBooking(token, booking_id string) *http.Response {
	bookingPath := fmt.Sprintf("%s/%s", "/booking", booking_id)
	response := controller.MakeDeleteRequest(bookingPath, token)
	return response
}
