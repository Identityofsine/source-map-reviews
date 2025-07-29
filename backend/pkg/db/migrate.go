package db

import (
	"database/sql"
	"log"
	"os"

	"github.com/pressly/goose"
)

func Migrate() DatabaseError {
	db, err := Connect()
	env := os.Getenv("GO_ENV")

	if err != nil {
		return err
	}

	if err := goose.SetDialect("postgres"); err != nil {
		return NewDatabaseError("db", "Error setting dialect", err.Error(), 500)
	}

	log.Printf("[DB:INFO] Migrating database %s", os.Getenv("DB_NAME"))

	// Use a separate goose version table for each migration folder
	goose.SetTableName("goose_db_version_functions")
	log.Printf("[DB:INFO] Running migrations in functions")
	if err := goose.Up(db, "./migrations/functions"); err != nil {
		log.Fatalf("[DB:ERROR] %s in %s", err, "./migrations/functions")
		return NewDatabaseError("db", "Error running migrations in functions", err.Error(), 500)
	}

	goose.SetTableName("goose_db_version")
	log.Printf("[DB:INFO] Running migrations in init")
	if err := goose.Up(db, "./migrations/init"); err != nil {
		log.Fatalf("[DB:ERROR] %s in %s", err, "./migrations/init")
		return NewDatabaseError("db", "Error running migrations in init", err.Error(), 500)
	}

	goose.SetTableName("goose_db_version_changelogs")
	log.Printf("[DB:INFO] Running migrations in changelogs")
	if err := goose.Up(db, "./migrations/changelogs"); err != nil {
		log.Fatalf("[DB:ERROR] %s in %s", err, "./migrations/changelogs")
		return NewDatabaseError("db", "Error running migrations in changelogs", err.Error(), 500)
	}

	if env == "development" {
		goose.SetTableName("goose_db_version_localonly")
		log.Printf("[DB:INFO] Running migrations in localonly")
		if err := goose.Up(db, "./migrations/localonly"); err != nil {
			log.Fatalf("[DB:ERROR] %s in %s", err, "./migrations/localonly")
			return NewDatabaseError("db", "Error running migrations in localonly", err.Error(), 500)
		}
	}

	return nil
}

func Close(db *sql.DB) DatabaseError {
	if db == nil {
		return NewDatabaseError("db", "Error closing database", "Database is nil", 500)
	}
	database.Close()
	database = nil
	db = nil
	return nil
}
