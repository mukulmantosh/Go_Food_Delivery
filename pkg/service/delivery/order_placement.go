package delivery

import (
	"Go_Food_Delivery/pkg/database"
	"Go_Food_Delivery/pkg/database/models/delivery"
	"Go_Food_Delivery/pkg/database/models/order"
	"context"
	"errors"
	"fmt"
	"time"
)

func (deliverSrv *DeliveryService) updateOrderStatus(ctx context.Context, orderID int64, status string) error {
	_, err := deliverSrv.db.Update(ctx, "orders", database.Filter{"order_status": status},
		database.Filter{"order_id": orderID})
	if err != nil {
		return err
	}
	_, err = deliverSrv.db.Update(ctx, "deliveries", database.Filter{"delivery_status": status},
		database.Filter{"order_id": orderID})
	return nil
}

func (deliverSrv *DeliveryService) OrderPlacement(ctx context.Context,
	deliveryPersonID int64, orderID int64, deliveryStatus string) (bool, error) {
	var orderInfo order.Order
	setFilter := database.Filter{"order_status": deliveryStatus}
	whereFilter := database.Filter{"order_id": orderID}

	// Check the order is valid or not.
	err := deliverSrv.db.Select(ctx, &orderInfo, "order_id", orderID)
	if err != nil {
		return false, err
	}

	// Perform generic validation.
	_, err = deliverSrv.orderValidation(ctx, &orderInfo, deliveryPersonID)
	if err != nil {
		return false, err
	}

	invalidStatuses := map[string]bool{
		"cancelled": true,
		"completed": true,
		"failed":    true,
		"delivered": true,
	}

	if invalidStatuses[orderInfo.OrderStatus] {
		return false, errors.New("this order is invalid or it has been already delivered,failed or cancelled")
	}

	switch orderInfo.OrderStatus {
	case "in_progress", "on_the_way":
		// Update Order
		_, err := deliverSrv.db.Update(ctx, "orders", setFilter, whereFilter)
		if err != nil {
			return false, err
		}

		// Add info. Into the delivery table.
		deliver := new(delivery.Deliveries)
		deliver.DeliveryPersonID = deliveryPersonID
		deliver.OrderID = orderID
		deliver.DeliveryStatus = deliveryStatus

		if deliveryStatus == "delivered" {
			deliver.DeliveryTime = time.Now()

		}

		_, err = deliverSrv.db.Insert(ctx, deliver)
		if err != nil {
			return false, err
		}
		// Notify User.
		err = deliverSrv.notifyDeliveryStatusToUser(&orderInfo, deliveryStatus)
		if err != nil {
			return false, err
		}
		return true, nil
	default:
		return false, errors.New("unknown order status")
	}

}

func (deliverSrv *DeliveryService) notifyDeliveryStatusToUser(order *order.Order, status string) error {
	var message string

	switch status {
	case "delivered":
		message = fmt.Sprintf("USER_ID:%d|MESSAGE:Your order no.%d has been successfully %s", order.UserID, order.OrderID, status)
	case "failed":
		message = fmt.Sprintf("USER_ID:%d|MESSAGE:Your order no.%d has been %s", order.UserID, order.OrderID, status)
	case "cancelled":
		message = fmt.Sprintf("USER_ID:%d|MESSAGE:Your order no.%d has been %s", order.UserID, order.OrderID, status)
	case "on_the_way":
		message = fmt.Sprintf("USER_ID:%d|MESSAGE:Your order no.%d is %s", order.UserID, order.OrderID, status)
	default:
		return fmt.Errorf("invalid status: %s", status)
	}

	topic := fmt.Sprintf("orders.status.%d", order.OrderID)
	err := deliverSrv.nats.Pub(topic, []byte(message))
	if err != nil {
		return err
	}
	return nil
}
