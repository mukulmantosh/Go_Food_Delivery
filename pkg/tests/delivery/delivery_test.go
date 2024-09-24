package delivery

import (
	"Go_Food_Delivery/cmd/api/middleware"
	"Go_Food_Delivery/pkg/handler"
	delv "Go_Food_Delivery/pkg/handler/delivery"
	"Go_Food_Delivery/pkg/handler/user"
	natsPkg "Go_Food_Delivery/pkg/nats"
	"Go_Food_Delivery/pkg/service/delivery"
	usr "Go_Food_Delivery/pkg/service/user"
	"Go_Food_Delivery/pkg/tests"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/nats"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"
)

func TestAddDeliveryUser(t *testing.T) {
	t.Setenv("APP_ENV", "TEST")
	t.Setenv("STORAGE_TYPE", "local")
	t.Setenv("STORAGE_DIRECTORY", "uploads")
	t.Setenv("LOCAL_STORAGE_PATH", "./tmp")
	testDB := tests.Setup()
	AppEnv := os.Getenv("APP_ENV")
	testServer := handler.NewServer(testDB, false)
	middlewares := []gin.HandlerFunc{middleware.AuthMiddleware()}

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
	natTestServer, _ := natsPkg.NewNATS(connectionString)

	validate := validator.New()
	userService := usr.NewUserService(testDB, AppEnv)
	user.NewUserHandler(testServer, "/user", userService, validate)

	deliveryService := delivery.NewDeliveryService(testDB, os.Getenv("APP_ENV"), natTestServer)
	delv.NewDeliveryHandler(testServer, "/delivery", deliveryService, middlewares, validate)

	type FakeDeliveryUser struct {
		Name           string `json:"name"`
		Phone          string `json:"phone"`
		VehicleDetails string `json:"vehicle_details"`
	}

	t.Run("Delivery::User::Create", func(t *testing.T) {

		var customUser FakeDeliveryUser
		customUser.Name = "Test"
		customUser.Phone = "08090909090"
		customUser.VehicleDetails = "OX-25895-8547"
		payload, err := json.Marshal(&customUser)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(string(payload))

		req, _ := http.NewRequest(http.MethodPost, "/delivery/add", strings.NewReader(string(payload)))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		testServer.Gin.ServeHTTP(w, req)
		t.Log(w.Body.String())
		assert.Equal(t, http.StatusCreated, w.Code)
	})

}
