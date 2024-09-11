package delivery

import "Go_Food_Delivery/pkg/database"

type DeliveryService struct {
	db  database.Database
	env string
}

func NewDeliveryService(db database.Database, env string) *DeliveryService {
	return &DeliveryService{db, env}
}
