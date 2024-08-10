package restaurant

import (
	"Go_Food_Delivery/pkg/database"
	"context"
)

func (restSrv *RestaurantService) DeleteMenu(ctx context.Context, menuId int64, restaurantId int64) (bool, error) {
	filter := database.Filter{"menu_id": menuId, "restaurant_id": restaurantId}

	_, err := restSrv.db.Delete(ctx, "menu_item", filter)
	if err != nil {
		return false, err
	}
	return true, nil
}
