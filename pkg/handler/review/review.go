package review

import (
	reviewModel "Go_Food_Delivery/pkg/database/models/review"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func (s *ReviewProtectedHandler) addReview(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	userID := c.GetInt64("userID")
	restaurantId, err := strconv.ParseInt(c.Param("restaurant_id"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Invalid RestaurantID"})
		return
	}

	var reviewParam reviewModel.ReviewParams
	var review reviewModel.Review
	if err := c.BindJSON(&reviewParam); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if err := s.validate.Struct(reviewParam); err != nil {
		validationError := reviewModel.ReviewValidationError(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": validationError})
		return
	}

	comment := reviewParam.Comment
	rating := reviewParam.Rating
	review.UserID = userID
	review.RestaurantID = restaurantId
	review.Rating = rating
	review.Comment = comment
	_, err = s.service.Add(ctx, &review)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Review Added!"})

}

func (s *ReviewProtectedHandler) listReviews(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	restaurantId, err := strconv.ParseInt(c.Param("restaurant_id"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Invalid RestaurantID"})
		return
	}

	results, err := s.service.ListReviews(ctx, restaurantId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	if len(results) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No results found"})
		return
	}
	c.JSON(http.StatusOK, results)
}

func (s *ReviewProtectedHandler) deleteReview(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	reviewId := c.Param("review_id")
	userID := c.GetInt64("userID")

	// Convert to integer
	reviewID, _ := strconv.ParseInt(reviewId, 10, 64)

	_, err := s.service.DeleteReview(ctx, reviewID, userID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)

}
