package register

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Register) addUser(c *gin.Context) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	c.JSON(http.StatusOK, resp)
}
