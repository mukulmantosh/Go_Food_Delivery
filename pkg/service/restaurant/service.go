package restaurant

import (
	"Go_Food_Delivery/pkg/database"
)

type RestaurantService struct {
	db  database.Database
	Env string
}

func NewRestaurantService(db database.Database, Environment string) *RestaurantService {
	return &RestaurantService{db: db, Env: Environment}
}
