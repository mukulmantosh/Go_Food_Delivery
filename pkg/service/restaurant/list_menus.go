package restaurant

import (
	restaurantModel "Go_Food_Delivery/pkg/database/models/restaurant"
	"context"
)

func (restSrv *RestaurantService) ListMenus(ctx context.Context, restaurantId int64) ([]restaurantModel.MenuItem, error) {
	var menuItems []restaurantModel.MenuItem

	if err := restSrv.db.NewSelect().
		Table("menu_item").Where("restaurant_id = ?", restaurantId).
		Scan(ctx, &menuItems); err != nil {
		return menuItems, err
	}
	return menuItems, nil
}
