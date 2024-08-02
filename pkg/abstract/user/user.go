package user

import (
	"Uber_Food_Delivery/pkg/database/models/user"
	"context"
)

type User interface {
	Add(ctx context.Context, user *user.User) (bool, error)
	//GetUserById(ctx context.Context, ID string) (*models.DisplayUser, error)
	//UpdateUser(ctx context.Context, user *models.User) (bool, error)
	//DeleteUser(ctx context.Context, ID string) error
}
