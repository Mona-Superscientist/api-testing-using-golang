package utils

import (
	"api-testing-example/model"
	"fmt"
	"math/rand"
	"time"

	"github.com/go-faker/faker/v4"
)

func getFormattedDate() string {
	date := time.Now().UTC().Format("2006-01-01")
	fmt.Println("date is", date)
	return date
}

func generateBookingDates() model.BookingDates {
	booingDate := model.BookingDates{
		Checkin:  getFormattedDate(),
		Checkout: getFormattedDate(),
	}
	return booingDate
}

func GenerateBooking() model.Booking {
	booking := model.Booking{
		FirstName:       faker.FirstName(),
		LastName:        faker.LastName(),
		TotalPrice:      rand.Intn(10000),
		DepositPaid:     true,
		BookingDates:    generateBookingDates(),
		AdditionalNeeds: faker.Sentence(),
	}

	return booking
}
