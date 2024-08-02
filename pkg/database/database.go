package database

import (
	"Uber_Food_Delivery/pkg/database/models/user"
	"context"
	"database/sql"
	"fmt"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"log"
	"log/slog"
	"os"
	"strconv"
	"time"
)

type Database interface {
	Db() *bun.DB
	Migrate() error
	HealthCheck() bool
	Close() error
}

type DB struct {
	db *bun.DB
}

func (d *DB) Db() *bun.DB {
	return d.db
}

func (d *DB) HealthCheck() bool {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	err := d.db.PingContext(ctx)
	if err != nil {
		slog.Error("DB::error", err)
		return false
	}
	return true
}

func (d *DB) Close() error {
	slog.Info("DB::Closing database connection")
	return d.db.Close()
}

func New() Database {
	dbHost := os.Getenv("DB_HOST")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	databasePort, err := strconv.Atoi(dbPort)
	if err != nil {
		log.Fatal("Invalid DB Port")
	}

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", dbUsername, dbPassword, dbHost, databasePort, dbName)
	database := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	db := bun.NewDB(database, pgdialect.New())
	return &DB{db: db}

}

func (d *DB) Migrate() error {
	models := []interface{}{
		(*user.User)(nil),
	}

	for _, model := range models {
		if err := d.db.ResetModel(context.Background(), model); err != nil {
			return err
		}
	}
	return nil
}
