package maps

import (
	"github.com/identityofsine/fofx-go-gin-api-template/internal/components/maps/mapsearchform"
	"github.com/identityofsine/fofx-go-gin-api-template/internal/constants/exception"
	"github.com/identityofsine/fofx-go-gin-api-template/internal/repository"
	"github.com/identityofsine/fofx-go-gin-api-template/internal/types/routeexception"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/db/dbmapper"
)

func GetMaps() (*[]Map, routeexception.RouteError) {

	dbs, err := repository.GetMaps()
	if err != nil {
		return nil, routeexception.NewRouteError(
			err,
			"Failed to get maps",
			"get-maps-failed",
			err.Code,
		)
	}

	maps := dbmapper.MapAllDbFields[repository.MapDB, Map](*dbs)

	if dbs == nil || len(*dbs) == 0 {
		return &[]Map{}, nil // Return empty slice if no maps found
	}

	rerr := populateAllMaps(maps)

	return maps, rerr
}

func GetMap(mapName string) (*Map, routeexception.RouteError) {

	db, err := repository.GetMap(mapName)
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

	mapObj := dbmapper.MapDbFields[repository.MapDB, Map](*db)
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

func SearchMaps(form mapsearchform.MapSearchForm) (*[]Map, routeexception.RouteError) {

	dbs, err := repository.SearchMaps(form)
	if err != nil {
		return nil, routeexception.NewRouteError(
			err,
			"Failed to search maps",
			"search-maps-failed",
			err.Code,
		)
	}

	if dbs == nil || len(*dbs) == 0 {
		return &[]Map{}, nil // Return empty slice if no maps found
	}

	maps := dbmapper.MapAllDbFields[repository.MapDB, Map](*dbs)

	rerr := populateAllMaps(maps)

	return maps, rerr
}

func populateAllMaps(maps *[]Map) routeexception.RouteError {
	tagsMap, err := GetTagsByMaps(
		*maps,
	)
	if err != nil {
		if err.Code == exception.CODE_RESOURCE_NOT_FOUND {
			// No tags found for any maps, return without error
			return nil
		}
		return err
	}
	for i, mapObj := range *maps {
		tags, ok := tagsMap[mapObj.MapName]
		if !ok {
			(*maps)[i].Tags = []MapTag{} // No tags found for this map
			continue
		}
		(*maps)[i].Tags = tags
	}

	return nil

}

func populateMap(
	mapObj *Map,
) routeexception.RouteError {

	maps := make([]Map, 1)
	maps[0] = *mapObj // Create a slice with a single map object
	err := populateAllMaps(&maps)
	if err != nil {
		return err
	}

	// Copy updated tags back into original pointer
	mapObj.Tags = maps[0].Tags

	return nil
}
