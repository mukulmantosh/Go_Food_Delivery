package delivery

import (
	"Go_Food_Delivery/pkg/database/models/delivery"
	"context"
)

func (deliverSrv *DeliveryService) AddDeliveryPerson(ctx context.Context, deliveryPerson *delivery.DeliveryPerson) (*delivery.DeliveryPerson, error) {
	_, err := deliverSrv.db.Insert(ctx, deliveryPerson)
	if err != nil {
		return nil, err
	}
	return deliveryPerson, nil
}
