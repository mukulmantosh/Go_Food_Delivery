package restaurant

import (
	"Go_Food_Delivery/pkg/database/models/restaurant"
	"context"
)

func (restSrv *RestaurantService) AddMenu(ctx context.Context, menu *restaurant.MenuItem) (bool, error) {
	_, err := restSrv.db.NewInsert().Model(menu).Exec(ctx)
	if err != nil {
		return false, err
	}
	return true, nil
}
