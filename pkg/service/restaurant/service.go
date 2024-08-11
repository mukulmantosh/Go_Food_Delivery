package restaurant

import (
	"Go_Food_Delivery/pkg/database"
)

type RestaurantService struct {
	db  database.Database
	env string
}

func NewRestaurantService(db database.Database, env string) *RestaurantService {
	return &RestaurantService{db, env}
}
