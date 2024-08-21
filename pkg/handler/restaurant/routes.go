package restaurant

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *RestaurantHandler) registerGroup() *gin.RouterGroup {
	return s.Serve.Gin.Group(s.group)
}

func (s *RestaurantHandler) routes() http.Handler {
	s.router.POST("/", s.addRestaurant)
	s.router.GET("/", s.listRestaurants)
	s.router.GET("/:id", s.listRestaurantById)
	s.router.DELETE("/:id", s.deleteRestaurant)
	s.router.POST("/menu", s.addMenu)
	s.router.GET("/menu", s.listMenus)
	s.router.DELETE("/menu/:restaurant_id/:menu_id", s.deleteMenu)
	return s.Serve.Gin
}
