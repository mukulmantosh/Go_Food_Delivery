package cart

import (
	"Go_Food_Delivery/pkg/database/models/cart"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func (s *CartHandler) addToCart(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	var cartID int64

	userID := c.GetInt64("userID")

	// Check Cart Exists in DB
	cartInfo, err := s.service.GetCartId(ctx, userID)
	if err != nil {
		// Create a new cart.
		var cartData cart.Cart
		cartData.UserID = userID

		newCart, err := s.service.Create(ctx, &cartData)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		cartID = newCart.CartID
	} else {
		cartID = cartInfo.CartID
	}

	fmt.Println("cart ", cartID)

	c.JSON(http.StatusCreated, gin.H{"message": "Cart created!"})

}
