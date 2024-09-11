package delivery

import (
	"Go_Food_Delivery/pkg/database/models/order"
	"Go_Food_Delivery/pkg/database/models/utils"
	"github.com/uptrace/bun"
	"time"
)

type DeliveryPerson struct {
	bun.BaseModel    `bun:"table:delivery_person"`
	DeliveryPersonID int64  `bun:",pk,autoincrement" json:"delivery_person_id"`
	Name             string `bun:"name,notnull" json:"name"`
	Phone            string `bun:"phone,notnull" json:"phone"`
	VehicleDetails   string `bun:"vehicle_details,notnull" json:"vehicle_details"`
	Status           string `bun:"status,notnull" json:"status"`
	utils.Timestamp
}

type Deliveries struct {
	bun.BaseModel    `bun:"table:deliveries"`
	DeliveryID       int64        `bun:",pk,autoincrement" json:"delivery_id"`
	DeliveryPersonID int64        `bun:"delivery_person_id,notnull" json:"delivery_person_id"`
	OrderID          int64        `bun:"order_id,notnull" json:"order_id"`
	DeliveryStatus   string       `bun:"delivery_status,notnull" json:"delivery_status"`
	DeliveryTime     time.Time    `bun:",nullzero"`
	Order            *order.Order `bun:"rel:belongs-to,join:order_id=order_id" json:"-"`
	utils.Timestamp
}
