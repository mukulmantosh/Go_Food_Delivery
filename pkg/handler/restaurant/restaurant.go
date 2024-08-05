package restaurant

import (
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

	_, err = s.registerServe.Storage().Upload(newFileName, file)
	if err != nil {
		fmt.Println("Error:", err)
	}

	uploadedFile := filepath.Join(os.Getenv("STORAGE_DIRECTORY"), newFileName)
	fmt.Println("UPLOAD", uploadedFile)

	// Upload the file to specific dst.
	//.SaveUploadedFile(file, dst)

	//var user userModel.User
	//if err := c.BindJSON(&user); err != nil {
	//	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
	//	return
	//}
	//
	//userService := database.NewUserService(s.registerServe.Engine())
	//_, err := userService.Add(ctx, &user)
	//if err != nil {
	//	c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
	//	return
	//}

	c.JSON(http.StatusCreated, gin.H{"message": "Restaurant created successfully"})

}
