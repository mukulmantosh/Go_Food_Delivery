package restaurant

import (
	"Go_Food_Delivery/pkg/database/models/restaurant"
	restro "Go_Food_Delivery/pkg/service/restaurant"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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
	restroService := restro.NewRestaurantService(s.Serve.Engine(), s.Environment)
	menuObject, err := restroService.AddMenu(ctx, &menuItem)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	} else {
		// Update Photo from UnSplash
		restroService.UpdateMenuPhoto(ctx, menuObject)
	}

	c.JSON(http.StatusCreated, gin.H{"message": "New Menu Added!"})
}

func (s *Restaurant) listMenus(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	restaurantId, err := strconv.ParseInt(c.Param("restaurant_id"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Invalid RestaurantID"})
		return
	}

	restroService := restro.NewRestaurantService(s.Serve.Engine(), s.Environment)
	results, err := restroService.ListMenus(ctx, restaurantId)
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

func (s *Restaurant) deleteMenu(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	menuId, err := strconv.ParseInt(c.Param("menu_id"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Invalid MenuID"})
		return
	}
	restaurantId, err := strconv.ParseInt(c.Param("restaurant_id"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Invalid RestaurantID"})
		return
	}

	restroService := restro.NewRestaurantService(s.Serve.Engine(), s.Environment)
	_, err = restroService.DeleteMenu(ctx, menuId, restaurantId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)

}
