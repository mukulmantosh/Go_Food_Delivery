package main

import (
	"Go_Food_Delivery/cmd/api/middleware"
	"Go_Food_Delivery/pkg/database"
	"Go_Food_Delivery/pkg/handler"
	"Go_Food_Delivery/pkg/handler/annoucements"
	crt "Go_Food_Delivery/pkg/handler/cart"
	delv "Go_Food_Delivery/pkg/handler/delivery"
	notify "Go_Food_Delivery/pkg/handler/notification"
	"Go_Food_Delivery/pkg/handler/restaurant"
	revw "Go_Food_Delivery/pkg/handler/review"
	"Go_Food_Delivery/pkg/handler/user"
	"Go_Food_Delivery/pkg/nats"
	"Go_Food_Delivery/pkg/service/announcements"
	"Go_Food_Delivery/pkg/service/cart_order"
	"Go_Food_Delivery/pkg/service/delivery"
	"Go_Food_Delivery/pkg/service/notification"
	restro "Go_Food_Delivery/pkg/service/restaurant"
	"Go_Food_Delivery/pkg/service/review"
	usr "Go_Food_Delivery/pkg/service/user"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/websocket"
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

	// Connect NATS
	natServer, err := nats.NewNATS("nats://127.0.0.1:4222")

	// WebSocket Clients
	wsClients := make(map[string]*websocket.Conn)

	s := handler.NewServer(db, true)

	// Initialize Validator
	validate := validator.New()

	// Middlewares List
	middlewares := []gin.HandlerFunc{middleware.AuthMiddleware()}

	// User
	userService := usr.NewUserService(db, env)
	user.NewUserHandler(s, "/user", userService, validate)

	// Restaurant
	restaurantService := restro.NewRestaurantService(db, env)
	restaurant.NewRestaurantHandler(s, "/restaurant", restaurantService)

	// Reviews
	reviewService := review.NewReviewService(db, env)
	revw.NewReviewProtectedHandler(s, "/review", reviewService, middlewares, validate)

	// Cart
	cartService := cart_order.NewCartService(db, env, natServer)
	crt.NewCartHandler(s, "/cart", cartService, middlewares, validate)

	// Delivery
	deliveryService := delivery.NewDeliveryService(db, env, natServer)
	delv.NewDeliveryHandler(s, "/delivery", deliveryService, middlewares, validate)

	// Notification
	notifyService := notification.NewNotificationService(db, env, natServer)

	// Subscribe to multiple events.
	_ = notifyService.SubscribeNewOrders(wsClients)
	_ = notifyService.SubscribeOrderStatus(wsClients)

	notify.NewNotifyHandler(s, "/notify", notifyService, middlewares, validate, wsClients)

	// Events/Announcements
	announceService := announcements.NewAnnouncementService(db, env)
	annoucements.NewAnnouncementHandler(s, "/announcements", announceService, middlewares, validate)
	log.Fatal(s.Run())

}
