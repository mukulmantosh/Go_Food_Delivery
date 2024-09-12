package delivery

import (
	"Go_Food_Delivery/pkg/database/models/delivery"
	"context"
)

type DeliveryPerson interface {
	AddDeliveryPerson(ctx context.Context, deliveryPerson *delivery.DeliveryPerson) (bool, error)
	GenerateTOTP(_ context.Context, phone string) (string, string, error)
}
