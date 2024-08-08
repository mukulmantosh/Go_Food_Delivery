package main

import (
	"Go_Food_Delivery/pkg/database"
	"Go_Food_Delivery/pkg/handler"
	registration "Go_Food_Delivery/pkg/handler/register"
	"Go_Food_Delivery/pkg/handler/restaurant"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	// load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	AppEnv := os.Getenv("APP_ENV")
	db := database.New()
	// Create Tables
	if err := db.Migrate(); err != nil {
		log.Fatalf("Error migrating database: %s", err)
	}

	s := handler.NewServer(db)

	registration.NewRegister(s, "/register")
	restaurant.NewRestaurant(s, "/restaurant", AppEnv)

	log.Fatal(s.Run())

}
