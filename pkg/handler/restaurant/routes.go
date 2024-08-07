package restaurant

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Restaurant) registerGroup() *gin.RouterGroup {
	return s.Serve.Gin().Group(s.group)
}

func (s *Restaurant) routes() http.Handler {
	s.router.POST("/", s.addRestaurant)
	s.router.GET("/", s.listRestaurants)
	s.router.DELETE("/:id", s.deleteRestaurant)
	s.router.POST("/menu", s.addMenu)
	s.router.GET("/menu/:restaurant_id", s.listMenus)
	return s.Serve.Gin()
}
