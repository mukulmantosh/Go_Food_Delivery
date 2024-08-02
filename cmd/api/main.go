package main

import (
	"Uber_Food_Delivery/internal/database"
	"Uber_Food_Delivery/internal/handler"
	registration "Uber_Food_Delivery/internal/handler/register"
	"log"
)

func main() {
	db := database.New()
	s := handler.NewServer(db)

	registration.NewRegister(s, "/register")

	log.Fatal(s.Run())

}
