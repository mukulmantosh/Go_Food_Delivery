package cart

import (
	"Go_Food_Delivery/pkg/database/models/cart"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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

	var cartItemParam cart.CartItemParams
	cartItemParam.CartID = cartID

	if err := c.BindJSON(&cartItemParam); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	cartItem := &cart.CartItems{
		CartID:       cartItemParam.CartID,
		ItemID:       cartItemParam.ItemID,
		RestaurantID: cartItemParam.RestaurantID,
		Quantity:     cartItemParam.Quantity,
	}

	_, err = s.service.AddItem(ctx, cartItem)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Items added to cart!"})

}

func (s *CartHandler) getItems(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	userID := c.GetInt64("userID")
	cartInfo, err := s.service.GetCartId(ctx, userID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	items, err := s.service.ListItems(ctx, cartInfo.CartID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"items": items})
	return
}

func (s *CartHandler) deleteItemFromCart(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	cartItemId := c.Param("id")
	cartItemID, _ := strconv.ParseInt(cartItemId, 10, 64)

	_, err := s.service.DeleteItem(ctx, cartItemID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)

}
