package cart_order

import (
	"Go_Food_Delivery/pkg/database/models/cart"
	"context"
)

func (cartSrv *CartService) AddItem(ctx context.Context, Item *cart.CartItems) (*cart.CartItems, error) {
	_, err := cartSrv.db.Insert(ctx, Item)
	if err != nil {
		return nil, err
	}
	return Item, nil
}
