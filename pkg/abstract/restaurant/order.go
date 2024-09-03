package restaurant

import (
	"Go_Food_Delivery/pkg/database/models/order"
	"context"
)

type Order interface {
	AddOrder(ctx context.Context, order *order.Order) (*order.Order, error)
}
