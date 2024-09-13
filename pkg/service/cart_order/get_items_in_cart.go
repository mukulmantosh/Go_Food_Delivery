package cart_order

import (
	"Go_Food_Delivery/pkg/database"
	"Go_Food_Delivery/pkg/database/models/cart"
	"context"
)

func (cartSrv *CartService) ListItems(ctx context.Context, cartId int64) (*[]cart.CartItems, error) {
	var cartItems []cart.CartItems
	var relatedFields = []string{"Cart", "Restaurant", "MenuItem"}
	whereFilter := database.Filter{"cart_items.cart_id": cartId}
	err := cartSrv.db.SelectWithRelation(ctx, &cartItems, relatedFields, whereFilter)

	if err != nil {
		return nil, err
	}
	return &cartItems, nil
}
