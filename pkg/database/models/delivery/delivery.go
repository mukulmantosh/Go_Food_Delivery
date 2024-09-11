package delivery

import (
	"Go_Food_Delivery/pkg/database/models/utils"
	"github.com/uptrace/bun"
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
