package cart

import (
	"Go_Food_Delivery/cmd/api/middleware"
	restroTypes "Go_Food_Delivery/pkg/database/models/restaurant"
	userModel "Go_Food_Delivery/pkg/database/models/user"
	"Go_Food_Delivery/pkg/handler"
	crt "Go_Food_Delivery/pkg/handler/cart"
	"Go_Food_Delivery/pkg/handler/restaurant"
	"Go_Food_Delivery/pkg/handler/user"
	natsPkg "Go_Food_Delivery/pkg/nats"
	"Go_Food_Delivery/pkg/service/cart_order"
	restro "Go_Food_Delivery/pkg/service/restaurant"
	usr "Go_Food_Delivery/pkg/service/user"
	"Go_Food_Delivery/pkg/tests"
	common "Go_Food_Delivery/pkg/tests/restaurant"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/nats"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-faker/faker/v4"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestCart(t *testing.T) {
	t.Setenv("APP_ENV", "TEST")
	t.Setenv("STORAGE_TYPE", "local")
	t.Setenv("STORAGE_DIRECTORY", "uploads")
	t.Setenv("LOCAL_STORAGE_PATH", "./tmp")
	t.Setenv("GIN_MODE", "release")

	testDB := tests.Setup()
	AppEnv := os.Getenv("APP_ENV")
	testServer := handler.NewServer(testDB, false)
	validate := validator.New()

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(60*time.Second))
	defer cancel()

	natsContainer, err := nats.Run(ctx, "nats:2.10.20", testcontainers.WithHostPortAccess(4222))

	if err != nil {
		t.Logf("failed to start NATS container: %s", err)
		return
	}
	connectionString, err := natsContainer.ConnectionString(ctx)
	if err != nil {
		t.Log("NATS Connection String Error::", err)
		return
	}

	// Connect NATS
	natTestServer, err := natsPkg.NewNATS(connectionString)
	middlewares := []gin.HandlerFunc{middleware.AuthMiddleware()}

	// User
	userService := usr.NewUserService(testDB, AppEnv)
	user.NewUserHandler(testServer, "/user", userService, validate)

	// Restaurant
	restaurantService := restro.NewRestaurantService(testDB, AppEnv)
	restaurant.NewRestaurantHandler(testServer, "/restaurant", restaurantService)

	// Cart
	cartService := cart_order.NewCartService(testDB, AppEnv, natTestServer)
	crt.NewCartHandler(testServer, "/cart", cartService, middlewares, validate)

	var RestaurantResponseID int64
	var RestaurantMenuID int64
	name := faker.Name()
	file := []byte{10, 10, 10, 10, 10} // fake image bytes
	description := faker.Paragraph()
	address := faker.Word()
	city := faker.Word()
	state := faker.Word()

	type FakeRestaurantMenu struct {
		RestaurantID int64   `json:"restaurant_id"`
		Name         string  `json:"name"`
		Description  string  `json:"description"`
		Price        float64 `json:"price"`
		Category     string  `json:"category"`
		Available    bool    `json:"available"`
	}

	type CartParams struct {
		ItemID       int64 `json:"item_id"`
		RestaurantID int64 `json:"restaurant_id"`
		Quantity     int64 `json:"quantity"`
	}

	form := common.FakeRestaurant{
		Name:        name,
		File:        file,
		Description: description,
		Address:     address,
		City:        city,
		State:       state,
	}

	body, contentType, err := common.GenerateData(form)
	if err != nil {
		t.Fatalf("Error generating form-data: %v", err)
	}

	type FakeUser struct {
		User     string `json:"user" faker:"name"`
		Email    string `json:"email" faker:"email"`
		Password string `json:"password" faker:"password"`
	}

	var customUser FakeUser
	var userInfo userModel.User
	_ = faker.FakeData(&customUser)
	userInfo.Email = customUser.Email
	userInfo.Password = customUser.Password

	_, err = userService.Add(ctx, &userInfo)
	if err != nil {
		t.Error(err)
	}

	loginToken, err := userService.Login(ctx, userInfo.ID, "Food Delivery")
	if err != nil {
		t.Fatal(err)
	}

	Token := fmt.Sprintf("Bearer %s", loginToken)

	t.Run("Cart::Restaurant::Create", func(t *testing.T) {

		req, _ := http.NewRequest(http.MethodPost, "/restaurant/", body)
		req.Header.Set("Content-Type", contentType)
		w := httptest.NewRecorder()
		testServer.Gin.ServeHTTP(w, req)
		assert.Equal(t, http.StatusCreated, w.Code)

	})

	t.Run("Cart::Restaurant::Listing", func(t *testing.T) {

		type RestaurantResponse struct {
			RestaurantID int64  `json:"restaurant_id"`
			Name         string `json:"name"`
			StoreImage   string `json:"store_image"`
			Description  string `json:"description"`
			Address      string `json:"address"`
			City         string `json:"city"`
			State        string `json:"state"`
			CreatedAt    string `json:"CreatedAt"`
			UpdatedAt    string `json:"UpdatedAt"`
		}

		req, _ := http.NewRequest(http.MethodGet, "/restaurant/", nil)
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		testServer.Gin.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)

		var restaurants []RestaurantResponse
		err := json.NewDecoder(strings.NewReader(w.Body.String())).Decode(&restaurants)
		if err != nil {
			t.Fatalf("Failed to decode response body: %v", err)
		}

		// set the restaurantID
		RestaurantResponseID = restaurants[0].RestaurantID

	})

	t.Run("Cart::RestaurantMenu::Create", func(t *testing.T) {
		var customMenu FakeRestaurantMenu

		customMenu.Available = true
		customMenu.Price = 40.35
		customMenu.Name = "burger"
		customMenu.Description = "burger"
		customMenu.Category = "FAST_FOODS"
		customMenu.RestaurantID = RestaurantResponseID
		payload, err := json.Marshal(&customMenu)
		if err != nil {
			t.Fatal("Error::", err)
		}
		req, _ := http.NewRequest(http.MethodPost, "/restaurant/menu", bytes.NewBuffer(payload))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		testServer.Gin.ServeHTTP(w, req)
		assert.Equal(t, http.StatusCreated, w.Code)

	})

	t.Run("Cart::RestaurantMenu::List", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/restaurant/menu", nil)
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		testServer.Gin.ServeHTTP(w, req)
		var menuItems []restroTypes.MenuItem
		err := json.Unmarshal(w.Body.Bytes(), &menuItems)
		if err != nil {
			fmt.Println("Error unmarshalling JSON:", err)
			return
		}

		RestaurantMenuID = menuItems[0].MenuID

		assert.Equal(t, http.StatusOK, w.Code)

	})

	t.Run("Cart::AddItemToCart", func(t *testing.T) {
		var cartParams CartParams
		cartParams.ItemID = RestaurantMenuID
		cartParams.RestaurantID = RestaurantResponseID
		cartParams.Quantity = 1
		payload, err := json.Marshal(&cartParams)
		if err != nil {
			t.Fatal("Error::", err)
		}
		req, _ := http.NewRequest(http.MethodPost, "/cart/add", bytes.NewBuffer(payload))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", Token)
		w := httptest.NewRecorder()
		testServer.Gin.ServeHTTP(w, req)
		assert.Equal(t, http.StatusCreated, w.Code)

	})

	t.Run("Cart::List", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/cart/list", nil)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", Token)

		w := httptest.NewRecorder()
		testServer.Gin.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)

	})

	t.Run("Cart::PlaceOrder", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodPost, "/cart/order/new", nil)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", Token)

		w := httptest.NewRecorder()
		testServer.Gin.ServeHTTP(w, req)
		assert.Equal(t, http.StatusCreated, w.Code)

	})

	t.Cleanup(func() {
		tests.Teardown(testDB)
	})
}
