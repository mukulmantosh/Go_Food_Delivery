package restaurant

import (
	"Go_Food_Delivery/pkg/database/models/restaurant"
	"context"
)

type MenuItems interface {
	AddMenu(ctx context.Context, menu *restaurant.MenuItem) (bool, error, int64, string)
	UpdateMenuPhoto(ctx context.Context, menuID int64, imageURL string)
}
