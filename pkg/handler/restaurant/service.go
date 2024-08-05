package restaurant

import (
	"Go_Food_Delivery/pkg/handler"
	"github.com/gin-gonic/gin"
)

type Restaurant struct {
	Serve  *handler.Server
	group  string
	router *gin.RouterGroup
}

func NewRestaurant(s *handler.Server, groupName string) *Restaurant {

	restaurantService := &Restaurant{
		s,
		groupName,
		&gin.RouterGroup{},
	}

	restaurantService.router = restaurantService.registerGroup()
	restaurantService.routes()
	return restaurantService
}
