package reviews

import (
	"fmt"

	"github.com/identityofsine/fofx-go-gin-api-template/internal/constants/exception"
	"github.com/identityofsine/fofx-go-gin-api-template/internal/repository"
	"github.com/identityofsine/fofx-go-gin-api-template/internal/types/routeexception"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/db/dbmapper"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/storedlogs"
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

	for i, review := range *reviews {
		images, err := repository.GetMapReviewImagesByReviewId(review.MapReviewID)
		if err != nil {
			storedlogs.LogWarn(fmt.Sprintf("Error fetching images for review %d: %v", review.MapReviewID, err))
			continue // Skip this review if images cannot be fetched
		}
		if images != nil && len(images) > 0 {
			imgModels := dbmapper.MapAllDbFields[repository.MapReviewImageDB, MapReviewImage](images)
			if imgModels == nil {
				return nil, exception.InternalServerError
			}
			err := (*reviews)[i].Populate(*imgModels)
			if err != nil {
				if err.Code == exception.CODE_INTERNAL_SERVER_ERROR {
					return nil, err
				}
				continue
			}
		} else {
			storedlogs.LogInfo(fmt.Sprintf("No images found for review %d", review.MapReviewID))
		}
	}

	return *reviews, nil
}
