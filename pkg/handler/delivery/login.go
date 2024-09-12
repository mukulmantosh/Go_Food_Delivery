package delivery

import (
	"Go_Food_Delivery/pkg/database/models/delivery"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func (s *DeliveryHandler) loginDelivery(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()
	var token string
	var deliverLoginPerson delivery.DeliveryLoginParams

	if err := c.BindJSON(&deliverLoginPerson); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	verify := s.service.Verify(ctx, deliverLoginPerson.Phone, deliverLoginPerson.OTP)
	if !verify {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Either Phone or OTP is incorrect or user is inactive. Please contact administrator."})
		return
	} else {
		deliveryLoginDetails, err := s.service.ValidateAccountDetails(ctx, deliverLoginPerson.Phone)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Unable to fetch delivery person details. Please contact administrator."})
			return
		}
		token, err = s.service.GenerateJWT(ctx, deliveryLoginDetails.DeliveryPersonID, deliveryLoginDetails.Name)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Unable to generate login information. Please contact administrator."})
			return
		}

	}

	c.JSON(http.StatusCreated, gin.H{"token": token})
}
