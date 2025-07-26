package maptagdb

import (
	"fmt"
	"strings"

	"github.com/identityofsine/fofx-go-gin-api-template/pkg/db"
	"github.com/identityofsine/fofx-go-gin-api-template/util"
)

const (
	table = "map_tags"
)

type MapTagRelationshipDbs = map[string][]MapTagDb

// GetMapTags retrieves all map-tag links from the map_tags table
func GetMapTags() (*[]MapTagDb, db.DatabaseError) {
	dbs, err := db.Query[MapTagDb]("SELECT * from " + table)
	if err != nil {
		return nil, err
	}
	return dbs, nil
}

// GetMapTagsByMapName retrieves all map-tag links from the map_tags table for a given map name
func GetMapTagsByMapName(mapName string) (*[]MapTagDb, db.DatabaseError) {
	dbs, err := db.Query[MapTagDb]("SELECT * from "+table+" where map_name = $1", mapName)
	if err != nil {
		return nil, err
	}
	return dbs, nil
}

func GetMapTagsByMapNames(mapNames []string) (*MapTagRelationshipDbs, db.DatabaseError) {

	placeholders := make([]string, len(mapNames))
	args := make([]interface{}, len(mapNames))
	if len(mapNames) == 0 {
		return nil, nil
	}
	for i, name := range mapNames {
		placeholders[i] = fmt.Sprintf("$%d", i+1) // PostgreSQL style
		args[i] = name
	}
	query := fmt.Sprintf("SELECT * FROM "+table+" WHERE map_name IN (%s)", strings.Join(placeholders, ","))

	dbs, err := db.Query[MapTagDb](query, args...)
	if err != nil {
		return nil, err
	}
	grouped := util.GroupBy[MapTagDb](*dbs, func(item MapTagDb) string {
		return item.MapName
	})

	return &grouped, nil
}
