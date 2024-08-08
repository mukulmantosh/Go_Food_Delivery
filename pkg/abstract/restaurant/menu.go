package restaurant

import (
	"Go_Food_Delivery/pkg/database/models/restaurant"
	"context"
)

type MenuItems interface {
	AddMenu(ctx context.Context, menu *restaurant.MenuItem) (*restaurant.MenuItem, error)
	UpdateMenuPhoto(ctx context.Context, menu *restaurant.MenuItem)
}
