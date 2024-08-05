package user

import (
	"Go_Food_Delivery/pkg/database/models/user/utils"
	"github.com/uptrace/bun"
	"log"
)

type User struct {
	bun.BaseModel `bun:"table:users"`
	ID            int64  `bun:",pk,autoincrement" json:"id"`
	Name          string `bun:",notnull" json:"name"`
	Email         string `bun:",unique,notnull" json:"email"`
	Password      string `bun:",notnull" json:"password"`
}

func (u *User) HashPassword() {
	salt, err := utils.GenerateSalt()
	if err != nil {
		log.Fatal(err)
	}
	passwordHash := utils.Hash(u.Password, salt)
	u.Password = passwordHash
	return
}
