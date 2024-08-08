package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Register) registerGroup() *gin.RouterGroup {
	return s.Serve.Gin().Group(s.group)
}

func (s *Register) routes() http.Handler {
	s.router.POST("/", s.addUser)
	s.router.DELETE("/:id", s.deleteUser)
	s.router.POST("/login", s.loginUser)
	return s.Serve.Gin()
}
