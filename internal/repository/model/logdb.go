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
