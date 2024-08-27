package user

import (
	"Go_Food_Delivery/pkg/database/models/user"
	"context"
	"errors"
)

func ValidateAccount(login func(ctx context.Context, userID int64, userName string) (string, error),
	accountExists func(ctx context.Context, email string, recordRequired bool) (bool, int64, string, error),
	validatePassword func(ctx context.Context, user *user.LoginUser) (bool, error)) func(ctx context.Context, user *user.LoginUser) (string, error) {
	return func(ctx context.Context, user *user.LoginUser) (string, error) {
		exists, userId, name, err := accountExists(ctx, user.Email, true)
		if err != nil {
			return "", err
		}
		if !exists {
			return "", errors.New("we did not find any account with this user")
		}

		_, err = validatePassword(ctx, user)
		if err != nil {
			return "", err
		}

		return login(ctx, userId, name)
	}
}
