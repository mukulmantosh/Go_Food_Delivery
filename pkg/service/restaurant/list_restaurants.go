package restaurant

import (
	restaurantModel "Go_Food_Delivery/pkg/database/models/restaurant"
	"context"
)

func (restSrv *RestaurantService) ListRestaurants(ctx context.Context) ([]restaurantModel.Restaurant, error) {
	var restaurants []restaurantModel.Restaurant

	if err := restSrv.db.NewSelect().Table("restaurant").Scan(ctx, &restaurants); err != nil {
		return restaurants, err
	}

	return restaurants, nil
}
