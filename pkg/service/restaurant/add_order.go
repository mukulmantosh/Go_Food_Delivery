package restaurant

import (
	"Go_Food_Delivery/pkg/database/models/order"
	"context"
)

func (restSrv *RestaurantService) AddOrder(ctx context.Context, order *order.Order) (*order.Order, error) {
	_, err := restSrv.db.Insert(ctx, order)
	if err != nil {
		return nil, err
	}
	return order, nil
}
