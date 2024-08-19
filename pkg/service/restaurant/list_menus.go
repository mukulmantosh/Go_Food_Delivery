package restaurant

import (
	restaurantModel "Go_Food_Delivery/pkg/database/models/restaurant"
	"context"
)

func (restSrv *RestaurantService) ListMenus(ctx context.Context, restaurantId int64) ([]restaurantModel.MenuItem, error) {
	var menuItems []restaurantModel.MenuItem

	err := restSrv.db.Select(ctx, &menuItems, "restaurant_id", restaurantId)
	if err != nil {
		return nil, err
	}
	return menuItems, nil
}

func (restSrv *RestaurantService) ListAllMenus(ctx context.Context) ([]restaurantModel.MenuItem, error) {
	var menuItems []restaurantModel.MenuItem

	err := restSrv.db.SelectAll(ctx, "menu_item", &menuItems)
	if err != nil {
		return nil, err
	}
	return menuItems, nil
}
