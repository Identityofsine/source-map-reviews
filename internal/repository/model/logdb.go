package model

import (
	"time"

	"github.com/identityofsine/fofx-go-gin-api-template/pkg/db"
)

type LogDB struct {
	name      string
	db        string
	level     string
	message   string
	timestamp time.Time
}

func getLogs() []LogDB {
	query := "SELECT * FROM public.logs"
	rows, err := db.Query[LogDB](query, nil)
	if err != nil {
		return nil
	}

	return *rows
}
