package models

type Seat struct {
	ID             int    `json:"id"  gorm:"primaryKey" `
	SeatIdentifier string `json:"seat_id"`
	SeatClass      string `json:"seat_class"`
	IsBooked       bool   `json:"is_booked"`
}
