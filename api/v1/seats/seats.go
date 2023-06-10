package seats

import (
	"fmt"
	"net/http"
	"ticket-booking/config/dbconfig"
	"ticket-booking/models"

	"github.com/gin-gonic/gin"
)

func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/seats")
	{
		g.GET("", getSeats)
		g.GET("/:id", GetSeatById)
	}
}

func getSeats(c *gin.Context) {

	var seats []models.Seat
	db := dbconfig.GetDB()

	err := db.Order("seat_class").Find(&seats).Error
	if err != nil {
		fmt.Println("Fail to find all seats", err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, seats)
}

func GetSeatById(c *gin.Context) {

	var seat models.Seat
	seat_identifier := c.Param("id")

	db := dbconfig.GetDB()
	err := db.First(&seat, "seat_identifier = ?", seat_identifier).Error
	if err != nil {
		fmt.Println("Fail to fetch seat id", err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	//create response fro the request
	var response Response = Response{
		ID:             seat.ID,
		SeatIdentifier: seat.SeatIdentifier,
		SeatClass:      seat.SeatClass,
		IsBooked:       seat.IsBooked,
	}

	var book_count int64
	err = db.Model(&models.Seat{}).Select("is_booked").Where("seat_class = ? AND is_booked = ?", seat.SeatClass, true).Count(&book_count).Error
	if err != nil {
		fmt.Println("Fail to fetch book count for class ", err)
		c.JSON(http.StatusBadRequest, "Fail to fetch book count for class ")
		return
	}

	response.Price = EvaluatePrice(book_count, seat.SeatClass)

	c.JSON(http.StatusOK, response)

}
