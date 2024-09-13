package cart_order

import (
	"Go_Food_Delivery/pkg/database/models/cart"
	"context"
)

func (cartSrv *CartService) Create(ctx context.Context, cart *cart.Cart) (*cart.Cart, error) {
	_, err := cartSrv.db.Insert(ctx, cart)
	if err != nil {
		return nil, err
	}
	return cart, nil
}
