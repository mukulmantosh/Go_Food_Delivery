package cart

import (
	"Go_Food_Delivery/pkg/database/models/cart"
	"context"
)

func (cartSrv *CartService) ListItems(ctx context.Context, cartId int64) (*[]cart.CartItems, error) {
	var cartItems []cart.CartItems

	err := cartSrv.db.Select(ctx, &cartItems, "cart_id", cartId)

	if err != nil {
		return nil, err
	}
	return &cartItems, nil
}
