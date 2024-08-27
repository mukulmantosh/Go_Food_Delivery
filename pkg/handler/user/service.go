package user

import (
	userValidate "Go_Food_Delivery/pkg/database/models/user"
	"Go_Food_Delivery/pkg/handler"
	"Go_Food_Delivery/pkg/service/user"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserHandler struct {
	Serve    *handler.Server
	group    string
	router   *gin.RouterGroup
	service  *user.UsrService
	validate *validator.Validate
}

func NewUserHandler(s *handler.Server, groupName string, service *user.UsrService, validate *validator.Validate) {

	usrHandler := &UserHandler{
		s,
		groupName,
		&gin.RouterGroup{},
		service,
		validate,
	}
	usrHandler.router = usrHandler.registerGroup()
	usrHandler.routes()
	usrHandler.registerValidator()
}

func (s *UserHandler) registerValidator() {
	_ = s.validate.RegisterValidation("name", userValidate.NameValidator)
	_ = s.validate.RegisterValidation("email", userValidate.EmailValidator)
}
