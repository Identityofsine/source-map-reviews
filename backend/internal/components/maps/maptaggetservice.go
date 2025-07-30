package maps

import (
	"fmt"

	"github.com/identityofsine/fofx-go-gin-api-template/internal/constants/exception"
	"github.com/identityofsine/fofx-go-gin-api-template/internal/repository"
	"github.com/identityofsine/fofx-go-gin-api-template/internal/types/routeexception"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/db/dbmapper"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/storedlogs"
	"github.com/identityofsine/fofx-go-gin-api-template/util"
)

type MapTagLkDb = repository.LkTagDB
type MapTagLkDbSlice = []repository.LkTagDB
type MapTagDb = repository.MapTagDb

type MapTagRelationship map[string][]MapTag

func CastMapTagRelationship(m map[string][]MapTagDb) MapTagRelationship {
	tags := make(MapTagRelationship, len(m))
	for mapName, tag := range m {
		tags[mapName] = *dbmapper.MapAllDbFields[MapTagDb, MapTag](tag)
	}
	return tags
}

func ReverseCastMapTagRelationship(m MapTagRelationship) map[string][]MapTagDb {
	tags := make(map[string][]MapTagDb, len(m))
	for mapName, tag := range m {
		tags[mapName] = *dbmapper.MapAllDbFields[MapTag, MapTagDb](tag)
	}
	return tags
}

func GetTagsByMaps(mapsObject []Map) (MapTagRelationship, routeexception.RouteError) {

	var maps []Map = make([]Map, len(mapsObject))
	for _, mapObj := range mapsObject {
		maps = append(maps, Map(mapObj))
	}

	mapNames := util.Map[Map, string](maps, func(item Map) string {
		return item.MapName
	})

	tags, err := GetMapTagsByMapNames(mapNames)
	if err != nil {
		return nil, routeexception.NewRouteError(
			err,
			"Failed to get tags for maps",
			"get-tags-failed",
			err.Code,
		)
	}

	if tags == nil || len(*tags) == 0 {
		return nil, exception.ResourceNotFound
	}

	tagsModelMap := util.MapToMap[[]MapTagDb, []MapTag](*tags, func(item []MapTagDb) []MapTag {
		return *dbmapper.MapAllDbFields[MapTagDb, MapTag](item)
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
