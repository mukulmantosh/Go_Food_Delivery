package restaurant

import (
	restaurantModel "Go_Food_Delivery/pkg/database/models/restaurant"
	"context"
)

func (restSrv *RestaurantService) ListRestaurants(ctx context.Context) ([]restaurantModel.Restaurant, error) {
	var restaurants []restaurantModel.Restaurant

	err := restSrv.db.SelectAll(ctx, "restaurant", &restaurants)
	if err != nil {
		return nil, err
	}

	return restaurants, nil
}
