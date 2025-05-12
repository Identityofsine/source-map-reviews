package model

import (
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/db"
	"time"
)

type LogDB struct {
	Id        int
	Severity  string
	Message   string
	CreatedAt time.Time
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
	query := "INSERT INTO public.logs (severity, message, created_at) VALUES ($1, $2, $3)"
	_, err := db.Query[LogDB](query, logDB.Severity, logDB.Message, logDB.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}
