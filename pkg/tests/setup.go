package tests

import (
	"Go_Food_Delivery/pkg/database"
	"log"
	"log/slog"
)

// Setup will be bootstrapping our test db.
func Setup() database.Database {
	slog.Info("Initializing Setup..")
	testDb := database.NewTestDB()

	if err := testDb.Migrate(); err != nil {
		log.Fatalf("Error migrating database: %s", err)
	}
	return testDb
}

func Teardown(testDB database.Database) {
	err := testDB.Close()
	if err != nil {
		log.Fatalf("Error closing testDB: %s", err)
	}
}
