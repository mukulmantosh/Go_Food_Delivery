package restaurant

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Restaurant) registerGroup() *gin.RouterGroup {
	return s.registerServe.Gin().Group(s.group)
}

func (s *Restaurant) routes() http.Handler {
	s.router.POST("/", s.addRestaurant)
	return s.registerServe.Gin()
}
