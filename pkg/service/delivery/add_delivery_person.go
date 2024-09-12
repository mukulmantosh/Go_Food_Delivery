package delivery

import (
	"Go_Food_Delivery/pkg/database/models/delivery"
	"context"
)

func (deliverSrv *DeliveryService) AddDeliveryPerson(ctx context.Context, deliveryPerson *delivery.DeliveryPerson) (bool, error) {
	_, err := deliverSrv.db.Insert(ctx, deliveryPerson)
	if err != nil {
		return false, err
	}
	return true, nil
}
