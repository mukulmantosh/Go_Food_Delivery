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

func (s *UserHandler) addUser(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	var user userModel.User
	if err := c.BindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if err := s.validate.Struct(user); err != nil {
		validationError := userModel.UserValidationError(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": validationError})
		return
	}

	_, err := s.service.Add(ctx, &user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})

}

func (s *UserHandler) deleteUser(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	userId := c.Param("id")

	// Convert to integer
	userID, _ := strconv.ParseInt(userId, 10, 64)

	_, err := s.service.Delete(ctx, userID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)

}

func (s *UserHandler) loginUser(c *gin.Context) {
	_, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	var user userModel.LoginUser
	if err := c.BindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	login := database.ValidateAccount(s.service.Login, s.service.UserExist, s.service.ValidatePassword)
	result, err := login(c, &userModel.LoginUser{Email: user.Email, Password: user.Password})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": result})
}
