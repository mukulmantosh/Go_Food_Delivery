package user

import (
	"Go_Food_Delivery/pkg/database/models/user"
	"context"
	"log/slog"
)

func (usrSrv *UsrService) accountExists(ctx context.Context, email string) (bool, error) {
	var count int
	err := usrSrv.db.NewSelect().Model((*user.User)(nil)).
		ColumnExpr("COUNT(*)").Where("email = ?", email).
		Scan(ctx, &count)

	if err != nil {
		slog.Info("UserService.accountExists: %v", err)
		return false, err
	}
	return count > 0, nil
}
