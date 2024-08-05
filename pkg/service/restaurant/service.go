package restaurant

import (
	"Go_Food_Delivery/pkg/database"
	"github.com/uptrace/bun"
)

type RestaurantService struct {
	db *bun.DB
}

func NewRestaurantService(db database.Database) *RestaurantService {
	return &RestaurantService{db: db.Db()}
}
