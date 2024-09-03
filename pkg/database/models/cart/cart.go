package cart

import (
	"Go_Food_Delivery/pkg/database/models/restaurant"
	userModel "Go_Food_Delivery/pkg/database/models/user"
	"Go_Food_Delivery/pkg/database/models/utils"
	"github.com/uptrace/bun"
)

type Cart struct {
	bun.BaseModel `bun:"table:cart"`
	CartID        int64 `bun:",pk,autoincrement" json:"cart_id"`
	UserID        int64 `bun:"user_id,notnull" json:"user_id"`
	utils.Timestamp
	User *userModel.User `bun:"rel:belongs-to,join:user_id=id" json:"-"`
}

type CartItems struct {
	bun.BaseModel `bun:"table:cart_items"`
	CartItemID    int64 `bun:",pk,autoincrement" json:"cart_item_id"`
	CartID        int64 `bun:"cart_id,notnull" json:"cart_id"`
	ItemID        int64 `bun:"item_id,notnull" json:"item_id"`
	RestaurantID  int64 `bun:"restaurant_id,notnull" json:"restaurant_id"`
	Quantity      int64 `bun:"quantity,notnull" json:"quantity"`
	utils.Timestamp
	Restaurant *restaurant.Restaurant `bun:"rel:belongs-to,join:restaurant_id=restaurant_id" json:"-"`
	MenuItem   *restaurant.MenuItem   `bun:"rel:belongs-to,join:item_id=menu_id" json:"menu_item"`
	Cart       *Cart                  `bun:"rel:belongs-to,join:cart_id=cart_id" json:"-"`
}

type CartItemParams struct {
	CartID       int64 `json:"cart_id"`
	ItemID       int64 `json:"item_id"`
	RestaurantID int64 `json:"restaurant_id"`
	Quantity     int64 `json:"quantity"`
}
