package cart_order

import (
	"Go_Food_Delivery/pkg/database/models/cart"
	"context"
)

func (cartSrv *CartService) GetCartId(ctx context.Context, UserId int64) (*cart.Cart, error) {
	var cartInfo cart.Cart

	err := cartSrv.db.Select(ctx, &cartInfo, "user_id", UserId)
	if err != nil {
		return nil, err
	}
	return &cartInfo, nil
}
