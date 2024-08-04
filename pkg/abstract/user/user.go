package user

import (
	"Go_Food_Delivery/pkg/database/models/user"
	"context"
)

type User interface {
	Add(ctx context.Context, user *user.User) (bool, error)
	Delete(ctx context.Context, userId int64) (bool, error)
}
