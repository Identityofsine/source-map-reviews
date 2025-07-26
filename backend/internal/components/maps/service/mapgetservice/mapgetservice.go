package mapgetservice

import (
	"github.com/identityofsine/fofx-go-gin-api-template/internal/components/maps/model/mapmodel"
	"github.com/identityofsine/fofx-go-gin-api-template/internal/components/maps/model/mapsearchform"
	"github.com/identityofsine/fofx-go-gin-api-template/internal/components/maps/model/maptags"
	"github.com/identityofsine/fofx-go-gin-api-template/internal/components/maps/service/maptaggetservice"
	"github.com/identityofsine/fofx-go-gin-api-template/internal/constants/exception"
	"github.com/identityofsine/fofx-go-gin-api-template/internal/repository/model/mapdb"
	"github.com/identityofsine/fofx-go-gin-api-template/internal/types/routeexception"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/db/dbmapper"
)

func GetMaps() (*[]mapmodel.Map, routeexception.RouteError) {

	dbs, err := mapdb.GetMaps()
	if err != nil {
		return nil, routeexception.NewRouteError(
			err,
			"Failed to get maps",
			"get-maps-failed",
			err.Code,
		)
	}

	maps := dbmapper.MapAllDbFields[mapdb.MapDb, mapmodel.Map](*dbs)

	if dbs == nil || len(*dbs) == 0 {
		return &[]mapmodel.Map{}, nil // Return empty slice if no maps found
	}

	rerr := populateAllMaps(maps)

	return maps, rerr
}

func GetMap(mapName string) (*mapmodel.Map, routeexception.RouteError) {

	db, err := mapdb.GetMap(mapName)
	if err != nil {
		if err.Code == exception.CODE_RESOURCE_NOT_FOUND {
			return nil, exception.ResourceNotFound
		}
		return nil, routeexception.NewRouteError(
			err,
			"Failed to get map",
			"get-map-failed",
			err.Code,
		)
	}

	mapObj := dbmapper.MapDbFields[mapdb.MapDb, mapmodel.Map](*db)
	if mapObj == nil {
		return nil, routeexception.NewRouteError(
			err,
			"Failed to map database fields",
			"map-db-fields-failed",
			exception.CODE_INTERNAL_SERVER_ERROR,
		)
	}

	rerr := populateMap(mapObj)

	return mapObj, rerr
}

func SearchMaps(form mapsearchform.MapSearchForm) (*[]mapmodel.Map, routeexception.RouteError) {

	dbs, err := mapdb.SearchMaps(form)
	if err != nil {
		return nil, routeexception.NewRouteError(
			err,
			"Failed to search maps",
			"search-maps-failed",
			err.Code,
		)
	}

	if dbs == nil || len(*dbs) == 0 {
		return &[]mapmodel.Map{}, nil // Return empty slice if no maps found
	}

	maps := dbmapper.MapAllDbFields[mapdb.MapDb, mapmodel.Map](*dbs)

	rerr := populateAllMaps(maps)

	return maps, rerr
}

func populateAllMaps(maps *[]mapmodel.Map) routeexception.RouteError {
	tagsMap, err := maptaggetservice.GetTagsByMaps(
		*maps,
	)
	if err != nil {
		return err
	}
	for i, mapObj := range *maps {
		tags, ok := tagsMap[mapObj.MapName]
		if !ok {
			(*maps)[i].Tags = []maptags.MapTags{} // No tags found for this map
			continue
		}
		(*maps)[i].Tags = tags
	}

	return nil

}

func populateMap(
	mapObj *mapmodel.Map,
) routeexception.RouteError {

	maps := []mapmodel.Map{*mapObj} // only one element, copied
	err := populateAllMaps(&maps)
	if err != nil {
		return err
	}

	// Copy updated tags back into original pointer
	mapObj.Tags = maps[0].Tags

	return nil
}
