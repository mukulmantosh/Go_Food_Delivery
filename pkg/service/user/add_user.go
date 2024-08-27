package user

import (
	"Go_Food_Delivery/pkg/database/models/user"
	"context"
	"errors"
)

func (usrSrv *UsrService) Add(ctx context.Context, user *user.User) (bool, error) {
	accountExists, _, _, _ := usrSrv.UserExist(ctx, user.Email, false)
	if accountExists {
		return false, errors.New("the user you are trying to register already exists")
	} else {
		user.HashPassword()
		_, err := usrSrv.db.Insert(ctx, user)
		if err != nil {
			return false, err
		}
		return true, nil
	}
}
