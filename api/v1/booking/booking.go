package booking

import (
	"fmt"
	"net/http"
	"strconv"
	"ticket-booking/api/v1/seats"
	"ticket-booking/config/dbconfig"
	"ticket-booking/models"

	"github.com/gin-gonic/gin"
)

func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/booking")
	{
		g.POST("", createBooking)
		g.GET("/bookings", retrieveBooking)
	}
}

var booking BookingRequest
var payload Response

func retrieveBooking(c *gin.Context) {
	db := dbconfig.GetDB()
	var request Request

	err := c.BindJSON(&request)
	if err != nil {
		fmt.Println("Fail to Bind request :", err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	var user_bookings []models.User
	err = db.Where("email = ?", request.Email).Or("contact = ?", request.Contact).Find(&user_bookings).Error
	if err != nil {
		fmt.Println("Fail to fetch user_bookings", err)
		c.JSON(http.StatusBadRequest, "Fail to fetch user_bookings")
		return
	}

	c.JSON(http.StatusOK, user_bookings)
}

func createBooking(c *gin.Context) {
	db := dbconfig.GetDB()

	err := c.BindJSON(&booking)
	if err != nil {
		fmt.Println("Fail to Bind booking request :", err)
		c.JSON(http.StatusBadRequest, err)
		return
	}
	//check if all the seat ids booking availability
	var all_seats_available bool = true
	var seat SeatPayload
	var seat_ids_string string
	for _, v := range booking.SeatIds {
		seat_ids_string = seat_ids_string + v + " ,"
		err := db.Model(&models.Seat{}).Where("seat_identifier = ?", v).Find(&seat).Error
		if err != nil {
			fmt.Println("Fail to fetch booking status :", err)
			c.JSON(http.StatusBadRequest, "Fail to fetch booking status")
			return
		}
		if seat.IsBooked {
			all_seats_available = false
			break
		}
	}

	if !all_seats_available {
		payload.Message = "Sorry One of the seat is already booked. Please Select another seat"
		c.JSON(http.StatusForbidden, payload)
		return
	} else {
		//Get ticket current prices
		amount, err := getAmount()
		if err != nil {
			payload.Message = err.Error()
			c.JSON(http.StatusBadRequest, payload)
			return
		}
		//Block SeatIds : "is_booked" status set to "true" as block
		err = setSeatIdsStatus(c)
		if err != nil {
			payload.Message = "Fail to Block seats"
			c.JSON(http.StatusBadRequest, payload)
			return
		}
		//create booking now
		var bookTickets models.User = models.User{
			Email:         booking.Email,
			Contact:       booking.Contact,
			SeatIdsBooked: seat_ids_string,
		}
		//result := db.Create(&booking)
		result := db.Create(&bookTickets)
		if result.Error != nil {
			err = unSeatIdsStatus(c)
			if err != nil {
				payload.Message = "Fail to unBlock seats"
				c.JSON(http.StatusBadRequest, payload)
				return
			}
			fmt.Println("Unable to create booking :", result.Error.Error())
			payload.Message = "Unable to create booking"
			c.JSON(http.StatusBadRequest, payload)
			return
		} else {
			payload.Message = "booking created successfully"
			payload.Booking_ID = bookTickets.ID
			payload.Total_Amount = amount
		}

	}
	notify()
	c.JSON(http.StatusOK, payload)

}

func setSeatIdsStatus(c *gin.Context) error {
	db := dbconfig.GetDB()
	for _, v := range booking.SeatIds {
		err := db.Model(&models.Seat{}).Where("seat_identifier = ?", v).UpdateColumn("is_booked", true).Error
		if err != nil {
			fmt.Println("Fail to update seat status :", err)
			return err
		}
	}
	return nil
}

func unSeatIdsStatus(c *gin.Context) error {
	db := dbconfig.GetDB()
	for _, v := range booking.SeatIds {
		err := db.Model(&models.Seat{}).Where("seat_identifier = ?", v).UpdateColumn("is_booked", false).Error
		if err != nil {
			fmt.Println("Fail to update seat status :", err)
			return err
		}
	}
	return nil
}

func getAmount() (string, error) {
	db := dbconfig.GetDB()
	var amount float64
	for _, v := range booking.SeatIds {
		var seat models.Seat
		err := db.First(&seat, "seat_identifier = ?", v).Error
		if err != nil {
			fmt.Println("Fail to fetch seat from seat-id", err)
			return "", err
		}

		var book_count int64
		err = db.Model(&models.Seat{}).Select("is_booked").Where("seat_class = ? AND is_booked = ?", seat.SeatClass, true).Count(&book_count).Error
		if err != nil {
			fmt.Println("Fail to fetch book count for class ", err)
			return "", err
		}

		price := seats.EvaluatePrice(book_count, seat.SeatClass)
		cost, err := strconv.ParseFloat(price[1:], 64)
		if err != nil {
			fmt.Println("Fail convert string price ", err)
			return "", err
		}
		amount = amount + cost
	}

	return "$" + strconv.FormatFloat(amount, 'f', 0, 64), nil
}
