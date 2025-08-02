package reviews

import (
	"github.com/identityofsine/fofx-go-gin-api-template/internal/constants/exception"
	"github.com/identityofsine/fofx-go-gin-api-template/internal/repository"
	"github.com/identityofsine/fofx-go-gin-api-template/internal/types/routeexception"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/db/dbmapper"
)

func GetReviewsByMapName(mapName string) ([]MapReview, routeexception.RouteError) {
	dbs, err := repository.GetMapReviewDBByMapName(mapName)
	if err != nil {
		return nil, routeexception.NewRouteError(
			err,
			"Error getting reviews by map name",
			"get-reviews-by-map-name",
			err.Code,
		)
	}
	reviews := dbmapper.MapAllDbFields[repository.MapReviewDB, MapReview](dbs)
	if reviews == nil {
		return nil, exception.InternalServerError
	}
	return *reviews, nil
}
