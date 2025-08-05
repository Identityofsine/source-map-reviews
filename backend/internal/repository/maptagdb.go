package repository

import (
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

	whereClause := "map_name IN (" + db.Placeholders(len(mapNames)) + ")"

	mapNameMutated := util.ToGenericArray(mapNames...)

	dbs, err := dao.SelectFromDatabaseByStruct(MapTagDB{}, whereClause, mapNameMutated...)
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

// SaveMapTagDb handles inserting map-tag relationships
func SaveMapTagDb(mapTag MapTagDB) (*MapTagDB, db.DatabaseError) {
	// Check if this map-tag combination already exists
	existing, err := GetMapTagsByMapName(mapTag.MapName)
	if err == nil && existing != nil {
		for _, existingTag := range *existing {
			if existingTag.LkTag == mapTag.LkTag {
				// Already exists, return the existing one
				return &existingTag, nil
			}
		}
	}

	// Insert new map-tag relationship
	inserted, err := dao.InsertIntoDatabaseByStruct(mapTag)
	if err != nil {
		return nil, err
	}
	return inserted, nil
}

// DeleteMapTagDb deletes a map-tag relationship
func DeleteMapTagDb(mapName string, lkTag string) db.DatabaseError {
	_, err := db.Delete("DELETE FROM "+map_table+" WHERE map_name = $1 AND lk_tag = $2", mapName, lkTag)
	return err
}

// DeleteMapTagsByMapName deletes all tags for a specific map
func DeleteMapTagsByMapName(mapName string) db.DatabaseError {
	_, err := db.Delete("DELETE FROM "+map_table+" WHERE map_name = $1", mapName)
	return err
}
