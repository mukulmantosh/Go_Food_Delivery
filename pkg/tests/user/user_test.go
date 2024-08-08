package user

import (
	"Go_Food_Delivery/pkg/handler"
	"Go_Food_Delivery/pkg/handler/user"
	"Go_Food_Delivery/pkg/tests"
	"encoding/json"
	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAddUser(t *testing.T) {
	t.Setenv("STORAGE_TYPE", "local")
	t.Setenv("STORAGE_DIRECTORY", "uploads")
	t.Setenv("LOCAL_STORAGE_PATH", "./tmp")
	testDB := tests.Setup()
	testServer := handler.NewServer(testDB)
	user.NewRegister(testServer, "/register")

	type FakeUser struct {
		User     string `json:"user" faker:"name"`
		Email    string `json:"email" faker:"email"`
		Password string `json:"password" faker:"password"`
	}

	t.Run("User::Create", func(t *testing.T) {

		var customUser FakeUser
		_ = faker.FakeData(&customUser)
		payload, err := json.Marshal(&customUser)
		if err != nil {
			t.Fatal(err)
		}

		req, _ := http.NewRequest(http.MethodPost, "/register/user", strings.NewReader(string(payload)))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		testServer.Gin().ServeHTTP(w, req)
		assert.Equal(t, http.StatusCreated, w.Code)

		tests.Teardown(testDB)
	})

}
