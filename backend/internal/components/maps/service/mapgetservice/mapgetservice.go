package mapgetservice

import (
	"github.com/identityofsine/fofx-go-gin-api-template/internal/components/maps/model/mapmodel"
	"github.com/identityofsine/fofx-go-gin-api-template/internal/components/maps/model/maptags"
	"github.com/identityofsine/fofx-go-gin-api-template/internal/repository/model/mapdb"
	"github.com/identityofsine/fofx-go-gin-api-template/internal/repository/model/maptagdb"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/db"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/db/dbmapper"
	"github.com/identityofsine/fofx-go-gin-api-template/util"
)

func GetMaps() (*[]mapmodel.Map, db.DatabaseError) {

	dbs, err := mapdb.GetMaps()
	if err != nil {
		return nil, err
	}

	mapNames := util.Map[mapdb.MapDb, string](*dbs, func(item mapdb.MapDb) string {
		return item.MapName
	})

	tagsMap, err := maptagdb.GetMapTagsByMapNames(mapNames)
	if err != nil {
		return nil, err
	}

	maps := make([]mapmodel.Map, len(mapNames))

	for i, dbMap := range *dbs {

		if tags, ok := (*tagsMap)[dbMap.MapName]; ok {
			newTags := dbmapper.MapAllDbFields[maptagdb.MapTagDb, maptags.MapTags](tags)
			mutatedMap := dbmapper.MapDbFullFields[mapdb.MapDb, mapmodel.Map](dbMap, newTags)
			if mutatedMap == nil {
				return nil, db.NewDatabaseError("GetMaps", "Failed to map db fields to model", "mapper-failed", 500)
			}
			maps[i] = *mutatedMap
		} else {
			maps[i] = *dbmapper.MapDbFields[mapdb.MapDb, mapmodel.Map](dbMap)
		}

	}

	return &maps, nil
}

func populateAllMaps(maps *[]mapmodel.Map) (*[]mapmodel.Map, db.DatabaseError) {
	newMaps := make([]mapmodel.Map, len(*maps))
	for _, mapobj := range *maps {
		mapobj, err := populateMap(mapobj)
		newMaps = append(newMaps, *mapobj)
		if err != nil {
			return nil, err
		}
	}
	return &newMaps, nil
}

func populateMap(mapobj mapmodel.Map) (*mapmodel.Map, db.DatabaseError) {

	dbs, err := maptagdb.GetMapTagsByMapName(mapobj.MapName)
	if err != nil {
		return nil, err
	}

	mapTags := dbmapper.MapAllDbFields[maptagdb.MapTagDb, maptags.MapTags](*dbs)

	mapobj.Tags = *mapTags

	return &mapobj, nil
}
