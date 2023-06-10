package dbconfig

import (
	"fmt"
	"os"
	"strconv"

	//"ticket-booking/config/envconfig"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func GetDB() *gorm.DB {
	if db != nil {
		return db
	}

	var (
		host     = os.Getenv("DB_HOST")
		username = os.Getenv("DB_USER")
		password = os.Getenv("DB_PASSWORD")
		dbname   = os.Getenv("DB_NAME")
		dbport   = os.Getenv("DB_PORT")
	)
	var err error
	port, err := strconv.Atoi(dbport)
	if err != nil {
		fmt.Println("Port conversion error :", err)
	}

	dns := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable port=%d", host, username, password, dbname, port)
	db, err = gorm.Open(postgres.New(postgres.Config{
		DSN: dns,
	}))
	if err != nil {
		fmt.Println("DB connection error :", err)
	}

	sql_db, err := db.DB()
	if err != nil {
		fmt.Println("Database pinging error :", err)
	}
	if err := sql_db.Ping(); err != nil {
		fmt.Println("Fail to ping database error :", err)
	}

	return db
}
