package specs

import (
	"api-testing-example/controller"
	"api-testing-example/controller/endpoints"
	"api-testing-example/model"
	"api-testing-example/utils"
	"reflect"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Booking Test Suite", func() {
	Describe("GET Bookings", func() {
		It("Should return a list of booking ids", func() {
			response := endpoints.GetBookingIds()

			controller.Assert_StatusOK(response)

			responseBody := utils.GetResponseBody(response)
			var body []interface{}
			rv := reflect.ValueOf(responseBody)
			if rv.Kind() == reflect.Slice {
				for i := 0; i < rv.Len(); i++ {
					body = append(body, rv.Index(i).Interface())
				}
			}
			for _, booking := range body {
				Expect(booking).Should(HaveKey("bookingid"))
			}
		})
	})

	Describe("Create Booking", func() {
		Context("Positive Cases", func() {
			var token string
			var booking model.Booking

			BeforeEach(func() {
				By("Generate auth token and booking sample", func() {
					token = endpoints.GetAuthToken("admin", "password123")
					booking = utils.GenerateBooking()
				})
			})

			It("Should create a new booking successfully", func() {
				response := endpoints.CreateBooking(booking, token)
				controller.Assert_StatusOK(response)
			})
		})
	})
})
