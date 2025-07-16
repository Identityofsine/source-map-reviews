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

	// Flatten all map values into a single slice
	var allTags []maptagdb.MapTagDb
	for _, tagSlice := range *tagsMap {
		allTags = append(allTags, tagSlice...)
	}
	tagsList := util.GroupIntoLists[maptagdb.MapTagDb](allTags, func(item maptagdb.MapTagDb) string {
		return item.MapName
	})

	// Map the database models to the application models
	maps := dbmapper.MapAllDbFullFields[mapdb.MapDb, mapmodel.Map](*dbs, tagsList)
	if maps == nil {
		return nil, db.NewDatabaseError("GetMaps", "Mapping failed", "mapping-failed", 500)
	}

	return maps, nil
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
