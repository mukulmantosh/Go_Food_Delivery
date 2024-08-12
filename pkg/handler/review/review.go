package review

import (
	reviewModel "Go_Food_Delivery/pkg/database/models/review"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func (s *ReviewProtectedHandler) addReview(c *gin.Context) {
	_, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	userID, _ := c.Get("userID")
	fmt.Println("addReview", userID)

	//if r.Rating < 1 || r.Rating > 5 {
	//	return fmt.Errorf("rating must be between 1 and 5")
	//}

	var review reviewModel.ReviewParams
	if err := c.BindJSON(&review); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if err := s.validate.Struct(review); err != nil {
		validationError := reviewModel.ReviewValidationError(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": validationError})
		return
	}

	//_, err := s.service.Add(ctx, &review)
	//if err != nil {
	//	c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
	//	return
	//}

	c.JSON(http.StatusCreated, gin.H{"message": "Review Added!"})

}
