package cart

import (
	"Go_Food_Delivery/pkg/database"
	"Go_Food_Delivery/pkg/database/models/cart"
	"Go_Food_Delivery/pkg/database/models/order"
	"context"
	"errors"
)

func (cartSrv *CartService) PlaceOrder(ctx context.Context, cartId int64) (*order.Order, error) {
	var cartItems []cart.CartItems
	var newOrder order.Order
	var newOrderItems []order.OrderItems
	var orderTotal float64 = 0.0
	var relatedFields = []string{"MenuItem"}
	err := cartSrv.db.SelectWithRelation(ctx, &cartItems, relatedFields, "cart_items.cart_id", cartId)
	if err != nil {
		return nil, err
	}

	if len(cartItems) == 0 {
		return nil, errors.New("no items in cart")
	}

	// Creating a new order.
	newOrder.UserID = 1
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

	//newOrder.TotalAmount = orderTotal
	_, err = cartSrv.db.Update(ctx, "orders", database.Filter{"total_amount": orderTotal, "order_status": "in_progress"},
		database.Filter{"order_id": newOrder.OrderID})
	if err != nil {
		return nil, err
	}

	return nil, err
}

//type CartItems struct {
//	bun.BaseModel `bun:"table:cart_items"`
//	CartItemID    int64 `bun:",pk,autoincrement" json:"cart_item_id"`
//	CartID        int64 `bun:"cart_id,notnull" json:"cart_id"`
//	ItemID        int64 `bun:"item_id,notnull" json:"item_id"`
//	RestaurantID  int64 `bun:"restaurant_id,notnull" json:"restaurant_id"`
//	Quantity      int64 `bun:"quantity,notnull" json:"quantity"`
//	utils.Timestamp
//	Restaurant *restaurant.Restaurant `bun:"rel:belongs-to,join:restaurant_id=restaurant_id" json:"-"`
//	MenuItem   *restaurant.MenuItem   `bun:"rel:belongs-to,join:item_id=menu_id" json:"menu_item"`
//	Cart       *Cart                  `bun:"rel:belongs-to,join:cart_id=cart_id" json:"-"`
//}

//type Order struct {
//	bun.BaseModel   `bun:"table:orders"`
//	OrderID         int64   `bun:",pk,autoincrement" json:"order_id"`
//	UserID          int64   `bun:"user_id,notnull" json:"user_id"`
//	OrderStatus     string  `bun:"order_status,notnull" json:"order_status"`
//	TotalAmount     float64 `bun:"total_amount,notnull" json:"total_amount"`
//	DeliveryAddress string  `bun:"delivery_address,notnull" json:"delivery_address"`
//	utils.Timestamp
//	User *userModel.User `bun:"rel:belongs-to,join:user_id=id"`
//}
//
//type OrderItems struct {
//	bun.BaseModel `bun:"table:order_items"`
//	OrderItemID   int64   `bun:",pk,autoincrement" json:"order_item_id"`
//	OrderID       int64   `bun:"order_id,notnull" json:"order_id"`
//	ItemID        int64   `bun:"item_id,notnull" json:"item_id"`
//	RestaurantID  int64   `bun:"restaurant_id,notnull" json:"restaurant_id"`
//	Quantity      int64   `bun:"quantity,notnull" json:"quantity"`
//	Price         float64 `bun:"price,notnull" json:"price"`
//	utils.Timestamp
//	MenuItem   *restaurant.MenuItem   `bun:"rel:belongs-to,join:item_id=menu_id" json:"-"`
//	Restaurant *restaurant.Restaurant `bun:"rel:belongs-to,join:restaurant_id=restaurant_id"`
//	Order      *Order                 `bun:"rel:belongs-to,join:order_id=order_id" json:"-"`
//}
