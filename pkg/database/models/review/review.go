package review

import (
	"Go_Food_Delivery/pkg/database/models/restaurant"
	"Go_Food_Delivery/pkg/database/models/user"
	"Go_Food_Delivery/pkg/database/models/utils"
	"github.com/uptrace/bun"
)

type Review struct {
	bun.BaseModel `bun:"table:reviews"`

	ReviewID     int    `bun:"review_id,pk,autoincrement"`
	UserID       int    `bun:"user_id"`
	RestaurantID int    `bun:"restaurant_id"`
	Rating       int    `bun:"rating"`
	Comment      string `bun:"comment"`
	utils.Timestamp

	User       *user.User             `bun:"rel:belongs-to,join:user_id=id" json:"-"`
	Restaurant *restaurant.Restaurant `bun:"rel:belongs-to,join:restaurant_id=restaurant_id" json:"-"`
}
