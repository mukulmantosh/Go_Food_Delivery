package restaurant

import (
	"Go_Food_Delivery/pkg/database/models/utils"
	"github.com/uptrace/bun"
)

type Restaurant struct {
	bun.BaseModel `bun:"table:restaurant"`
	ID            int64  `bun:",pk,autoincrement" json:"id"`
	Name          string `bun:",notnull" json:"name"`
	Photo         string `bun:"store_image,nullzero" json:"store_image"`
	Description   string `bun:",notnull" json:"description"`
	Address       string `bun:"address,notnull" json:"address"`
	City          string `bun:"city,notnull" json:"city"`
	State         string `bun:"state,notnull" json:"state"`
	utils.Timestamp
}
