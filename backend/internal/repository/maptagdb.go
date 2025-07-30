package repository

import (
	"fmt"

	"github.com/identityofsine/fofx-go-gin-api-template/internal/constants/exception"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/db"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/db/dao"
	"github.com/identityofsine/fofx-go-gin-api-template/util"
)

// MapTagDb represents a row in the map_tags table (junction table linking maps to tags)
type MapTagDb struct {
	LkTag     string `db:"lk_tag"`
	MapName   string `db:"map_name"`
	CreatedAt string `db:"created_at" dao:"omit"`
	UpdatedAt string `db:"updated_at" dao:"omit"`
}

const (
	map_table = "map_tags"
)

type MapTagRelationshipDbs = map[string][]MapTagDb

// GetMapTags retrieves all map-tag links from the map_tags table
func GetMapTags() (*[]MapTagDb, db.DatabaseError) {
	dbs, err := dao.SelectFromDatabaseByStruct(MapTagDb{}, "")
	if err != nil {
		return nil, err
	}

	return &dbs, nil
}

// GetMapTagsByMapName retrieves all map-tag links from the map_tags table for a given map name
func GetMapTagsByMapName(mapName string) (*[]MapTagDb, db.DatabaseError) {
	dbs, err := dao.SelectFromDatabaseByStruct(MapTagDb{}, "map_name = $1", mapName)
	if err != nil {
		return nil, err
	}

	return &dbs, nil

}

func GetMapTagsByMapNames(mapNames []string) (*MapTagRelationshipDbs, db.DatabaseError) {

	if mapNames == nil || len(mapNames) == 0 {
		return nil, exception.BadRequestDatabase
	}

	dbs, err := dao.SelectFromDatabaseByStruct(MapTagDb{}, fmt.Sprintf("map_name IN (%s)", db.Placeholders(len(mapNames))), mapNames)
	if err != nil {
		return nil, err
	}

	mapped := util.GroupBy(dbs, func(item MapTagDb) string {
		return item.MapName
	})

	return &mapped, nil

}
