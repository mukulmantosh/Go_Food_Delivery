package restaurant

import (
	"Go_Food_Delivery/pkg/database"
	"context"
)

func (restSrv *RestaurantService) DeleteRestaurant(ctx context.Context, restaurantId int64) (bool, error) {
	filter := database.Filter{"restaurant_id": restaurantId}

	_, err := restSrv.db.Delete(ctx, "restaurant", filter)
	if err != nil {
		return false, err
	}
	return true, nil
}
