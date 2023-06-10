package booking

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func notify() {
	api_key := os.Getenv("SMS_API_KEY")
	url := "https://api.melroselabs.com/sms/message"
	method := "POST"
	msgDetail := "Booked Successfully. Your Seats are " + strings.Join(booking.SeatIds, ",")

	//payload := strings.NewReader("{ \"destination\": \"447712345678\", \"message\": \"Hello World #$£\", \"source\": \"MelroseLabs\" }")
	payload := strings.NewReader("{ \"destination\": \"" + strconv.Itoa(booking.Contact) + "\", \"message\": \" " + msgDetail + " \", \"source\": \"Booking App\" }")
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("x-api-key", api_key)
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))
}
