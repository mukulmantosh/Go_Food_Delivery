package cart

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *CartHandler) registerGroup(middleware ...gin.HandlerFunc) gin.IRoutes {
	return s.serve.Gin.Group(s.group).Use(middleware...)
}

func (s *CartHandler) routes() http.Handler {
	s.router.POST("/add", s.addToCart)
	s.router.GET("/list", s.getItems)
	s.router.DELETE("/remove/:id", s.deleteItemFromCart)
	s.router.POST("/order/new", s.PlaceNewOrder)

	return s.serve.Gin
}
