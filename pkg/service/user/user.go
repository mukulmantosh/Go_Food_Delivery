package database

import (
	"Uber_Food_Delivery/pkg/database"
	"Uber_Food_Delivery/pkg/database/models/user"
	"context"
	"fmt"
	"github.com/uptrace/bun"
	"time"
)

type UserService struct {
	db *bun.DB
}

func NewUserService(db database.Database) *UserService {
	return &UserService{db: db.Db()}
}

func (userService *UserService) Add(ctx context.Context, user *user.User) (bool, error) {
	time.Sleep(10 * time.Second) // This simulates a long processing delay

	// Perform the actual database operation
	userInfo, err := userService.db.NewInsert().Model(user).Exec(ctx)
	fmt.Println(userInfo, err)
	return true, err
}
