package restaurant

import (
	"Go_Food_Delivery/pkg/handler"
	"Go_Food_Delivery/pkg/handler/restaurant"
	restro "Go_Food_Delivery/pkg/service/restaurant"
	"Go_Food_Delivery/pkg/tests"
	"encoding/json"
	"fmt"
	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestRestaurant(t *testing.T) {
	t.Setenv("APP_ENV", "TEST")
	t.Setenv("STORAGE_TYPE", "local")
	t.Setenv("STORAGE_DIRECTORY", "uploads")
	t.Setenv("LOCAL_STORAGE_PATH", "./tmp")
	testDB := tests.Setup()
	testServer := handler.NewServer(testDB, false)
	AppEnv := os.Getenv("APP_ENV")

	// Restaurant
	restaurantService := restro.NewRestaurantService(testDB, AppEnv)
	restaurant.NewRestaurantHandler(testServer, "/restaurant", restaurantService)

	var RestaurantResponseID int64
	name := faker.Name()
	file := []byte{10, 10, 10, 10, 10} // fake image bytes
	description := faker.Paragraph()
	address := faker.Word()
	city := faker.Word()
	state := faker.Word()

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
		req.Header.Set("Content-Type", "application/Json")
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

	t.Run("Restaurant::GetById", func(t *testing.T) {
		url := fmt.Sprintf("/restaurant/%d", RestaurantResponseID)
		req, _ := http.NewRequest(http.MethodGet, url, nil)
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		testServer.Gin.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("Restaurant::Delete", func(t *testing.T) {
		url := fmt.Sprintf("/restaurant/%d", RestaurantResponseID)
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
