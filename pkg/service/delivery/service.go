package delivery

import (
	"Go_Food_Delivery/pkg/database"
	"Go_Food_Delivery/pkg/nats"
)

type DeliveryService struct {
	db   database.Database
	env  string
	nats *nats.NATS
}

func NewDeliveryService(db database.Database, env string, nats *nats.NATS) *DeliveryService {
	return &DeliveryService{db, env, nats}
}
