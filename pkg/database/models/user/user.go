package user

import (
	"github.com/uptrace/bun"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type User struct {
	bun.BaseModel `bun:"table:users"`
	ID            int64  `bun:",pk,autoincrement" json:"id"`
	Name          string `bun:",notnull" json:"name"`
	Email         string `bun:",unique,notnull" json:"email"`
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
