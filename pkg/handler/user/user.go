package user

import (
	userModel "Go_Food_Delivery/pkg/database/models/user"
	database "Go_Food_Delivery/pkg/service/user"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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

	userService := database.NewUserService(s.Serve.Engine())
	_, err := userService.Add(ctx, &user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})

}

func (s *Register) deleteUser(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	userId := c.Param("id")

	// Convert to integer
	userID, _ := strconv.ParseInt(userId, 10, 64)

	userService := database.NewUserService(s.Serve.Engine())
	_, err := userService.Delete(ctx, userID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)

}

func (s *Register) loginUser(c *gin.Context) {
	_, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	var user userModel.LoginUser
	if err := c.BindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	userService := database.NewUserService(s.Serve.Engine())
	decoratedLogin := database.ValidateAccount(userService.Login, userService.UserExist, userService.ValidatePassword)
	result, err := decoratedLogin(c, &userModel.LoginUser{Email: user.Email, Password: user.Password})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": result})
}
