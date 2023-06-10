package models

import "time"

type User struct {
	ID            uint      `json:"id" gorm:"primaryKey"`
	Email         string    `json:"email"`
	Contact       int       `json:"contact-number"`
	SeatIdsBooked string    `json:"seat-ids-booked"`
	CreatedAt     time.Time `json:"created_at,omitempty"`
}
