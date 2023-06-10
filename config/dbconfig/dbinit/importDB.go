package dbinit

import (
	"fmt"
	"os"
	"path/filepath"
	"ticket-booking/config/dbconfig"

	"github.com/gocarina/gocsv"
)

type SeatPrice struct {
	ID          int    `csv:"id"`
	SeatClass   string `csv:"seat_class"`
	MinPrice    string `csv:"min_price"`
	NormalPrice string `csv:"normal_price"`
	MaxPrice    string `csv:"max_price"`
}

type Seat struct {
	ID             int    `csv:"id"  gorm:"primaryKey" `
	SeatIdentifier string `csv:"seat_identifier"`
	SeatClass      string `csv:"seat_class"`
	IsBooked       bool   `csv:"-"`
}

func loadDatabase() error {
	// Open the CSV file for reading
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Fail to get curr work dir :", err)
		return err
	}

	currentDir := filepath.Join(dir, "config", "dbconfig", "dbinit")
	seat_price_file_path := filepath.Join(currentDir, "SeatPricing.csv")
	file1, err := os.Open(seat_price_file_path)
	if err != nil {
		fmt.Println("Fail to open csv file :", err)
		return err
	}
	defer file1.Close()

	var price_data []SeatPrice
	err = gocsv.Unmarshal(file1, &price_data)
	if err != nil {
		fmt.Println("Fail to parse csv entries :", err)
		return err
	}

	db := dbconfig.GetDB()
	result := db.Create(&price_data)
	if result.Error != nil {
		fmt.Println("Unable to load entries :", result.Error)
		return result.Error
	}

	// loading seats data
	var seats_data []Seat
	seats_file_path := filepath.Join(currentDir, "Seats.csv")
	file2, err := os.Open(seats_file_path)
	if err != nil {
		fmt.Println("Fail to open csv file :", err)
		return err
	}
	defer file2.Close()
	err = gocsv.Unmarshal(file2, &seats_data)
	if err != nil {
		fmt.Println("Fail to parse csv entries :", err)
		return err
	}

	result = db.Create(&seats_data)
	if result.Error != nil {
		fmt.Println("Unable to load entries :", result.Error)
		return result.Error
	}

	return nil
}
