package register

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Register) registerGroup() *gin.RouterGroup {
	return s.registerServe.Gin().Group(s.group)
}

func (s *Register) routes() http.Handler {
	s.router.POST("/user", s.addUser)
	return s.registerServe.Gin()
}
