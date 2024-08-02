package main

import (
	"Uber_Food_Delivery/pkg/database"
	"Uber_Food_Delivery/pkg/handler"
	registration "Uber_Food_Delivery/pkg/handler/register"
	"log"
)

func main() {
	db := database.New()
	s := handler.NewServer(db)

	registration.NewRegister(s, "/register")

	log.Fatal(s.Run())

}
