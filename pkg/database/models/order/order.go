package order

import (
	"Go_Food_Delivery/pkg/database/models/restaurant"
	userModel "Go_Food_Delivery/pkg/database/models/user"
	"Go_Food_Delivery/pkg/database/models/utils"
	"github.com/uptrace/bun"
)

type Order struct {
	bun.BaseModel   `bun:"table:orders"`
	OrderID         int64   `bun:",pk,autoincrement" json:"order_id"`
	UserID          int64   `bun:"user_id,notnull" json:"-"`
	OrderStatus     string  `bun:"order_status,notnull" json:"order_status"`
	TotalAmount     float64 `bun:"total_amount,notnull" json:"total_amount"`
	DeliveryAddress string  `bun:"delivery_address,notnull" json:"delivery_address"`
	utils.Timestamp
	User *userModel.User `bun:"rel:belongs-to,join:user_id=id" json:"-"`
}

type OrderItems struct {
	bun.BaseModel `bun:"table:order_items"`
	OrderItemID   int64   `bun:",pk,autoincrement" json:"order_item_id"`
	OrderID       int64   `bun:"order_id,notnull" json:"order_id"`
	ItemID        int64   `bun:"item_id,notnull" json:"item_id"`
	RestaurantID  int64   `bun:"restaurant_id,notnull" json:"restaurant_id"`
	Quantity      int64   `bun:"quantity,notnull" json:"quantity"`
	Price         float64 `bun:"price,notnull" json:"price"`
	utils.Timestamp
	MenuItem   *restaurant.MenuItem   `bun:"rel:belongs-to,join:item_id=menu_id" json:"MenuItem"`
	Restaurant *restaurant.Restaurant `bun:"rel:belongs-to,join:restaurant_id=restaurant_id"`
	Order      *Order                 `bun:"rel:belongs-to,join:order_id=order_id" json:"-"`
}
