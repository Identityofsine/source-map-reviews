package repository

import (
	"fmt"

	"github.com/identityofsine/fofx-go-gin-api-template/internal/constants/exception"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/db"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/db/dao"
	"github.com/identityofsine/fofx-go-gin-api-template/util"
)

// MapTagDB represents a row in the map_tags table (junction table linking maps to tags)
type MapTagDB struct {
	LkTag     string `db:"lk_tag"`
	MapName   string `db:"map_name"`
	CreatedAt string `db:"created_at" dao:"omit"`
	UpdatedAt string `db:"updated_at" dao:"omit"`
}

const (
	map_table = "map_tags"
)

type MapTagRelationshipDbs = map[string][]MapTagDB

// GetMapTags retrieves all map-tag links from the map_tags table
func GetMapTags() (*[]MapTagDB, db.DatabaseError) {
	dbs, err := dao.SelectFromDatabaseByStruct(MapTagDB{}, "")
	if err != nil {
		return nil, err
	}

	return &dbs, nil
}

// GetMapTagsByMapName retrieves all map-tag links from the map_tags table for a given map name
func GetMapTagsByMapName(mapName string) (*[]MapTagDB, db.DatabaseError) {
	dbs, err := dao.SelectFromDatabaseByStruct(MapTagDB{}, "map_name = $1", mapName)
	if err != nil {
		return nil, err
	}

	return &dbs, nil

}

func GetMapTagsByMapNames(mapNames []string) (*MapTagRelationshipDbs, db.DatabaseError) {

	if mapNames == nil || len(mapNames) == 0 {
		return nil, exception.BadRequestDatabase
	}

	dbs, err := dao.SelectFromDatabaseByStruct(MapTagDB{}, fmt.Sprintf("map_name IN (%s)", db.Placeholders(len(mapNames))), util.ToGenericArray(mapNames)...)
	if err != nil {
		return nil, err
	}

	if dbs == nil || len(dbs) == 0 {
		return nil, exception.ResourceNotFoundDatabase
	}

	mapped := util.GroupBy(dbs, func(item MapTagDB) string {
		return item.MapName
	})

	return &mapped, nil

}
