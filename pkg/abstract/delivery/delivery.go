package delivery

import (
	"Go_Food_Delivery/pkg/database/models/delivery"
	"context"
)

type DeliveryPerson interface {
	AddDeliveryPerson(ctx context.Context, deliveryPerson *delivery.DeliveryPerson) (bool, error)
}

type Validation interface {
	GenerateTOTP(_ context.Context, phone string) (string, string, error)
	ValidateOTP(_ context.Context, secretKey string, otp string) bool
	ValidateAccountDetails(ctx context.Context, phone string) (*delivery.DeliveryPerson, error)
	Verify(ctx context.Context, phone string, otp string) bool
}

type DeliveryLogin interface {
	GenerateJWT(ctx context.Context, userId int64, name string) (string, error)
}
