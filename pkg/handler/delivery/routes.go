package delivery

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *DeliveryHandler) registerMiddlewareGroup(middleware ...gin.HandlerFunc) gin.IRoutes {
	return s.serve.Gin.Group(s.group).Use(middleware...)
}

func (s *DeliveryHandler) registerGroup() gin.IRoutes {
	return s.serve.Gin.Group(s.group)
}

func (s *DeliveryHandler) regularRoutes() http.Handler {
	s.router.POST("/add", s.addDeliveryPerson)
	s.router.POST("/login", s.loginDelivery)
	return s.serve.Gin
}

func (s *DeliveryHandler) middlewareRoutes() http.Handler {
	s.middlewareGuarded.POST("/update-order", s.updateOrder)
	s.middlewareGuarded.GET("/deliveries/:order_id", s.deliveryListing)
	return s.serve.Gin
}
