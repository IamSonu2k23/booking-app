package booking

type SeatPayload struct {
	ID             int    `json:"id"`
	SeatIdentifier string `json:"seat_id"`
	SeatClass      string `json:"seat_class"`
	IsBooked       bool   `json:"is_booked"`
}

type Request struct {
	Email   string `json:"email" `
	Contact int    `json:"contact-number"`
}

type BookingRequest struct {
	Email   string   `json:"email" `
	Contact int      `json:"contact-number"`
	SeatIds []string `json:"seat-ids-booked"`
}

type Response struct {
	Message      string `json:"message" `
	Booking_ID   uint   `json:"booking-id" `
	Total_Amount string `json:"total-amount"`
}
