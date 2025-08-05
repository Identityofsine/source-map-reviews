package reviews

import (
	"strconv"

	"github.com/identityofsine/fofx-go-gin-api-template/internal/constants/exception"
	"github.com/identityofsine/fofx-go-gin-api-template/internal/repository"
	"github.com/identityofsine/fofx-go-gin-api-template/internal/types/routeexception"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/db/dbmapper"
	"github.com/identityofsine/fofx-go-gin-api-template/util"
)

func GetReviewImagesByMapReviewId(mapReviewId int64) ([]MapReviewImage, error) {
	dbs, err := repository.GetMapReviewImagesByReviewId(mapReviewId)
	if err != nil {
		return nil, err
	}

	mapped := dbmapper.MapAllDbFields[repository.MapReviewImageDB, MapReviewImage](dbs)
	if mapped == nil {
		return nil, exception.InternalServerError
	}

	return *mapped, nil
}

func GetReviewImageByMapReview(mapReview MapReviewImage) ([]MapReviewImage, error) {
	if mapReview.MapReviewId == 0 {
		return nil, exception.ResourceNotFound
	}

	dbs, err := repository.GetMapReviewImagesByReviewId(mapReview.MapReviewId)
	if err != nil {
		return nil, err
	}

	mapped := dbmapper.MapAllDbFields[repository.MapReviewImageDB, MapReviewImage](dbs)
	if mapped == nil {
		return nil, exception.InternalServerError
	}

	return *mapped, nil
}

func GetReviewImagesByMapReviewIds(mapReviewIds []int64) (map[string]MapReviewImage, routeexception.RouteError) {
	if len(mapReviewIds) == 0 {
		return nil, exception.ResourceNotFound
	}

	dbs, err := repository.GetMapReviewImagesByReviewIds(mapReviewIds)
	if err != nil {
		return nil, routeexception.NewRouteError(
			err,
			"Failed to get review images by review IDs",
			"get-review-images-by-review-ids-failed",
			err.Code,
		)
	}

	mapped := dbmapper.MapAllDbFields[repository.MapReviewImageDB, MapReviewImage](dbs)
	if mapped == nil {
		return nil, exception.InternalServerError
	}

	mapModel := util.MapBy(
		*mapped,
		func(item MapReviewImage) string {
			return strconv.Itoa(int(item.ImageId))
		},
		func(item MapReviewImage) MapReviewImage {
			return item
		},
	)

	return mapModel, nil
}
