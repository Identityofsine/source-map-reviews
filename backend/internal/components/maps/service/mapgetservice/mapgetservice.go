package mapgetservice

import (
	"github.com/identityofsine/fofx-go-gin-api-template/internal/components/maps/model/mapmodel"
	"github.com/identityofsine/fofx-go-gin-api-template/internal/components/maps/model/maptags"
	"github.com/identityofsine/fofx-go-gin-api-template/internal/repository/model/lk_tagdb"
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

	newMaps, err := populateAllMaps(&maps)

	return newMaps, err
}

func populateAllMaps(maps *[]mapmodel.Map) (*[]mapmodel.Map, db.DatabaseError) {
	newMaps := make([]mapmodel.Map, len(*maps))
	for i, mapobj := range *maps {
		populated, err := populateMap(&mapobj)
		if err != nil {
			return maps, err
		}
		newMaps[i] = *populated
	}
	return &newMaps, nil
}

func populateMap(mapobj *mapmodel.Map) (*mapmodel.Map, db.DatabaseError) {

	if mapobj.Tags == nil || len(mapobj.Tags) == 0 {
		return mapobj, nil
	}

	tagNames := util.Map[maptags.MapTags, string](mapobj.Tags, func(item maptags.MapTags) string {
		return item.TagName
	})

	dbs, err := lk_tagdb.GetLkTagsByLkTags(tagNames)
	if err != nil {
		return nil, err
	}

	if dbs == nil || len(*dbs) == 0 {
		return mapobj, nil
	}

	tagDbs := util.MapBy(
		*dbs,
		func(item lk_tagdb.LkTagDb) string {
			return item.LkTag
		},
		func(item lk_tagdb.LkTagDb) lk_tagdb.LkTagDb {
			return item
		},
	)

	for i, tag := range mapobj.Tags {
		tagDb, ok := tagDbs[tag.TagName]
		if !ok {
			continue
		}
		mapobj.Tags[i].TagDescription = tagDb.Description
		mapobj.Tags[i].TagDescriptionShort = tagDb.ShortDescription
	}

	return mapobj, nil
}
