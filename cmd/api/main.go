package main

import (
	"Uber_Food_Delivery/internal/database"
	"Uber_Food_Delivery/internal/server"
	registration "Uber_Food_Delivery/internal/server/register"
	"log"
)

func main() {
	db := database.New()
	s := server.NewServer(db)

	registration.NewRegister(s, "/register")

	log.Fatal(s.Run())

}
