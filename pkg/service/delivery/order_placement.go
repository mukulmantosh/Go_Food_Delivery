package delivery

import (
	"Go_Food_Delivery/pkg/database"
	"Go_Food_Delivery/pkg/database/models/delivery"
	"Go_Food_Delivery/pkg/database/models/order"
	"context"
	"errors"
	"time"
)

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

		// Add info. into the delivery table.
		deliver := new(delivery.Deliveries)
		deliver.DeliveryPersonID = deliveryPersonID
		deliver.OrderID = orderID
		deliver.DeliveryStatus = deliveryStatus

		if orderInfo.OrderStatus == "on_the_way" {
			deliver.DeliveryTime = time.Now()

		}

		_, err = deliverSrv.db.Insert(ctx, deliver)
		if err != nil {
			return false, err
		}

		return true, nil
	default:
		return false, errors.New("unknown order status")
	}

}
