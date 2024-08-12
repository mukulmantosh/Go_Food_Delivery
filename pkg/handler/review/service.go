package review

import (
	reviewValidate "Go_Food_Delivery/pkg/database/models/review"
	"Go_Food_Delivery/pkg/handler"
	"Go_Food_Delivery/pkg/service/review"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"log/slog"
)

type ReviewProtectedHandler struct {
	serve      *handler.Server
	group      string
	router     gin.IRoutes
	service    *review.ReviewService
	middleware []gin.HandlerFunc
	validate   *validator.Validate
}

func NewReviewProtectedHandler(s *handler.Server, groupName string,
	service *review.ReviewService, middleware []gin.HandlerFunc, validate *validator.Validate) {

	reviewHandler := &ReviewProtectedHandler{
		s,
		groupName,
		nil,
		service,
		middleware,
		validate,
	}

	reviewHandler.router = reviewHandler.registerGroup(middleware...)
	reviewHandler.routes()
	reviewHandler.registerValidator()

}

func (s *ReviewProtectedHandler) registerValidator() {
	err := s.validate.RegisterValidation("rating", reviewValidate.RatingValidator)
	if err != nil {
		slog.Error("registerValidator", "NewReviewProtectedHandler", err)
	}

}
