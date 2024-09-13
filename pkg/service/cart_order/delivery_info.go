package cart_order

import (
	"Go_Food_Delivery/pkg/database/models/delivery"
	"context"
)

func (cartSrv *CartService) DeliveryInformation(ctx context.Context, orderId int64, userId int64) (*[]delivery.DeliveryListResponse, error) {
	var deliveryList []delivery.DeliveryListResponse

	err := cartSrv.db.Raw(ctx, &deliveryList, `
	SELECT d.delivery_id, d.delivery_person_id, 
       d.delivery_status, d.delivery_time, d.created_at,
       dp.name, dp.vehicle_details, dp.phone
	FROM deliveries AS d
         JOIN orders AS o ON d.order_id = o.order_id
         JOIN delivery_person AS dp ON d.delivery_person_id = dp.delivery_person_id
	WHERE o.user_id = ? AND d.order_id = ?;`,
		userId, orderId)
	if err != nil {
		return nil, err
	}
	return &deliveryList, nil
}
