package delivery

import (
	"Go_Food_Delivery/pkg/handler"
	"Go_Food_Delivery/pkg/service/delivery"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type DeliveryHandler struct {
	serve      *handler.Server
	group      string
	router     gin.IRoutes
	service    *delivery.DeliveryService
	middleware []gin.HandlerFunc
	validate   *validator.Validate
}

func NewDeliveryHandler(s *handler.Server, groupName string,
	service *delivery.DeliveryService, middleware []gin.HandlerFunc,
	validate *validator.Validate) {

	cartHandler := &DeliveryHandler{
		s,
		groupName,
		nil,
		service,
		middleware,
		validate,
	}
	cartHandler.router = cartHandler.registerGroup(middleware...)
	cartHandler.routes()
	cartHandler.registerValidator()
}

func (s *DeliveryHandler) registerValidator() {

}
