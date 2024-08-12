package review

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func (s *ReviewProtectedHandler) addReview(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	userID, _ := c.Get("userID")
	fmt.Println("addReview", userID)

	_ = ctx
	//var user userModel.User
	//if err := c.BindJSON(&user); err != nil {
	//	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
	//	return
	//}
	//
	//_, err := s.service.Add(ctx, &user)
	//if err != nil {
	//	c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
	//	return
	//}

	c.JSON(http.StatusCreated, gin.H{"message": "Review Added!"})

}
