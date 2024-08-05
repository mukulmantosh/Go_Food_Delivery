package restaurant

import (
	"Go_Food_Delivery/pkg/database/models/restaurant"
	"context"
)

func (restSrv *RestaurantService) Add(ctx context.Context, restaurant *restaurant.Restaurant) (bool, error) {
	_, err := restSrv.db.NewInsert().Model(restaurant).Exec(ctx)
	if err != nil {
		return false, err
	}
	return true, nil
}
