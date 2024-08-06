package restaurant

import (
	"context"
)

func (restSrv *RestaurantService) DeleteRestaurant(ctx context.Context, restaurantId int64) (bool, error) {

	_, err := restSrv.db.NewDelete().Table("restaurant").Where("restaurant_id = ?", restaurantId).Exec(ctx)

	if err != nil {
		return false, err
	}
	return true, nil
}
