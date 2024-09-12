package delivery

import (
	"Go_Food_Delivery/pkg/database/models/delivery"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func (s *DeliveryHandler) addDeliveryPerson(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	var deliverPerson delivery.DeliveryPersonParams
	var deliverPersonModel delivery.DeliveryPerson

	if err := c.BindJSON(&deliverPerson); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	deliverPersonModel.Name = deliverPerson.Name
	deliverPersonModel.Phone = deliverPerson.Phone
	deliverPersonModel.VehicleDetails = deliverPerson.VehicleDetails
	deliverPersonModel.Status = "AVAILABLE"

	authKey, authKeyURL, err := s.service.GenerateTOTP(ctx, deliverPerson.Phone)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	deliverPersonModel.AuthKey = authKey
	deliverPersonModel.AuthKeyURL = authKeyURL
	deliverPersonModel.IsAuthSet = false

	_, err = s.service.AddDeliveryPerson(ctx, &deliverPersonModel)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Delivery Person Added Successfully!"})

}
