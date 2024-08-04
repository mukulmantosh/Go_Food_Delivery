package user

import (
	"Go_Food_Delivery/pkg/database"
	"github.com/uptrace/bun"
)

type UsrService struct {
	db *bun.DB
}

func NewUserService(db database.Database) *UsrService {
	return &UsrService{db: db.Db()}
}
