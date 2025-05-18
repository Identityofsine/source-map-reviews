package model

import (
	"database/sql"
	"time"

	"github.com/identityofsine/fofx-go-gin-api-template/pkg/db"
)

type LogDB struct {
	Id        int
	Severity  string
	Message   string
	CreatedAt time.Time
	Version   sql.NullString
	Commit    sql.NullString
}

func GetLogs() ([]LogDB, db.DatabaseError) {
	query := "SELECT * FROM public.logs"
	rows, err := db.Query[LogDB](query)

	if err != nil {
		return nil, err
	}

	return *rows, err
}

func SaveLogs(logDB LogDB) db.DatabaseError {
	query := "INSERT INTO public.logs (severity, message, created_at, version, commit_hash) VALUES ($1, $2, $3, $4, $5)"
	_, err := db.Query[LogDB](query, logDB.Severity, logDB.Message, logDB.CreatedAt, logDB.Version, logDB.Commit)
	if err != nil {
		return err
	}
	return nil
}
