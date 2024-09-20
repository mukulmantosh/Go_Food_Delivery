package delivery

import (
	"Go_Food_Delivery/pkg/database/models/order"
	"context"
	"errors"
	"time"
)

func (deliverSrv *DeliveryService) orderValidation(ctx context.Context, order *order.Order, deliveryPersonID int64) (bool, error) {
	orderPlacedTime := order.CreatedAt
	currentTime := time.Now()

	var deliveryExists int64
	var deliveryCancelCount int64
	var totalDeliveryCancellationCount int64
	err := deliverSrv.db.Raw(ctx, &deliveryExists,
		`SELECT COUNT(*) FROM deliveries WHERE order_id=? AND delivery_person_id=? AND delivery_status='on_the_way';`,
		order.OrderID, deliveryPersonID)
	if err != nil {
		return false, err
	}

	err = deliverSrv.db.Raw(ctx, &deliveryCancelCount,
		`SELECT COUNT(*) FROM deliveries WHERE order_id=? AND delivery_person_id=? AND delivery_status='cancelled';`,
		order.OrderID, deliveryPersonID)
	if err != nil {
		return false, err
	}

	err = deliverSrv.db.Raw(ctx, &totalDeliveryCancellationCount,
		`SELECT count(*) FROM deliveries WHERE delivery_person_id = ? 
                                  AND delivery_status = 'cancelled' 
                                  AND created_at >= now() - interval '1 hour';`, deliveryPersonID)
	if err != nil {
		return false, err
	}

	// If the order remains unclaimed by any delivery partner for more than 5 minutes, it will be canceled.
	if order.OrderStatus == "in_progress" {
		if currentTime.Sub(orderPlacedTime) > 5*time.Minute {
			_ = deliverSrv.updateOrderStatus(ctx, order.OrderID, "cancelled")
			return false, errors.New("this order was not accepted by the restaurant " +
				"for more than 5 minutes. It has been cancelled")
		}
	}

	//If a delivery partner accepts an order, it should no longer be visible to other delivery partners.
	if order.OrderStatus == "on_the_way" && deliveryExists == 0 {
		return false, errors.New("this order was already accepted by another delivery partner")
	}
	// If a delivery partner cancels an order, it should no longer be visible to them but will remain visible to other delivery partners.
	if deliveryCancelCount == 1 {
		return false, errors.New("you won't be able to accept this order again. It has been cancelled")
	}

	// If a delivery partner cancels three or more orders within an hour of accepting them, they will be blocked from receiving orders for three hours.

	return true, nil
}
