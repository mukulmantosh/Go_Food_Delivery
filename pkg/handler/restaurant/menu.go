package restaurant

import (
	"Go_Food_Delivery/pkg/database/models/restaurant"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func (s *RestaurantHandler) addMenu(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()

	var menuItem restaurant.MenuItem
	if err := c.BindJSON(&menuItem); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	menuObject, err := s.service.AddMenu(ctx, &menuItem)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	} else {
		// Update Photo from UnSplash
		s.service.UpdateMenuPhoto(ctx, menuObject)
	}

	c.JSON(http.StatusCreated, gin.H{"message": "New Menu Added!"})
}

func (s *RestaurantHandler) listMenus(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()
	restaurantId := c.Query("restaurant_id")
	if restaurantId == "" {
		results, err := s.service.ListAllMenus(ctx)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, results)
		return

	} else {
		restaurantID, err := strconv.ParseInt(restaurantId, 10, 64)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Invalid RestaurantID"})
			return
		}
		results, err := s.service.ListMenus(ctx, restaurantID)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		if len(results) == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "No results found"})
			return
		}
		c.JSON(http.StatusOK, results)
		return
	}
}

func (s *RestaurantHandler) deleteMenu(c *gin.Context) {
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

	_, err = s.service.DeleteMenu(ctx, menuId, restaurantId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)

}
