package user

import (
	"Go_Food_Delivery/pkg/database"
)

type UsrService struct {
	db  database.Database
	env string
}

func NewUserService(db database.Database, env string) *UsrService {
	return &UsrService{db, env}
}
