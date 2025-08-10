package reviews

import (
	"github.com/identityofsine/fofx-go-gin-api-template/internal/constants/exception"
	"github.com/identityofsine/fofx-go-gin-api-template/internal/repository"
	"github.com/identityofsine/fofx-go-gin-api-template/internal/types/routeexception"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/db/dbmapper"
)

// GetReviewsByMapNames returns all reviews for multiple maps with their images populated efficiently
func GetReviewsByMapNames(mapNames []string) (map[string][]MapReview, routeexception.RouteError) {
	if mapNames == nil || len(mapNames) == 0 {
		return make(map[string][]MapReview), nil
	}

	// Get all reviews for the maps in bulk
	reviewsByMapDB, err := repository.GetMapReviewDBByMapNames(mapNames)
	if err != nil {
		return nil, routeexception.NewRouteError(
			err,
			"Error getting reviews by map names",
			"get-reviews-by-map-names",
			err.Code,
		)
	}

	// Convert to domain objects
	reviewsByMap := make(map[string][]MapReview)
	var allReviews []MapReview

	for mapName, dbReviews := range reviewsByMapDB {
		if len(dbReviews) == 0 {
			reviewsByMap[mapName] = []MapReview{}
			continue
		}

		reviews := dbmapper.MapAllDbFields[repository.MapReviewDB, MapReview](dbReviews)
		if reviews == nil {
			return nil, exception.InternalServerError
		}

		reviewsByMap[mapName] = *reviews
		allReviews = append(allReviews, *reviews...)
	}

	// Populate all reviews with images using efficient bulk queries
	if len(allReviews) > 0 {
		populationErr := PopulateReviewsWithImages(allReviews)
		if populationErr != nil {
			return nil, routeexception.NewRouteError(
				populationErr,
				"Failed to populate reviews with images",
				"populate-reviews-with-images-failed",
				populationErr.Code,
			)
		}

		// Update the map with the populated reviews
		reviewIndex := 0
		for mapName, mapReviews := range reviewsByMap {
			if len(mapReviews) > 0 {
				for i := range mapReviews {
					reviewsByMap[mapName][i] = allReviews[reviewIndex]
					reviewIndex++
				}
			}
		}
	}

	return reviewsByMap, nil
}

// GetReviewsByMapName returns all reviews for a specific map with their images populated
func GetReviewsByMapName(mapName string) ([]MapReview, routeexception.RouteError) {
	// Get reviews from database
	dbs, err := repository.GetMapReviewDBByMapName(mapName)
	if err != nil {
		return nil, routeexception.NewRouteError(
			err,
			"Error getting reviews by map name",
			"get-reviews-by-map-name",
			err.Code,
		)
	}

	// Map to domain objects
	reviews := dbmapper.MapAllDbFields[repository.MapReviewDB, MapReview](dbs)
	if reviews == nil {
		return nil, exception.InternalServerError
	}

	// Populate with images using efficient bulk queries
	populationErr := PopulateReviewsWithImages(*reviews)
	if populationErr != nil {
		return nil, routeexception.NewRouteError(
			populationErr,
			"Failed to populate reviews with images",
			"populate-reviews-with-images-failed",
			populationErr.Code,
		)
	}

	return *reviews, nil
}

// GetReviewsByUser returns all reviews by a specific user with their images populated
func GetReviewsByUser(userId int64) ([]MapReview, routeexception.RouteError) {
	// Get reviews from database
	dbs, err := repository.GetMapReviewDBByReviewer(userId)
	if err != nil {
		return nil, routeexception.NewRouteError(
			err,
			"Failed to fetch user reviews",
			"fetch-user-reviews-failed",
			err.Code,
		)
	}

	// Map to domain objects
	reviews := make([]MapReview, len(dbs))
	for i, db := range dbs {
		mapped := dbmapper.MapDbFields[repository.MapReviewDB, MapReview](db)
		if mapped == nil {
			return nil, exception.InternalServerError
		}
		reviews[i] = *mapped
	}

	// Populate with images using efficient bulk queries
	populationErr := PopulateReviewsWithImages(reviews)
	if populationErr != nil {
		return nil, routeexception.NewRouteError(
			populationErr,
			"Failed to populate reviews with images",
			"populate-reviews-with-images-failed",
			populationErr.Code,
		)
	}

	return reviews, nil
}

// GetReviewByUserAndMap returns a user's review for a specific map with images populated
func GetReviewByUserAndMap(userId int64, mapName string) (*MapReview, routeexception.RouteError) {
	// Get review from database
	db, err := repository.GetMapReviewDBByMapNameAndReviewer(mapName, userId)
	if err != nil {
		return nil, routeexception.NewRouteError(
			err,
			"Failed to fetch review",
			"fetch-review-failed",
			err.Code,
		)
	}

	if db == nil {
		return nil, nil // No review found, but not an error
	}

	// Map to domain object
	review := dbmapper.MapDbFields[repository.MapReviewDB, MapReview](*db)
	if review == nil {
		return nil, exception.InternalServerError
	}

	// Populate with images
	populationErr := PopulateReviewWithImages(review)
	if populationErr != nil {
		return nil, routeexception.NewRouteError(
			populationErr,
			"Failed to populate review with images",
			"populate-review-with-images-failed",
			populationErr.Code,
		)
	}

	return review, nil
}

// GetReviewById returns a review by its ID with images populated
func GetReviewById(reviewId int64) (*MapReview, routeexception.RouteError) {
	// Get review from database
	db, err := repository.GetMapReviewDBById(reviewId)
	if err != nil {
		return nil, routeexception.NewRouteError(
			err,
			"Failed to fetch review",
			"fetch-review-failed",
			err.Code,
		)
	}

	if db == nil {
		return nil, nil // No review found, but not an error
	}

	// Map to domain object
	review := dbmapper.MapDbFields[repository.MapReviewDB, MapReview](*db)
	if review == nil {
		return nil, exception.InternalServerError
	}

	// Populate with images
	populationErr := PopulateReviewWithImages(review)
	if populationErr != nil {
		return nil, routeexception.NewRouteError(
			populationErr,
			"Failed to populate review with images",
			"populate-review-with-images-failed",
			populationErr.Code,
		)
	}

	return review, nil
}
