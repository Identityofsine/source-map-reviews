package db

import (
	"database/sql"
	_ "errors"
	"fmt"
	_ "fmt"
	_ "github.com/lib/pq" // This is important! The underscore is to import the package for side-effects
	"github.com/pressly/goose"
	"log"
	"os"
	"reflect"
	"strings"
)

var cfg = fmt.Sprintf(
	"user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
	os.Getenv("DB_USER"),
	os.Getenv("DB_PASSWORD"),
	os.Getenv("DB_NAME"),
	"db",   // e.g., "db" if using Docker Compose
	"5432", // e.g., "5432"
)

var database *sql.DB = nil

func Connect() (*sql.DB, DatabaseError) {
	if database == nil {
		db, err := sql.Open("postgres", cfg)
		if err != nil {
			return nil, NewDatabaseError("db", "Error connecting to database", err.Error(), 500)
		}
		if err := db.Ping(); err != nil {
			log.Printf("[DB:ERROR] %s", err)
			return nil, NewDatabaseError("db", "Error connecting to database", err.Error(), 500)
		}
		database = db
		return db, nil
	} else {
		return database, nil
	}
}

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

	log.Printf("[DB:INFO] Running migrations in functions")
	if err := goose.Up(db, "./migrations/functions"); err != nil {
		log.Fatalf("[DB:ERROR] %s in %s", err, "./migrations/functions")
		return NewDatabaseError("db", "Error running migrations in functions", err.Error(), 500)
	}

	log.Printf("[DB:INFO] Running migrations in init")
	if err := goose.Up(db, "./migrations/init"); err != nil {
		log.Fatalf("[DB:ERROR] %s in %s", err, "./migrations/init")
		return NewDatabaseError("db", "Error running migrations in init", err.Error(), 500)
	}

	log.Printf("[DB:INFO] Running migrations in changelogs")
	if err := goose.Up(db, "./migrations/changelogs"); err != nil {
		log.Fatalf("[DB:ERROR] %s in %s", err, "./migrations/changelogs")
		return NewDatabaseError("db", "Error running migrations in changelogs", err.Error(), 500)
	}

	if env == "development" {
		log.Printf("[DB:INFO] Running migrations in localonly")
		if err := goose.Up(db, "./migrations/localonly"); err != nil {
			log.Fatalf("[DB:ERROR] %s in %s", err, "./migrations/changelogs")
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

func Get() *sql.DB {
	return database
}

func Query[T interface{}](query string, placeholders ...any) (*[]T, DatabaseError) {
	db, err := Connect()
	//close
	if err != nil {
		return nil, NewDatabaseError("db", "Error connecting to database", err.Error(), 500)
	}
	rows, r_err := db.Query(query, placeholders...)
	if r_err != nil {
		return nil, NewDatabaseError("db", "Error executing query", r_err.Error(), 500)
	}
	defer rows.Close()
	results := []any{}
	obj := new(T)
	for rows.Next() {
		c := reflect.New(reflect.TypeOf(obj).Elem()).Interface()
		v := reflect.ValueOf(c).Elem()

		var cols []interface{}

		if v.Kind() == reflect.Struct {
			numCols := v.NumField()
			cols = make([]interface{}, numCols)
			for i := 0; i < numCols; i++ {
				f := v.Field(i)
				cols[i] = f.Addr().Interface()
			}
		} else {
			// Non-struct: scan directly into the single value
			cols = []interface{}{v.Addr().Interface()}
		}

		err := rows.Scan(cols...)
		if err != nil {
			return nil, NewDatabaseError("db", "Error scanning row", err.Error(), 500)
		}

		results = append(results, c)
	}
	dtos := make([]T, 0)
	for _, u := range results {
		castedObject := *u.(*T)
		dtos = append(dtos, castedObject)
	}
	return &dtos, nil
}

func Insert(query string, placeholders ...any) (int64, DatabaseError) {
	db, err := Connect()
	if err != nil {
		return 0, NewDatabaseError("db", "Error connecting to database", err.Error(), 500)
	}
	res, r_err := db.Exec(query, placeholders...)
	if r_err != nil {
		return 0, NewDatabaseError("db", "Error executing query", r_err.Error(), 500)
	}
	id, r_err := res.LastInsertId()
	if r_err != nil {
		return 0, NewDatabaseError("db", "Error getting last insert id", r_err.Error(), 500)
	}
	return id, nil
}

func Delete(query string, placeholders ...any) (int64, DatabaseError) {
	db, err := Connect()
	if err != nil {
		return 0, NewDatabaseError("db", "Error connecting to database", err.Error(), 500)
	}
	res, r_err := db.Exec(query, placeholders...)
	if r_err != nil {
		return 0, NewDatabaseError("db", "Error executing query", r_err.Error(), 500)
	}
	rows, r_err := res.RowsAffected()
	if r_err != nil {
		return 0, NewDatabaseError("db", "Error getting rows affected", r_err.Error(), 500)
	}
	return rows, nil
}

func Exists(query string, placeholders ...any) (bool, DatabaseError) {
	db, err := Connect()
	if err != nil {
		return false, NewDatabaseError("db", "Error connecting to database", err.Error(), 500)
	}
	rows, r_err := db.Query(query, placeholders...)
	if r_err != nil {
		return false, NewDatabaseError("db", "Error executing query", r_err.Error(), 500)
	}
	defer rows.Close()
	return rows.Next(), nil
}

func Sanitize(query string) string {
	return strings.ReplaceAll(query, "'", "''")
}

func Cast[T any](a any) *T {
	ok := *(a).(*T)
	return &ok
}
