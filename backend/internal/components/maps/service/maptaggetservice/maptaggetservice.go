package maptaggetservice

import (
	"fmt"

	"github.com/identityofsine/fofx-go-gin-api-template/internal/components/maps/model/mapmodel"
	"github.com/identityofsine/fofx-go-gin-api-template/internal/components/maps/model/maptags"
	"github.com/identityofsine/fofx-go-gin-api-template/internal/constants/exception"
	"github.com/identityofsine/fofx-go-gin-api-template/internal/repository/model/lktagdb"
	"github.com/identityofsine/fofx-go-gin-api-template/internal/repository/model/maptagdb"
	"github.com/identityofsine/fofx-go-gin-api-template/internal/types/routeexception"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/db/dbmapper"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/storedlogs"
	"github.com/identityofsine/fofx-go-gin-api-template/util"
)

type Map = mapmodel.Map
type MapTag = maptags.MapTags
type MapTagLkDb = lktagdb.LkTagDb
type MapTagLkDbSlice = []lktagdb.LkTagDb
type MapTagDb = maptagdb.MapTagDb

type MapTagRelationship map[string][]MapTag

func CastMapTagRelationship(m map[string][]maptagdb.MapTagDb) MapTagRelationship {
	tags := make(MapTagRelationship, len(m))
	for mapName, tag := range m {
		tags[mapName] = *dbmapper.MapAllDbFields[maptagdb.MapTagDb, MapTag](tag)
	}
	return tags
}

func ReverseCastMapTagRelationship(m MapTagRelationship) map[string][]maptagdb.MapTagDb {
	tags := make(map[string][]maptagdb.MapTagDb, len(m))
	for mapName, tag := range m {
		tags[mapName] = *dbmapper.MapAllDbFields[MapTag, maptagdb.MapTagDb](tag)
	}
	return tags
}

func GetTagsByMaps(mapsObject []mapmodel.Map) (MapTagRelationship, routeexception.RouteError) {

	var maps []Map = make([]Map, len(mapsObject))
	for _, mapObj := range mapsObject {
		maps = append(maps, Map(mapObj))
	}

	mapNames := util.Map[Map, string](maps, func(item Map) string {
		return item.MapName
	})

	tags, err := maptagdb.GetMapTagsByMapNames(mapNames)
	if err != nil {
		return nil, routeexception.NewRouteError(
			err,
			"Failed to get tags for maps",
			"get-tags-failed",
			err.Code,
		)
	}

	if len(*tags) == 0 {
		return nil, exception.ResourceNotFound
	}

	tagsModelMap := util.MapToMap[[]maptagdb.MapTagDb, []MapTag](*tags, func(item []maptagdb.MapTagDb) []MapTag {
		return *dbmapper.MapAllDbFields[maptagdb.MapTagDb, MapTag](item)
	})

	for mapName, tag := range tagsModelMap {
		tag, err := populateMapTags(tag)
		if err != nil {
			storedlogs.LogWarn(fmt.Sprintf("Failed to populate tags for map %s: %v", mapName, err))
			continue
		}
		tagsModelMap[mapName] = tag
	}

	return tagsModelMap, nil

}

func populateMapTags(tag []MapTag) ([]MapTag, routeexception.RouteError) {

	tag = util.Filter(tag, func(item MapTag) bool {
		if item.TagName == "" {
			return false
		}
		return true
	})

	tagLks, err := getTagLks(tag)
	if err != nil {
		return nil, err
	}

	lkMap := util.MapBy(tagLks,
		func(item lktagdb.LkTagDb) string {
			return item.LkTag
		}, func(item lktagdb.LkTagDb) MapTagLkDb {
			return MapTagLkDb(item)
		})

	for i := range tag {
		if lk, ok := lkMap[tag[i].TagName]; ok {
			tag[i].TagDescription = lk.Description
			tag[i].TagDescriptionShort = lk.ShortDescription
		}
	}

	return tag, nil
}

func getTagLks(tags []MapTag) (MapTagLkDbSlice, routeexception.RouteError) {

	if len(tags) == 0 {
		return nil, routeexception.NewRouteError(
			nil,
			"Map tag is empty",
			"map-tag-empty",
			exception.CODE_BAD_REQUEST,
		)
	}

	tagLks := util.Map(tags, func(item MapTag) string {
		return item.TagName
	})

	lks, err := lktagdb.GetLkTagsByLkTags(tagLks)
	if err != nil {
		return nil, routeexception.NewRouteError(
			err,
			"Failed to get tags for maps",
			"get-tags-failed",
			err.Code,
		)
	}

	if lks == nil || len(*lks) == 0 {
		return nil, routeexception.NewRouteError(
			nil,
			"Map tag not found",
			"map-tag-not-found",
			exception.CODE_RESOURCE_NOT_FOUND,
		)
	}

	return *lks, nil
}
