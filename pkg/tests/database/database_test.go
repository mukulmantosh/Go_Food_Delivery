package database

import (
	"Go_Food_Delivery/pkg/database"
	"context"
	"log"
	"os"
	"testing"
	"time"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

var (
	pgContainer    testcontainers.Container
	containerImage string = "postgres:16.3"
	dbHost         string
	dbPort         string
	dbName         string = "test-db"
	dbUsername     string = "postgres"
	dbPassword     string = "postgres"
	err            error
)

// Initialize the PostgreSQL container
func setup() {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	pgContainer, err = postgres.Run(ctx, containerImage,
		postgres.WithDatabase(dbName),
		postgres.WithUsername(dbUsername),
		postgres.WithPassword(dbPassword),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).WithStartupTimeout(10*time.Second)),
	)
	if err != nil {
		log.Fatalf("failed to start PostgreSQL container: %s", err)
	}

	// Set environment variables
	dbHost, _ = pgContainer.Host(ctx)
	mappedPort, _ := pgContainer.MappedPort(ctx, "5432/tcp")
	dbPort = mappedPort.Port()
	_ = os.Setenv("DB_USERNAME", dbUsername)
	_ = os.Setenv("DB_PASSWORD", dbPassword)
	_ = os.Setenv("DB_NAME", dbName)
	_ = os.Setenv("DB_PORT", dbPort)
	_ = os.Setenv("DB_HOST", dbHost)
	_ = os.Setenv("STORAGE_TYPE", "local")
	_ = os.Setenv("STORAGE_DIRECTORY", "uploads")
	_ = os.Setenv("LOCAL_STORAGE_PATH", "./tmp")

}

// Teardown function to terminate the container
func teardown() {
	if err := pgContainer.Terminate(context.Background()); err != nil {
		log.Fatalf("failed to terminate PostgreSQL container: %s", err)
	}
}

func TestMain(m *testing.M) {
	setup()

	result := m.Run()

	//teardown()
	os.Exit(result)
}

func TestDatabase(t *testing.T) {
	dbTest := database.New()

	if dbTest.HealthCheck() != true {
		t.Fatalf("database health check failed")
	}
	if err := dbTest.Migrate(); err != nil {
		t.Fatalf("Error migrating database: %s", err)
	}

}
