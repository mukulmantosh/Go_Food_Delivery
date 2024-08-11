package user

import (
	userModel "Go_Food_Delivery/pkg/database/models/user"
	"Go_Food_Delivery/pkg/handler"
	"Go_Food_Delivery/pkg/handler/user"
	usr "Go_Food_Delivery/pkg/service/user"
	"Go_Food_Delivery/pkg/tests"
	"encoding/json"
	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestAddUser(t *testing.T) {
	t.Setenv("APP_ENV", "TEST")
	t.Setenv("STORAGE_TYPE", "local")
	t.Setenv("STORAGE_DIRECTORY", "uploads")
	t.Setenv("LOCAL_STORAGE_PATH", "./tmp")
	testDB := tests.Setup()
	AppEnv := os.Getenv("APP_ENV")
	testServer := handler.NewServer(testDB)

	// User
	userService := usr.NewUserService(testDB, AppEnv)
	user.NewRegister(testServer, "/user", userService)

	type FakeUser struct {
		User     string `json:"user" faker:"name"`
		Email    string `json:"email" faker:"email"`
		Password string `json:"password" faker:"password"`
	}

	var loggedInUser userModel.LoginUser

	t.Run("User::Create", func(t *testing.T) {

		var customUser FakeUser
		_ = faker.FakeData(&customUser)
		payload, err := json.Marshal(&customUser)
		if err != nil {
			t.Fatal(err)
		}

		loggedInUser.Email = customUser.Email
		loggedInUser.Password = customUser.Password

		req, _ := http.NewRequest(http.MethodPost, "/user/", strings.NewReader(string(payload)))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		testServer.Gin().ServeHTTP(w, req)
		assert.Equal(t, http.StatusCreated, w.Code)

	})

	t.Run("User::Login", func(t *testing.T) {

		payload, err := json.Marshal(&loggedInUser)
		if err != nil {
			t.Fatal(err)
		}

		req, _ := http.NewRequest(http.MethodPost, "/user/login", strings.NewReader(string(payload)))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		testServer.Gin().ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)

	})

	tests.Teardown(testDB)

}
