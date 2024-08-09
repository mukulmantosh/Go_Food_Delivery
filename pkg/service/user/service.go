package user

import (
	"Go_Food_Delivery/pkg/database"
)

type UsrService struct {
	db database.Database
}

func NewUserService(db database.Database) *UsrService {
	return &UsrService{db: db}
}
