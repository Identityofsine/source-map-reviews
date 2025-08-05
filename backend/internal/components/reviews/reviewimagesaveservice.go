package reviews

import (
	"github.com/identityofsine/fofx-go-gin-api-template/internal/constants/exception"
	"github.com/identityofsine/fofx-go-gin-api-template/internal/repository"
	"github.com/identityofsine/fofx-go-gin-api-template/internal/types/routeexception"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/db/dbmapper"
)

func SaveReviewImage(reviewId int64, imageId int64) (*MapReviewImage, routeexception.RouteError) {
	if reviewId <= 0 || imageId <= 0 {
		return nil, routeexception.NewRouteError(
			nil,
			"Invalid review ID or image ID",
			"invalid-review-image-ids",
			400,
		)
	}

	reviewImage := repository.MapReviewImageDB{
		MapReviewId: reviewId,
		ImageId:     imageId,
	}

	savedImage, err := repository.SaveMapReviewImageDb(reviewImage)
	if err != nil {
		return nil, routeexception.NewRouteError(
			err,
			"Failed to save review image",
			"save-review-image-failed",
			err.Code,
		)
	}
	if savedImage == nil {
		return nil, routeexception.NewRouteError(
			nil,
			"Failed to save review image, no data returned",
			"save-review-image-no-data",
			exception.CODE_INTERNAL_SERVER_ERROR,
		)
	}

	mapReviewImage := dbmapper.MapDbFields[repository.MapReviewImageDB, MapReviewImage](*savedImage)

	return mapReviewImage, nil
}
