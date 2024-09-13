package cart

import (
	"Go_Food_Delivery/pkg/database/models/delivery"
	"Go_Food_Delivery/pkg/database/models/order"
	"context"
)

type Order interface {
	PlaceOrder(ctx context.Context, cartId int64) (*order.Order, error)
	OrderList(ctx context.Context, userId int64) (*[]order.Order, error)
	DeliveryInformation(ctx context.Context, orderId int64) (*[]delivery.Deliveries, error)
}
