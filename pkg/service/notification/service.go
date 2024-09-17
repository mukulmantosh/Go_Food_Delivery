package notification

import (
	"Go_Food_Delivery/pkg/database"
	"Go_Food_Delivery/pkg/nats"
	"log/slog"
)

type NotificationService struct {
	db   database.Database
	env  string
	nats *nats.NATS
}

func NewNotificationService(db database.Database, env string, nats *nats.NATS) *NotificationService {
	return &NotificationService{db, env, nats}
}

func (s *NotificationService) SubscribeNewOrders(ordersMessage *chan string) error {
	slog.Info("NotificationService::SubscribeNewOrders")
	err := s.nats.Sub("orders.new.*", ordersMessage)
	if err != nil {
		return err
	}
	return nil
}
