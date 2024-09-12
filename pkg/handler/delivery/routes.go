package delivery

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *DeliveryHandler) registerGroup(middleware ...gin.HandlerFunc) gin.IRoutes {
	return s.serve.Gin.Group(s.group).Use(middleware...)
}

func (s *DeliveryHandler) routes() http.Handler {
	s.router.POST("/add", s.addDeliveryPerson)
	s.router.POST("/login", s.loginDelivery)

	return s.serve.Gin
}
