package main

import (
	"Go_Food_Delivery/pkg/database"
	"Go_Food_Delivery/pkg/handler"
	"Go_Food_Delivery/pkg/handler/restaurant"
	"Go_Food_Delivery/pkg/handler/user"
	restro "Go_Food_Delivery/pkg/service/restaurant"
	usr "Go_Food_Delivery/pkg/service/user"
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

	env := os.Getenv("APP_ENV")
	db := database.New()
	// Create Tables
	if err := db.Migrate(); err != nil {
		log.Fatalf("Error migrating database: %s", err)
	}

	s := handler.NewServer(db)

	// User
	userService := usr.NewUserService(db, env)
	user.NewUserHandler(s, "/user", userService)

	// Restaurant
	restaurantService := restro.NewRestaurantService(db, env)
	restaurant.NewRestaurantHandler(s, "/restaurant", restaurantService)

	log.Fatal(s.Run())

}
