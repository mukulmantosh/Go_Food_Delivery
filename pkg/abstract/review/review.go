package review

import (
	"Go_Food_Delivery/pkg/database/models/review"
	"context"
)

type Review interface {
	Add(ctx context.Context, review *review.Review) (bool, error)
}
