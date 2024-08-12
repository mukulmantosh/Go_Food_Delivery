package user

import (
	"Go_Food_Delivery/pkg/handler"
	"Go_Food_Delivery/pkg/service/user"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	Serve   *handler.Server
	group   string
	router  *gin.RouterGroup
	service *user.UsrService
}

func NewUserHandler(s *handler.Server, groupName string, service *user.UsrService) {

	usrHandler := &UserHandler{
		s,
		groupName,
		&gin.RouterGroup{},
		service,
	}
	usrHandler.router = usrHandler.registerGroup()
	usrHandler.routes()
}
