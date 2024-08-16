package review

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *ReviewProtectedHandler) registerGroup(middleware ...gin.HandlerFunc) gin.IRoutes {
	return s.serve.Gin.Group(s.group).Use(middleware...)
}

func (s *ReviewProtectedHandler) routes() http.Handler {
	s.router.POST("/:restaurant_id", s.addReview)
	s.router.GET("/:restaurant_id", s.listReviews)
	s.router.DELETE("/:review_id", s.deleteReview)

	return s.serve.Gin
}
