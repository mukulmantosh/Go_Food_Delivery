package restaurant

import (
	"Go_Food_Delivery/cmd/api/middleware"
	restroModel "Go_Food_Delivery/pkg/database/models/restaurant"
	"Go_Food_Delivery/pkg/database/models/review"
	userModel "Go_Food_Delivery/pkg/database/models/user"
	"Go_Food_Delivery/pkg/handler"
	revw "Go_Food_Delivery/pkg/handler/review"
	restro "Go_Food_Delivery/pkg/service/restaurant"
	reviewSrv "Go_Food_Delivery/pkg/service/review"
	usr "Go_Food_Delivery/pkg/service/user"
	"Go_Food_Delivery/pkg/tests"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-faker/faker/v4"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestReview(t *testing.T) {
	t.Setenv("APP_ENV", "TEST")
	t.Setenv("STORAGE_TYPE", "local")
	t.Setenv("STORAGE_DIRECTORY", "uploads")
	t.Setenv("LOCAL_STORAGE_PATH", "./tmp")
	testDB := tests.Setup()
	validate := validator.New()
	AppEnv := os.Getenv("APP_ENV")
	testServer := handler.NewServer(testDB, false)
	middlewares := []gin.HandlerFunc{middleware.AuthMiddleware()}

	userService := usr.NewUserService(testDB, AppEnv)
	restaurantService := restro.NewRestaurantService(testDB, AppEnv)
	reviewService := reviewSrv.NewReviewService(testDB, AppEnv)
	revw.NewReviewProtectedHandler(testServer, "/review", reviewService, middlewares, validate)

	type FakeUser struct {
		User     string `json:"user" faker:"name"`
		Email    string `json:"email" faker:"email"`
		Password string `json:"password" faker:"password"`
	}
	var ReviewResponseID int64

	var customUser FakeUser
	var user userModel.User
	var restrro restroModel.Restaurant
	_ = faker.FakeData(&customUser)
	user.Email = customUser.Email
	user.Password = customUser.Password

	ctx := context.Background()
	_, err := userService.Add(ctx, &user)
	if err != nil {
		t.Error(err)
	}

	loginToken, err := userService.Login(ctx, user.ID, "Food Delivery")
	if err != nil {
		t.Fatal(err)
	}

	Token := fmt.Sprintf("Bearer %s", loginToken)

	// Restaurant
	name := faker.Name()
	description := faker.Paragraph()
	address := faker.Word()
	city := faker.Word()
	state := faker.Word()

	restrro.Name = name
	restrro.Description = description
	restrro.Address = address
	restrro.City = city
	restrro.State = state

	_, err = restaurantService.Add(ctx, &restrro)
	if err != nil {
		t.Fatal(err)
	}

	restaurants, err := restaurantService.ListRestaurants(ctx)
	if err != nil {
		t.Fatal(err)
	}
	restaurantId := restaurants[0].RestaurantID

	t.Run("Review::Create", func(t *testing.T) {
		var reviewParam review.ReviewParams
		reviewParam.Comment = faker.Word()
		reviewParam.Rating = 4
		payload, err := json.Marshal(&reviewParam)
		if err != nil {
			t.Fatal(err)
		}

		req, _ := http.NewRequest(http.MethodPost,
			fmt.Sprintf("/review/%d", restaurantId),
			strings.NewReader(string(payload)))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", Token)
		w := httptest.NewRecorder()
		testServer.Gin.ServeHTTP(w, req)
		assert.Equal(t, http.StatusCreated, w.Code)

	})

	t.Run("Review::List", func(t *testing.T) {
		type ReviewResponse struct {
			ReviewID     int64  `json:"review_id"`
			UserID       int64  `json:"user_id"`
			RestaurantID int64  `json:"restaurant_id"`
			Rating       int    `json:"rating"`
			Comment      string `json:"comment"`
		}
		req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/review/%d", restaurantId), nil)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", Token)
		w := httptest.NewRecorder()
		testServer.Gin.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)

		var allReviews []ReviewResponse
		err := json.Unmarshal(w.Body.Bytes(), &allReviews)
		if err != nil {
			t.Fatalf("Failed to decode response body: %v", err)
		}

		ReviewResponseID = allReviews[0].ReviewID

	})

	t.Run("Review::Delete", func(t *testing.T) {
		url := fmt.Sprintf("/review/%d", ReviewResponseID)
		req, _ := http.NewRequest(http.MethodDelete, url, nil)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", Token)
		w := httptest.NewRecorder()
		testServer.Gin.ServeHTTP(w, req)
		assert.Equal(t, http.StatusNoContent, w.Code)
	})

	t.Cleanup(func() {
		tests.Teardown(testDB)
	})

}
