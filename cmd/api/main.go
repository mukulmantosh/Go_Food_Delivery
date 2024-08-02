package main

import (
	"Uber_Food_Delivery/pkg/database"
	"Uber_Food_Delivery/pkg/handler"
	registration "Uber_Food_Delivery/pkg/handler/register"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	// load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := database.New()
	// Create Tables
	if err := db.Migrate(); err != nil {
		log.Fatalf("Error migrating database: %s", err)
	}

	s := handler.NewServer(db)

	registration.NewRegister(s, "/register")

	log.Fatal(s.Run())

}
