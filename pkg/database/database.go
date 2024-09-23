package database

import (
	"Go_Food_Delivery/pkg/database/models/cart"
	"Go_Food_Delivery/pkg/database/models/delivery"
	"Go_Food_Delivery/pkg/database/models/order"
	"Go_Food_Delivery/pkg/database/models/restaurant"
	"Go_Food_Delivery/pkg/database/models/review"
	"Go_Food_Delivery/pkg/database/models/user"
	"context"
	"database/sql"
	"fmt"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/driver/sqliteshim"
	"log"
	"log/slog"
	"os"
	"strconv"
	"strings"
	"time"
)

type Database interface {
	Insert(ctx context.Context, model any) (sql.Result, error)
	Delete(ctx context.Context, tableName string, filter Filter) (sql.Result, error)
	Select(ctx context.Context, model any, columnName string, parameter any) error
	SelectAll(ctx context.Context, tableName string, model any) error
	SelectWithRelation(ctx context.Context, model any, relations []string, Condition Filter) error
	SelectWithMultipleFilter(ctx context.Context, model any, Condition Filter) error
	Raw(ctx context.Context, model any, query string, args ...interface{}) error
	Update(ctx context.Context, tableName string, Set Filter, Condition Filter) (sql.Result, error)
	Count(ctx context.Context, tableName string, ColumnExpression string, columnName string, parameter any) (int64, error)
	Migrate() error
	HealthCheck() bool
	Close() error
}

type Filter map[string]any

type DB struct {
	db *bun.DB
}

func (d *DB) Insert(ctx context.Context, model any) (sql.Result, error) {
	result, err := d.db.NewInsert().Model(model).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (d *DB) Update(ctx context.Context, tableName string, Set Filter, Condition Filter) (sql.Result, error) {
	result, err := d.db.NewUpdate().Table(tableName).Set(d.whereCondition(Set, "SET")).Where(d.whereCondition(Condition, "AND")).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (d *DB) Delete(ctx context.Context, tableName string, filter Filter) (sql.Result, error) {
	result, err := d.db.NewDelete().Table(tableName).Where(d.whereCondition(filter, "AND")).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return result, err
}

func (d *DB) Select(ctx context.Context, model any, columnName string, parameter any) error {
	err := d.db.NewSelect().Model(model).Where(fmt.Sprintf("%s = ?", columnName), parameter).Scan(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (d *DB) SelectWithMultipleFilter(ctx context.Context, model any, Condition Filter) error {
	err := d.db.NewSelect().Model(model).Where(d.whereCondition(Condition, "AND")).Scan(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (d *DB) SelectAll(ctx context.Context, tableName string, model any) error {
	if err := d.db.NewSelect().Table(tableName).Scan(ctx, model); err != nil {
		return err
	}
	return nil
}

func (d *DB) SelectWithRelation(ctx context.Context, model any, relations []string, Condition Filter) error {
	query := d.db.NewSelect().Model(model)
	for _, relation := range relations {
		query.Relation(relation)
	}

	err := query.Where(d.whereCondition(Condition, "AND")).Scan(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (d *DB) Raw(ctx context.Context, model any, query string, args ...interface{}) error {
	if err := d.db.NewRaw(query, args...).Scan(ctx, model); err != nil {
		return err
	}
	return nil
}

func (d *DB) Count(ctx context.Context, tableName string, ColumnExpression string, columnName string, parameter any) (int64, error) {
	var count int64
	err := d.db.NewSelect().Table(tableName).ColumnExpr(ColumnExpression).
		Where(fmt.Sprintf("%s = ?", columnName), parameter).Scan(ctx, &count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (d *DB) whereCondition(filter Filter, ConditionType string) string {
	var whereClauses []string
	for key, value := range filter {
		var formattedValue string
		switch v := value.(type) {
		case string:
			// Quote string values
			formattedValue = fmt.Sprintf("'%s'", v)
		case int, int64:
			formattedValue = fmt.Sprintf("%d", v)
		case float64:
			formattedValue = fmt.Sprintf("%.2f", v)
		default:
			log.Fatal("DB::Query:: Un-handled type for where condition!")

		}
		whereClauses = append(whereClauses, fmt.Sprintf("%s=%s", key, formattedValue))
	}

	var result string
	if len(whereClauses) > 0 {
		if ConditionType == "SET" {
			result = strings.Join(whereClauses, " , ")
		} else if ConditionType == "AND" {
			result = strings.Join(whereClauses, " AND ")
		}
	}
	return result
}

func (d *DB) HealthCheck() bool {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	err := d.db.PingContext(ctx)
	if err != nil {
		slog.Error("DB::error", "error", err.Error())
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

// NewTestDB creates a new in-memory test database.
func NewTestDB() Database {
	database, err := sql.Open(sqliteshim.ShimName, "file::memory:?cache=shared")
	if err != nil {
		panic(err)
	}
	db := bun.NewDB(database, sqlitedialect.New())
	return &DB{db}
}

func (d *DB) Migrate() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	models := []interface{}{
		(*user.User)(nil),
		(*restaurant.Restaurant)(nil),
		(*restaurant.MenuItem)(nil),
		(*review.Review)(nil),
		(*order.Order)(nil),
		(*order.OrderItems)(nil),
		(*cart.Cart)(nil),
		(*cart.CartItems)(nil),
		(*delivery.DeliveryPerson)(nil),
		(*delivery.Deliveries)(nil),
	}

	for _, model := range models {
		if _, err := d.db.NewCreateTable().Model(model).WithForeignKeys().IfNotExists().Exec(ctx); err != nil {
			return err
		}
	}
	return nil
}
