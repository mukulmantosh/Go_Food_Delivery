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
	Phone            string `bun:"phone,unique,notnull" json:"phone"`
	VehicleDetails   string `bun:"vehicle_details,notnull" json:"vehicle_details"`
	Status           string `bun:"status,notnull" json:"status"`
	AuthKey          string `bun:"auth_key,notnull" json:"auth_key"`
	AuthKeyURL       string `bun:"auth_key_url,notnull" json:"auth_key_url"`
	IsAuthSet        bool   `bun:"is_auth_set,notnull" json:"is_auth_set"`
	utils.Timestamp
}

type Deliveries struct {
	bun.BaseModel    `bun:"table:deliveries"`
	DeliveryID       int64           `bun:",pk,autoincrement" json:"delivery_id"`
	DeliveryPersonID int64           `bun:"delivery_person_id,notnull" json:"-"`
	OrderID          int64           `bun:"order_id,notnull" json:"order_id"`
	DeliveryStatus   string          `bun:"delivery_status,notnull" json:"delivery_status"`
	DeliveryTime     time.Time       `bun:",nullzero"`
	Order            *order.Order    `bun:"rel:belongs-to,join:order_id=order_id" json:"-"`
	DeliveryPerson   *DeliveryPerson `bun:"rel:belongs-to,join:delivery_person_id=delivery_person_id" json:"-"`
	utils.Timestamp
}

type DeliveryPersonParams struct {
	Name           string `json:"name"`
	Phone          string `json:"phone"`
	VehicleDetails string `json:"vehicle_details"`
}

type DeliveryLoginParams struct {
	Phone string `json:"phone"`
	OTP   string `json:"otp"`
}

type DeliveryOrderPlacementParams struct {
	OrderID int64  `json:"order_id"`
	Status  string `json:"status"`
}

type DeliveryListResponse struct {
	DeliveryId       int       `json:"delivery_id"`
	DeliveryPersonId int       `json:"delivery_person_id"`
	DeliveryStatus   string    `json:"delivery_status"`
	DeliveryTime     time.Time `json:"delivery_time"`
	CreatedAt        time.Time `json:"created_at"`
	Name             string    `json:"name"`
	VehicleDetails   string    `json:"vehicle_details"`
	Phone            string    `json:"phone"`
}
