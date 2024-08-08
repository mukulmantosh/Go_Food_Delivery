package restaurant

import (
	"Go_Food_Delivery/pkg/handler"
	"github.com/gin-gonic/gin"
)

type Restaurant struct {
	Serve       *handler.Server
	group       string
	router      *gin.RouterGroup
	Environment string
}

func NewRestaurant(s *handler.Server, groupName string, Env string) *Restaurant {

	restaurantService := &Restaurant{
		s,
		groupName,
		&gin.RouterGroup{},
		Env,
	}

	restaurantService.router = restaurantService.registerGroup()
	restaurantService.routes()
	return restaurantService
}
