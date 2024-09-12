package delivery

import (
	"Go_Food_Delivery/pkg/database/models/delivery"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func (s *DeliveryHandler) updateOrder(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	var deliveryOrder delivery.DeliveryOrderPlacementParams
	if err := c.BindJSON(&deliveryOrder); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	_, err := s.service.OrderPlacement(ctx, deliveryOrder.OrderID, deliveryOrder.Status)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Order Updated!"})

}
