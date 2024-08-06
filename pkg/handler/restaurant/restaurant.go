package restaurant

import (
	restaurantModel "Go_Food_Delivery/pkg/database/models/restaurant"
	restro "Go_Food_Delivery/pkg/service/restaurant"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func (s *Restaurant) addRestaurant(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()
	_ = ctx

	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File is required"})
		return
	}

	originalFileName := fileHeader.Filename

	// Generate a new file name
	newFileName := generateFileName(originalFileName)

	_, err = s.Serve.Storage().Upload(newFileName, file)
	if err != nil {
		fmt.Println("Error:", err)
	}

	uploadedFile := filepath.Join(os.Getenv("STORAGE_DIRECTORY"), newFileName)

	var restaurant restaurantModel.Restaurant
	restaurant.Name = c.PostForm("name")
	restaurant.Description = c.PostForm("description")
	restaurant.Address = c.PostForm("address")
	restaurant.City = c.PostForm("city")
	restaurant.State = c.PostForm("state")
	restaurant.Photo = uploadedFile

	restroService := restro.NewRestaurantService(s.Serve.Engine())
	_, err = restroService.Add(ctx, &restaurant)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Restaurant created successfully"})
}

func (s *Restaurant) ListRestaurants(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	restroService := restro.NewRestaurantService(s.Serve.Engine())
	results, err := restroService.ListRestaurants(ctx)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, results)
}
