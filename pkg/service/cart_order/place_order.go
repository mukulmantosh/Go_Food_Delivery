package cart_order

import (
	"Go_Food_Delivery/pkg/database"
	"Go_Food_Delivery/pkg/database/models/cart"
	"Go_Food_Delivery/pkg/database/models/order"
	"context"
	"errors"
	"fmt"
)

func (cartSrv *CartService) PlaceOrder(ctx context.Context, cartId int64, userId int64) (*order.Order, error) {
	var cartItems []cart.CartItems
	var newOrder order.Order
	var newOrderItems []order.OrderItems
	var orderTotal float64 = 0.0
	var relatedFields = []string{"MenuItem"}
	whereFilter := database.Filter{"cart_items.cart_id": cartId}

	err := cartSrv.db.SelectWithRelation(ctx, &cartItems, relatedFields, whereFilter)
	if err != nil {
		return nil, err
	}

	if len(cartItems) == 0 {
		return nil, errors.New("no items in cart")
	}

	// Creating a new order.
	newOrder.UserID = userId
	newOrder.OrderStatus = "pending"
	newOrder.TotalAmount = orderTotal
	newOrder.DeliveryAddress = "New Delhi"

	_, err = cartSrv.db.Insert(ctx, &newOrder)
	if err != nil {
		return nil, err
	}

	newOrderItems = make([]order.OrderItems, len(cartItems))
	for i, cartItem := range cartItems {
		newOrderItems[i].OrderID = newOrder.OrderID
		newOrderItems[i].ItemID = cartItem.ItemID
		newOrderItems[i].RestaurantID = cartItem.RestaurantID
		newOrderItems[i].Quantity = cartItem.Quantity
		newOrderItems[i].Price = cartItem.MenuItem.Price * float64(cartItem.Quantity)
		_, err = cartSrv.db.Insert(ctx, &newOrderItems[i])
		if err != nil {
			return nil, err
		}
		orderTotal += newOrderItems[i].Price
	}

	_, err = cartSrv.db.Update(ctx, "orders", database.Filter{"total_amount": orderTotal, "order_status": "in_progress"},
		database.Filter{"order_id": newOrder.OrderID})
	if err != nil {
		return nil, err
	}

	return &newOrder, nil

}

func (cartSrv *CartService) RemoveItemsFromCart(ctx context.Context, cartId int64) error {
	//remove all items from the cart.
	filter := database.Filter{"cart_id": cartId}
	_, err := cartSrv.db.Delete(ctx, "cart_items", filter)
	if err != nil {
		return errors.New("failed to delete cart items")
	}
	return nil
}

func (cartSrv *CartService) NewOrderPlacedNotification(userId int64, orderId int64) error {
	message := fmt.Sprintf("USER_ID:%d|MESSAGE:Your order number %d has been successfully placed, and the chef has begun the cooking process.", userId, orderId)
	topic := fmt.Sprintf("orders.new.%d", userId)
	err := cartSrv.nats.Pub(topic, []byte(message))
	if err != nil {
		return err
	}
	return nil
}
