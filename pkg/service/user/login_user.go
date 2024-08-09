package user

import (
	"Go_Food_Delivery/pkg/database/models/user"
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"log/slog"
	"os"
	"time"
)

type Claims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func (usrSrv *UsrService) Login(ctx context.Context, user *user.LoginUser) (string, error) {

	claims := Claims{Email: user.Email,
		RegisteredClaims: jwt.RegisteredClaims{

			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(1))),
			Issuer:    "Go_Food_Delivery",
		}}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
}

func (usrSrv *UsrService) UserExist(ctx context.Context, email string) (bool, error) {
	count, err := usrSrv.db.Count(ctx, "users", "COUNT(*)", "email", email)
	if err != nil {
		slog.Info("UserService.UserExist::Error %v", err)
		return false, err
	}
	return count > 0, nil
}

func (usrSrv *UsrService) ValidatePassword(ctx context.Context, userInput *user.LoginUser) (bool, error) {
	var userAccount user.User
	err := usrSrv.db.Select(ctx, &userAccount, "email", userInput.Email)
	if err != nil {
		slog.Info("UserService.ValidatePassword::Error %v", err)
		return false, err
	}

	err = userInput.CheckPassword(userAccount.Password)
	if err != nil {
		return false, errors.New("invalid password")
	}
	return true, nil

}
