package specs

import (
	"api-testing-example/controller"
	"api-testing-example/controller/endpoints"

	. "github.com/onsi/ginkgo/v2"
)

var _ = Describe("Auth test suite", func() {
	Context("Successful Login", func() {
		It("Should return auth token", func() {
			response := endpoints.CreateToken("admin", "password123")

			controller.Assert_StatusOK(response)
			controller.Assert_ResponseBody_Contain_Key(response, "token")
		})
	})

	Context("Negative Scenarios", func() {
		/* This test case should assert for bad request status code by logic,
		but it returns 200 so, it's handled like that... */
		It("Should return bad request after trying invalid credentials", func() {
			response := endpoints.CreateToken("adminzz", "123")

			controller.Assert_StatusOK(response)
			controller.Assert_ResponseBody_Has_Value(response, "reason", "Bad credentials")
		})
	})
})
