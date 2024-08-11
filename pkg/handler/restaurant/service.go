package restaurant

import (
	"Go_Food_Delivery/pkg/handler"
	"Go_Food_Delivery/pkg/service/restaurant"
	"github.com/gin-gonic/gin"
)

type Restaurant struct {
	Serve   *handler.Server
	group   string
	router  *gin.RouterGroup
	service *restaurant.RestaurantService
}

func NewRestaurant(s *handler.Server, groupName string, service *restaurant.RestaurantService) *Restaurant {

	restaurantService := &Restaurant{
		s,
		groupName,
		&gin.RouterGroup{},
		service,
	}

	restaurantService.router = restaurantService.registerGroup()
	restaurantService.routes()
	return restaurantService
}
