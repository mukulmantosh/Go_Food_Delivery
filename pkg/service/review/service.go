package review

import (
	"Go_Food_Delivery/pkg/database"
)

type ReviewService struct {
	db  database.Database
	env string
}

func NewReviewService(db database.Database, env string) *ReviewService {
	return &ReviewService{db, env}
}
