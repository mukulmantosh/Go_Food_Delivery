package cart_order

import (
	"Go_Food_Delivery/pkg/database"
	"Go_Food_Delivery/pkg/nats"
)

type CartService struct {
	db   database.Database
	env  string
	nats *nats.NATS
}

func NewCartService(db database.Database, env string, nats *nats.NATS) *CartService {
	return &CartService{db, env, nats}
}
