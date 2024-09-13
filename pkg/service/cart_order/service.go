package cart_order

import "Go_Food_Delivery/pkg/database"

type CartService struct {
	db  database.Database
	env string
}

func NewCartService(db database.Database, env string) *CartService {
	return &CartService{db, env}
}
