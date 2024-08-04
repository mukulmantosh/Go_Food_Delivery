package user

import (
	"Go_Food_Delivery/pkg/database/models/user"
	"context"
)

func (usrSrv *UsrService) Delete(ctx context.Context, userId int64) (bool, error) {
	_, err := usrSrv.db.NewDelete().Model((*user.User)(nil)).Where("id = ?", userId).Exec(ctx)
	if err != nil {
		return false, err
	}
	return true, nil
}
