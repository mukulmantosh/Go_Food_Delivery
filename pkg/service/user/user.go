package database

import (
	"Go_Food_Delivery/pkg/database"
	"Go_Food_Delivery/pkg/database/models/user"
	"context"
	"errors"
	"github.com/uptrace/bun"
	"log/slog"
)

type UserService struct {
	db *bun.DB
}

func NewUserService(db database.Database) *UserService {
	return &UserService{db: db.Db()}
}

func (userService *UserService) accountExists(ctx context.Context, email string) (bool, error) {
	var count int
	err := userService.db.NewSelect().Model((*user.User)(nil)).
		ColumnExpr("COUNT(*)").Where("email = ?", email).
		Scan(ctx, &count)

	if err != nil {
		slog.Info("UserService.accountExists: %v", err)
		return false, err
	}
	return count > 0, nil
}

func (userService *UserService) Add(ctx context.Context, user *user.User) (bool, error) {
	accountExists, err := userService.accountExists(ctx, user.Email)
	if accountExists {
		return false, errors.New("the user you are trying to register already exists")
	} else {
		_, err = userService.db.NewInsert().Model(user).Exec(ctx)
		if err != nil {
			return false, err
		}
		return true, nil
	}

}
