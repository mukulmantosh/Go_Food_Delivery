package user

import (
	"context"
)

func (usrSrv *UsrService) Delete(ctx context.Context, userId int64) (bool, error) {
	_, err := usrSrv.db.Delete(ctx, "users", "id", userId)
	if err != nil {
		return false, err
	}
	return true, nil
}
