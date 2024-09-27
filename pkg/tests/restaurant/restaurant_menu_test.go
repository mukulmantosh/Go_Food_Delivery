package restaurant

import (
	"Go_Food_Delivery/pkg/handler"
	"Go_Food_Delivery/pkg/handler/restaurant"
	restro "Go_Food_Delivery/pkg/service/restaurant"
	"Go_Food_Delivery/pkg/tests"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestRestaurantMenu(t *testing.T) {
	t.Setenv("APP_ENV", "TEST")
	t.Setenv("STORAGE_TYPE", "local")
	t.Setenv("STORAGE_DIRECTORY", "uploads")
	t.Setenv("LOCAL_STORAGE_PATH", "./tmp")
	testDB := tests.Setup()
	AppEnv := os.Getenv("APP_ENV")
	testServer := handler.NewServer(testDB, false)

	// Restaurant
	restaurantService := restro.NewRestaurantService(testDB, AppEnv)
	restaurant.NewRestaurantHandler(testServer, "/restaurant", restaurantService)

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

	form := FakeRestaurant{
		Name:        name,
		File:        file,
		Description: description,
		Address:     address,
		City:        city,
		State:       state,
	}

	body, contentType, err := GenerateData(form)
	if err != nil {
		t.Fatalf("Error generating form-data: %v", err)
	}

	t.Run("Restaurant::Create", func(t *testing.T) {

		req, _ := http.NewRequest(http.MethodPost, "/restaurant/", body)
		req.Header.Set("Content-Type", contentType)
		w := httptest.NewRecorder()
		testServer.Gin.ServeHTTP(w, req)
		assert.Equal(t, http.StatusCreated, w.Code)

	})

	t.Run("Restaurant::Listing", func(t *testing.T) {

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
		err := json.Unmarshal(w.Body.Bytes(), &restaurants)
		if err != nil {
			t.Fatalf("Failed to decode response body: %v", err)
		}

		// set the restaurantID
		RestaurantResponseID = restaurants[0].RestaurantID

	})

	t.Run("RestaurantMenu::Create", func(t *testing.T) {
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

	t.Run("RestaurantMenu::List", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/restaurant/menu", nil)
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		testServer.Gin.ServeHTTP(w, req)

		var menuItems []MenuItem
		err := json.Unmarshal(w.Body.Bytes(), &menuItems)
		if err != nil {
			fmt.Println("Error unmarshalling JSON:", err)
			return
		}

		RestaurantMenuID = int64(menuItems[0].MenuID)

		assert.Equal(t, http.StatusOK, w.Code)

	})

	t.Run("RestaurantMenu::List::ById", func(t *testing.T) {
		url := fmt.Sprintf("%s%d", "/restaurant/menu?restaurant_id=", RestaurantMenuID)
		req, _ := http.NewRequest(http.MethodGet, url, nil)
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		testServer.Gin.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("RestaurantMenu::Delete", func(t *testing.T) {
		url := fmt.Sprintf("%s%d/%d", "/restaurant/menu/", RestaurantResponseID, RestaurantMenuID)
		req, _ := http.NewRequest(http.MethodDelete, url, nil)
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		testServer.Gin.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNoContent, w.Code)

	})
	// Cleanup
	t.Cleanup(func() {
		tests.Teardown(testDB)
	})

}
