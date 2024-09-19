package annoucements

import (
	"Go_Food_Delivery/pkg/handler"
	"Go_Food_Delivery/pkg/service/announcements"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AnnouncementHandler struct {
	serve             *handler.Server
	group             string
	middlewareGuarded gin.IRoutes
	router            gin.IRoutes
	service           *announcements.AnnouncementService
	middleware        []gin.HandlerFunc
	validate          *validator.Validate
}

func NewAnnouncementHandler(s *handler.Server, group string,
	service *announcements.AnnouncementService, middleware []gin.HandlerFunc,
	validate *validator.Validate) {

	cartHandler := &AnnouncementHandler{
		s,
		group,
		nil,
		nil,
		service,
		middleware,
		validate,
	}
	cartHandler.middlewareGuarded = cartHandler.registerMiddlewareGroup(middleware...)
	cartHandler.router = cartHandler.registerGroup()
	cartHandler.regularRoutes()
	cartHandler.middlewareRoutes()
	cartHandler.registerValidator()
}

func (s *AnnouncementHandler) registerValidator() {

}
