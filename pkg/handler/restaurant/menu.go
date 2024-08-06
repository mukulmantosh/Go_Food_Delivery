package restaurant

import (
	"Go_Food_Delivery/pkg/database/models/restaurant"
	restro "Go_Food_Delivery/pkg/service/restaurant"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func (s *Restaurant) addMenu(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	var menuItem restaurant.MenuItem
	if err := c.BindJSON(&menuItem); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	restroService := restro.NewRestaurantService(s.Serve.Engine())
	_, err := restroService.AddMenu(ctx, &menuItem)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "New Menu Added!"})
}
