package register

import (
	"Uber_Food_Delivery/internal/server"
	"github.com/gin-gonic/gin"
)

type Register struct {
	registerServe *server.Server
	group         string
	router        *gin.RouterGroup
}

func NewRegister(s *server.Server, groupName string) *Register {

	registrationService := &Register{
		s,
		groupName,
		&gin.RouterGroup{},
	}
	registrationService.router = registrationService.registerGroup()
	registrationService.routes()
	return registrationService
}
