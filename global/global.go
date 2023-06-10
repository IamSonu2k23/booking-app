package global

import (
	"log"
	"ticket-booking/config/dbconfig"
	"ticket-booking/models"
)

type seatCounts struct {
	SeatClass string
	Total     int64
}

var SeatCount = make(map[string]int64)

func InitGlobal() error {

	db := dbconfig.GetDB()

	//tested query
	// select seat_class, count(*) as c from seats group by seat_class order by seat_class ;
	var totalSeat []seatCounts

	err := db.Model(&models.Seat{}).Select("seat_class, count(*) as Total").Group("seat_class").Order("seat_class").Find(&totalSeat).Error
	if err != nil {
		log.Fatalf("Fail to parse csv entries : %s", err)
		return err
	}

	for _, value := range totalSeat {
		SeatCount[value.SeatClass] = value.Total
	}

	return nil
}
