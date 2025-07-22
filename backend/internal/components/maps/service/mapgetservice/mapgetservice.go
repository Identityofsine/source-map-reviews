package mapgetservice

import (
	"github.com/identityofsine/fofx-go-gin-api-template/internal/components/maps/model/mapmodel"
	"github.com/identityofsine/fofx-go-gin-api-template/internal/components/maps/model/mapsearchform"
	"github.com/identityofsine/fofx-go-gin-api-template/internal/components/maps/service/maptaggetservice"
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

	rerr := populateAllMaps(maps)

	return maps, rerr
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
			continue
		}
		(*maps)[i].Tags = tags
	}

	return nil

}
