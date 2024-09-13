package delivery

import (
	"Go_Food_Delivery/pkg/database/models/delivery"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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

	userID := c.GetInt64("userID")

	_, err := s.service.OrderPlacement(ctx, userID, deliveryOrder.OrderID, deliveryOrder.Status)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Order Updated!"})

}

func (s *DeliveryHandler) deliveryListing(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	orderId := c.Param("order_id")

	// Convert to integer
	orderID, _ := strconv.ParseInt(orderId, 10, 64)
	userID := c.GetInt64("userID")

	deliveries, err := s.service.DeliveryListing(ctx, orderID, userID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"deliveries": deliveries})
}
