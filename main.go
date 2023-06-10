package main

import (
	"fmt"
	"log"
	"os"
	"ticket-booking/app"
)

func main() {
	app.Init()
	addr := fmt.Sprintf(":%s", os.Getenv("APP_PORT"))
	err := app.GinApp.Run(addr)
	if err != nil {
		log.Fatal(err)
	}
}
