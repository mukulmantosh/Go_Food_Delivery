package review

import (
	"Go_Food_Delivery/pkg/handler"
	"Go_Food_Delivery/pkg/service/review"
	"github.com/gin-gonic/gin"
)

type ReviewProtectedHandler struct {
	Serve      *handler.Server
	group      string
	router     gin.IRoutes
	service    *review.ReviewService
	middleware []gin.HandlerFunc
}

func NewReviewProtectedHandler(s *handler.Server, groupName string,
	service *review.ReviewService, middleware []gin.HandlerFunc) {

	reviewHandler := &ReviewProtectedHandler{
		s,
		groupName,
		nil,
		service,
		middleware,
	}

	reviewHandler.router = reviewHandler.registerGroup(middleware...)
	reviewHandler.routes()
}
