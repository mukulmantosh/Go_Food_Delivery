package user

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/uptrace/bun"
	"golang.org/x/crypto/bcrypt"
	"log"
	"regexp"
)

type User struct {
	bun.BaseModel `bun:"table:users"`
	ID            int64  `bun:",pk,autoincrement" json:"id"`
	Name          string `bun:",notnull" json:"name" validate:"name"`
	Email         string `bun:",unique,notnull" json:"email" validate:"email"`
	Password      string `bun:",notnull" json:"password"`
}

type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *User) HashPassword() {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal("error hashing password")
	}
	u.Password = string(hashedPassword)
}

func (l *LoginUser) CheckPassword(hashPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(l.Password))
}

func NameValidator(fl validator.FieldLevel) bool {
	str, ok := fl.Field().Interface().(string)
	return ok && str != ""
}

func EmailValidator(fl validator.FieldLevel) bool {
	email, ok := fl.Field().Interface().(string)
	if !ok {
		return false
	}
	// Basic email regex pattern
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

func UserValidationError(err error) map[string]string {
	var validationErrors validator.ValidationErrors
	if !errors.As(err, &validationErrors) {
		return map[string]string{"error": "Unknown error"}
	}

	errorsMap := make(map[string]string)
	for _, e := range validationErrors {
		field := e.Field()
		switch e.Tag() {
		case "name":
			errorsMap[field] = "Provide your full name"
		case "email":
			errorsMap[field] = "Provide valid email address"
		default:
			errorsMap[field] = "Invalid"
		}
	}
	return errorsMap
}
