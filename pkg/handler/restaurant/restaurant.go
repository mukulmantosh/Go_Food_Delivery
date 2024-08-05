package restaurant

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func (s *Restaurant) addRestaurant(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()
	_ = ctx

	file, _ := c.FormFile("file")
	log.Println(file.Filename)
	fmt.Printf("%T", file)

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
