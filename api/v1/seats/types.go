package seats

type SeatsData struct {
	SeatClass      string `json:"class"`
	SeatIdentifier string `json:"seat-id"`
	IsBooked       bool   `json:"isBooked"`
}

type Response struct {
	ID             int    `json:"id"`
	SeatIdentifier string `json:"seat_id"`
	SeatClass      string `json:"seat_class"`
	IsBooked       bool   `json:"is_booked"`
	Price          string `json:"price"`
}
