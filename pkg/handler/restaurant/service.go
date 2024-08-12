package restaurant

import (
	"Go_Food_Delivery/pkg/handler"
	"Go_Food_Delivery/pkg/service/restaurant"
	"github.com/gin-gonic/gin"
)

type RestaurantHandler struct {
	Serve   *handler.Server
	group   string
	router  *gin.RouterGroup
	service *restaurant.RestaurantService
}

func NewRestaurantHandler(s *handler.Server, groupName string, service *restaurant.RestaurantService) {

	restroHandler := &RestaurantHandler{
		s,
		groupName,
		&gin.RouterGroup{},
		service,
	}

	restroHandler.router = restroHandler.registerGroup()
	restroHandler.routes()
}
