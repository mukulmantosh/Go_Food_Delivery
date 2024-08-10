package user

import (
	"Go_Food_Delivery/pkg/handler"
	"github.com/gin-gonic/gin"
)

type Register struct {
	Serve  *handler.Server
	group  string
	router *gin.RouterGroup
}

func NewRegister(s *handler.Server, groupName string) *Register {

	registrationService := &Register{
		s,
		groupName,
		&gin.RouterGroup{},
	}

	registrationService.router = registrationService.registerGroup()
	registrationService.routes()
	return registrationService
}
