package register

import (
	userModel "Go_Food_Delivery/pkg/database/models/user"
	database "Go_Food_Delivery/pkg/service/user"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func (s *Register) addUser(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	var user userModel.User
	if err := c.BindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	userService := database.NewUserService(s.registerServe.Engine())
	_, err := userService.Add(ctx, &user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})

}
