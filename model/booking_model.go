package model

type BookingDates struct {
	Checkin  string `json:"checkin"`
	Checkout string `json:"checkout"`
}

type Booking struct {
	FirstName       string       `json:"firstname"`
	LastName        string       `json:"lastname"`
	TotalPrice      int          `json:"totalprice"`
	DepositPaid     bool         `json:"depositpaid"`
	BookingDates    BookingDates `json:"bookingdates"`
	AdditionalNeeds string       `json:"additionalneeds"`
}

type BookingResponse struct {
	Id      int     `json:"bookingid"`
	Booking Booking `json:"booking"`
}
