package restaurant

import (
	"Go_Food_Delivery/pkg/database/models/restaurant"
	"context"
)

type Restaurant interface {
	Add(ctx context.Context, user *restaurant.Restaurant) (bool, error)
	ListRestaurants(ctx context.Context) ([]restaurant.Restaurant, error)
	ListRestaurantById(ctx context.Context, restaurantId int64) (restaurant.Restaurant, error)
	DeleteRestaurant(ctx context.Context, restaurantId int64) (bool, error)
}
