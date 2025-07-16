package mapdb

import (
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/db"
)

const (
	table = "maps"
)

func GetMaps() (*[]MapDb, db.DatabaseError) {
	dbs, err := db.Query[MapDb]("SELECT * from " + table)
	if err != nil {
		return nil, err
	}

	return dbs, nil
}
