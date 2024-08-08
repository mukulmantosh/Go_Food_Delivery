package user

import (
	"Go_Food_Delivery/pkg/database/models/user"
	"context"
	"errors"
)

func ValidateAccount(fn func(ctx context.Context, user *user.LoginUser) (string, error),
	accountExists func(ctx context.Context, email string) (bool, error),
	validatePassword func(ctx context.Context, user *user.LoginUser) (bool, error)) func(ctx context.Context, user *user.LoginUser) (string, error) {
	return func(ctx context.Context, user *user.LoginUser) (string, error) {
		exists, err := accountExists(ctx, user.Email)
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

		return fn(ctx, user)
	}
}
