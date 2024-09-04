package cart

import (
	"Go_Food_Delivery/pkg/database/models/cart"
	"context"
)

type Cart interface {
	Create(ctx context.Context, cart *cart.Cart) (*cart.Cart, error)
	GetCartId(ctx context.Context, UserId int64) (*cart.Cart, error)
	AddItem(ctx context.Context, Item *cart.CartItems) (*cart.CartItems, error)
}

type CartItems interface {
	ListItems(ctx context.Context, cartId int64) (*[]cart.CartItems, error)
	DeleteItem(ctx context.Context, cartItemId int64) (bool, error)
}
