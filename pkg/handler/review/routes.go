package review

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *ReviewProtectedHandler) registerGroup(middleware ...gin.HandlerFunc) gin.IRoutes {
	return s.serve.Gin().Group(s.group).Use(middleware...)
}

func (s *ReviewProtectedHandler) routes() http.Handler {
	s.router.POST("/", s.addReview)
	return s.serve.Gin()
}
