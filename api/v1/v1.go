package apiv1

import (
	"ticket-booking/api/v1/booking"
	"ticket-booking/api/v1/seats"

	"github.com/gin-gonic/gin"
)

func ApplyRoutes(r *gin.RouterGroup) {
	v1 := r.Group("/v1.0")
	{
		seats.ApplyRoutes(v1)
		booking.ApplyRoutes(v1)
	}
}
