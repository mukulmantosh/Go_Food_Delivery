package user

import (
	"Go_Food_Delivery/pkg/database/models/user"
	"context"
	"errors"
	"fmt"
)

func (usrSrv *UsrService) Add(ctx context.Context, user *user.User) (bool, error) {
	accountExists, _ := usrSrv.UserExist(ctx, user.Email)
	if accountExists {
		return false, errors.New("the user you are trying to register already exists")
	} else {
		user.HashPassword()
		_, err := usrSrv.db.Insert(ctx, user)
		fmt.Println("Inserted user", user.ID)
		if err != nil {
			return false, err
		}
		return true, nil
	}
}
