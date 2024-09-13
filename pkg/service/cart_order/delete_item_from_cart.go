package cart_order

import (
	"Go_Food_Delivery/pkg/database"
	"context"
)

func (cartSrv *CartService) DeleteItem(ctx context.Context, cartItemId int64) (bool, error) {
	filter := database.Filter{"cart_item_id": cartItemId}

	_, err := cartSrv.db.Delete(ctx, "cart_items", filter)
	if err != nil {
		return false, err
	}
	return true, nil
}
