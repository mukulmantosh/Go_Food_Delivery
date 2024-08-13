package review

import (
	"Go_Food_Delivery/pkg/database/models/review"
	"context"
)

func (revSrv *ReviewService) Add(ctx context.Context, review *review.Review) (bool, error) {
	_, err := revSrv.db.Insert(ctx, review)
	if err != nil {
		return false, err
	}
	return true, nil
}
