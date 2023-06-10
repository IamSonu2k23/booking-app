package dbinit

import (
	"fmt"
	"ticket-booking/config/dbconfig"
	"ticket-booking/models"
)

func Init() error {
	db := dbconfig.GetDB()

	err := db.AutoMigrate(&models.User{}, &models.Seat{}, &models.SeatPrice{})
	if err != nil {
		fmt.Println("Fail to migrate errror :", err)
		return err
	}
	//load seats price database entries
	err = loadDatabase()
	if err != nil {
		fmt.Println("Fail to Load seats-price database :", err)
		return err
	}

	return nil
}
