package delivery

import (
	"context"
	"errors"
	"github.com/pquerna/otp/totp"
)

func (deliverSrv *DeliveryService) GenerateTOTP(_ context.Context, phone string) (string, string, error) {
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "Food Delivery",
		AccountName: phone,
	})
	if err != nil {
		return "", "", errors.New("error generating key")
	}
	return key.Secret(), key.URL(), nil
}
