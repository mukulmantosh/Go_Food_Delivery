package review

import (
	"Go_Food_Delivery/pkg/database/models/restaurant"
	"Go_Food_Delivery/pkg/database/models/user"
	"Go_Food_Delivery/pkg/database/models/utils"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/uptrace/bun"
)

type Review struct {
	bun.BaseModel `bun:"table:reviews"`

	ReviewID     int64  `bun:"review_id,pk,autoincrement" json:"review_id"`
	UserID       int64  `bun:"user_id" json:"user_id"`
	RestaurantID int64  `bun:"restaurant_id" json:"restaurant_id"`
	Rating       int    `bun:"rating"`
	Comment      string `bun:"comment"`
	utils.Timestamp

	User       *user.User             `bun:"rel:belongs-to,join:user_id=id" json:"-"`
	Restaurant *restaurant.Restaurant `bun:"rel:belongs-to,join:restaurant_id=restaurant_id" json:"-"`
}

type ReviewParams struct {
	Rating  int    `json:"rating" validate:"rating"`
	Comment string `json:"comment"`
}

func RatingValidator(fl validator.FieldLevel) bool {
	rating, ok := fl.Field().Interface().(int)
	return ok && rating >= 1 && rating <= 5
}

func ReviewValidationError(err error) map[string]string {
	var validationErrors validator.ValidationErrors
	if !errors.As(err, &validationErrors) {
		return map[string]string{"error": "Unknown error"}
	}

	errorsMap := make(map[string]string)
	for _, e := range validationErrors {
		field := e.Field()
		switch e.Tag() {
		case "rating":
			errorsMap[field] = "Rating must be between 1 and 5"
		default:
			errorsMap[field] = "Invalid"
		}
	}
	return errorsMap
}
