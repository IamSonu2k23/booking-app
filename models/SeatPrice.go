package models

type SeatPrice struct {
	ID          int    `json:"id"  gorm:"primaryKey" `
	SeatClass   string `json:"class"`
	MinPrice    string `json:"min-price"`
	NormalPrice string `json:"normal-price"`
	MaxPrice    string `json:"max-price"`
}
