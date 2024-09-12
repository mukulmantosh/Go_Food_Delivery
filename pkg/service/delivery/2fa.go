package delivery

import (
	"Go_Food_Delivery/cmd/api/middleware"
	"Go_Food_Delivery/pkg/database/models/delivery"
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/pquerna/otp/totp"
	"log/slog"
	"os"
	"time"
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

func (deliverSrv *DeliveryService) ValidateAccountDetails(ctx context.Context, phone string) (*delivery.DeliveryPerson, error) {
	var deliveryAccountInfo delivery.DeliveryPerson
	err := deliverSrv.db.Select(ctx, &deliveryAccountInfo, "phone", phone)
	if err != nil {
		return nil, err
	}
	if deliveryAccountInfo.Status != "AVAILABLE" {
		return nil, errors.New("account is inactive or not available")
	}
	return &deliveryAccountInfo, nil
}

func (deliverSrv *DeliveryService) ValidateOTP(_ context.Context, secretKey string, otp string) bool {
	return totp.Validate(otp, secretKey)
}

func (deliverSrv *DeliveryService) Verify(ctx context.Context, phone string, otp string) bool {
	accDetail, err := deliverSrv.ValidateAccountDetails(ctx, phone)
	if err != nil {
		slog.Error("Error::validating account details", "err", err)
		return false
	}

	valid := deliverSrv.ValidateOTP(ctx, accDetail.AuthKey, otp)
	return valid
}

func (deliverSrv *DeliveryService) GenerateJWT(_ context.Context, userId int64, name string) (string, error) {

	claims := middleware.UserClaims{UserID: userId, Name: name,
		RegisteredClaims: jwt.RegisteredClaims{

			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(2))),
			Issuer:    "Go_Food_Delivery",
		}}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
}
