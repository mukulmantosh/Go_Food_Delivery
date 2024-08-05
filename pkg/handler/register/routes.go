package register

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Register) registerGroup() *gin.RouterGroup {
	return s.Serve.Gin().Group(s.group)
}

func (s *Register) routes() http.Handler {
	s.router.POST("/user", s.addUser)
	s.router.DELETE("/user/:id", s.deleteUser)
	return s.Serve.Gin()
}
