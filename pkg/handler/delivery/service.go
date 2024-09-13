package delivery

import (
	"Go_Food_Delivery/pkg/handler"
	"Go_Food_Delivery/pkg/service/delivery"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type DeliveryHandler struct {
	serve             *handler.Server
	group             string
	middlewareGuarded gin.IRoutes
	router            gin.IRoutes
	service           *delivery.DeliveryService
	middleware        []gin.HandlerFunc
	validate          *validator.Validate
}

func NewDeliveryHandler(s *handler.Server, group string,
	service *delivery.DeliveryService, middleware []gin.HandlerFunc,
	validate *validator.Validate) {

	cartHandler := &DeliveryHandler{
		s,
		group,
		nil,
		nil,
		service,
		middleware,
		validate,
	}
	cartHandler.middlewareGuarded = cartHandler.registerMiddlewareGroup(middleware...)
	cartHandler.router = cartHandler.registerGroup()
	cartHandler.regularRoutes()
	cartHandler.middlewareRoutes()
	cartHandler.registerValidator()
}

func (s *DeliveryHandler) registerValidator() {

}
