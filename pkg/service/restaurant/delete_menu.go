package restaurant

import (
	"context"
)

func (restSrv *RestaurantService) DeleteMenu(ctx context.Context, menuId int64, restaurantId int64) (bool, error) {
	_, err := restSrv.db.NewDelete().Table("menu_item").
		Where("menu_id = ?", menuId).
		Where("restaurant_id = ?", restaurantId).
		Exec(ctx)
	if err != nil {
		return false, err
	}
	return true, nil
}
