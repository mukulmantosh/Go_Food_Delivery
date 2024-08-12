package restaurant

import (
	"Go_Food_Delivery/pkg/database/models/utils"
	"github.com/uptrace/bun"
)

type Restaurant struct {
	bun.BaseModel `bun:"table:restaurant"`
	RestaurantID  int64  `bun:",pk,autoincrement" json:"restaurant_id"`
	Name          string `bun:",notnull" json:"name"`
	Photo         string `bun:"store_image,nullzero" json:"store_image"`
	Description   string `bun:",notnull" json:"description"`
	Address       string `bun:"address,notnull" json:"address"`
	City          string `bun:"city,notnull" json:"city"`
	State         string `bun:"state,notnull" json:"state"`
	utils.Timestamp
	MenuItems []MenuItem `bun:"rel:has-many,join:restaurant_id=menu_id" json:"-"`
}

type MenuItem struct {
	bun.BaseModel `bun:"table:menu_item"`
	MenuID        int64   `bun:",pk,autoincrement" json:"menu_id"`
	RestaurantID  int64   `bun:"restaurant_id,notnull" json:"restaurant_id"`
	Name          string  `bun:"name,notnull" json:"name"`
	Description   string  `bun:"description,notnull" json:"description"`
	Photo         string  `bun:"photo,nullzero" json:"photo"`
	Price         float64 `bun:"price,notnull" json:"price"`
	Category      string  `bun:"category,notnull" json:"category"`
	Available     bool    `bun:"available,default:True" json:"available"`
	utils.Timestamp
	Restaurant *Restaurant `bun:"rel:belongs-to,join:restaurant_id=restaurant_id" json:"-"`
}
