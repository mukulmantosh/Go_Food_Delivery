package user

import (
	"Go_Food_Delivery/pkg/handler"
	"Go_Food_Delivery/pkg/service/user"
	"github.com/gin-gonic/gin"
)

type Register struct {
	Serve   *handler.Server
	group   string
	router  *gin.RouterGroup
	service *user.UsrService
}

func NewRegister(s *handler.Server, groupName string, service *user.UsrService) *Register {

	registrationService := &Register{
		s,
		groupName,
		&gin.RouterGroup{},
		service,
	}

	registrationService.router = registrationService.registerGroup()
	registrationService.routes()
	return registrationService
}
