package seats

import (
	"fmt"
	"ticket-booking/config/dbconfig"
	"ticket-booking/global"
	"ticket-booking/models"
)

type SeatPrice struct {
	ID          int
	SeatClass   string
	MinPrice    string
	NormalPrice string
	MaxPrice    string
}

var seatPrice SeatPrice

func EvaluatePrice(count int64, seat_class string) string {

	var price string
	db := dbconfig.GetDB()
	err := db.Model(&models.SeatPrice{}).Where("seat_class = ?", seat_class).Find(&seatPrice).Error
	if err != nil {
		fmt.Println("Fail to fetch seat price record", err)
	}

	total := global.SeatCount[seat_class]
	percentage := (count / total) * 100

	if percentage >= 60 {
		price = setPrice("maximum")
	} else if percentage >= 40 && percentage < 60 {
		price = setPrice("normal")
	} else {
		price = setPrice("minimum")
	}

	return price
}

func setPrice(price_level string) string {
	var status bool
	var price string
	switch price_level {
	case "minimum":
		if seatPrice.MinPrice != "" {
			price = seatPrice.MinPrice
			status = true
		}
		if status {
			break
		}
		fallthrough
	case "normal":
		if seatPrice.NormalPrice != "" {
			price = seatPrice.NormalPrice
			status = true
		}
		if status {
			break
		}
		fallthrough
	case "maximum":
		if seatPrice.MaxPrice != "" {
			price = seatPrice.MaxPrice
			status = true
		}
		if status {
			break
		}
		fallthrough
	default:
		price = seatPrice.NormalPrice
	}

	return price
}
