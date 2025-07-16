package mapdb

import (
	"github.com/identityofsine/fofx-go-gin-api-template/internal/components/maps/model/mapmodel"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/db"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/db/dbmapper"
)

const (
	table = "maps"
)

func GetMaps() (*[]mapmodel.Map, db.DatabaseError) {
	dbs, err := db.Query[MapDb]("SELECT * from " + table)
	if err != nil {
		return nil, err
	}

	// Map the database models to the application models
	maps := dbmapper.MapAllDbFields[MapDb, mapmodel.Map](*dbs)
	if maps == nil {
		return nil, db.NewDatabaseError("GetMaps", "Mapping failed", "mapping-failed", 500)
	}

	return maps, nil
}
