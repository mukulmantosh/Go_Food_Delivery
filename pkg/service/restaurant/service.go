package restaurant

import (
	"Go_Food_Delivery/pkg/database"
	"github.com/uptrace/bun"
)

type RestaurantService struct {
	db  *bun.DB
	Env string
}

func NewRestaurantService(db database.Database, Environment string) *RestaurantService {
	return &RestaurantService{db: db.Db(), Env: Environment}
}
