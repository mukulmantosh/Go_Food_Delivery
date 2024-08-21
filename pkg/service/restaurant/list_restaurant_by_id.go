package restaurant

import (
	"Go_Food_Delivery/pkg/database/models/restaurant"
	"context"
)

func (restSrv *RestaurantService) ListRestaurantById(ctx context.Context, restaurantId int64) (restaurant.Restaurant, error) {
	var restro restaurant.Restaurant

	err := restSrv.db.Select(ctx, &restro, "restaurant_id", restaurantId)
	if err != nil {
		return restaurant.Restaurant{}, err
	}
	return restro, nil
}
