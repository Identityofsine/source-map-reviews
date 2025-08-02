package reviews

import (
	"github.com/identityofsine/fofx-go-gin-api-template/internal/components/user"
	"github.com/identityofsine/fofx-go-gin-api-template/internal/constants/exception"
	"github.com/identityofsine/fofx-go-gin-api-template/internal/repository"
	"github.com/identityofsine/fofx-go-gin-api-template/internal/types/routeexception"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/cookies"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/db/dbmapper"
)

// SaveReview handles both creating new reviews and updating existing ones
func SaveReview(review MapReview, cookies cookies.Cookies) (*MapReview, routeexception.RouteError) {
	currentUser, err := user.GetUserByCookies(&cookies)
	if err != nil || currentUser == nil {
		return nil, exception.BadRequest
	}

	// Validate required fields
	if review.MapName == "" {
		return nil, routeexception.NewRouteError(nil, "Map name is required", "map-name-required", exception.CODE_BAD_REQUEST)
	}

	if review.Stars < 1 || review.Stars > 5 {
		return nil, routeexception.NewRouteError(nil, "Stars must be between 1 and 5", "invalid-stars", exception.CODE_BAD_REQUEST)
	}

	// Ensure the reviewer ID matches the current user
	if review.ReviewerID != 0 && review.ReviewerID != currentUser.ID {
		return nil, routeexception.NewRouteError(nil, "Reviewer ID does not match current user", "reviewer-id-mismatch", exception.CODE_UNAUTHORIZED)
	}

	review.ReviewerID = currentUser.ID

	// Check if this is an update or insert
	if review.MapReviewID != 0 {
		// This is an update - verify the user owns this review
		existingReview, dbErr := repository.GetMapReviewDBById(review.MapReviewID)
		if dbErr != nil {
			return nil, routeexception.NewRouteError(dbErr, "Failed to fetch existing review", "fetch-review-failed", dbErr.Code)
		}

		if existingReview.ReviewerId != currentUser.ID {
			return nil, routeexception.NewRouteError(nil, "Cannot update review that doesn't belong to you", "unauthorized-update", exception.CODE_UNAUTHORIZED)
		}
	} else {
		// This is a new review - check if user already has a review for this map
		existingReview, dbErr := repository.GetMapReviewDBByMapNameAndReviewer(review.MapName, currentUser.ID)
		if dbErr != nil {
			return nil, routeexception.NewRouteError(dbErr, "Failed to check for existing review", "check-review-failed", dbErr.Code)
		}

		if existingReview != nil {
			return nil, routeexception.NewRouteError(nil, "You already have a review for this map. Use update instead.", "review-already-exists", exception.CODE_BAD_REQUEST)
		}
	}

	// Map to database model and save
	mapped := dbmapper.MapDbFields[MapReview, repository.MapReviewDB](review)
	savedReview, dbErr := repository.SaveMapReviewDB(*mapped)
	if dbErr != nil {
		return nil, routeexception.NewRouteError(dbErr, "Failed to save review", "save-review-failed", dbErr.Code)
	}

	// Map back to domain model
	result := dbmapper.MapDbFields[repository.MapReviewDB, MapReview](savedReview)
	return result, nil
}

// UpdateReview specifically handles updating an existing review
func UpdateReview(reviewId int64, review MapReview, cookies cookies.Cookies) (*MapReview, routeexception.RouteError) {
	currentUser, err := user.GetUserByCookies(&cookies)
	if err != nil || currentUser == nil {
		return nil, exception.BadRequest
	}

	// Verify the review exists and belongs to the current user
	existingReview, dbErr := repository.GetMapReviewDBById(reviewId)
	if dbErr != nil {
		return nil, routeexception.NewRouteError(dbErr, "Review not found", "review-not-found", dbErr.Code)
	}

	if existingReview.ReviewerId != currentUser.ID {
		return nil, routeexception.NewRouteError(nil, "Cannot update review that doesn't belong to you", "unauthorized-update", exception.CODE_UNAUTHORIZED)
	}

	// Set the review ID and reviewer ID
	review.MapReviewID = reviewId
	review.ReviewerID = currentUser.ID
	review.MapName = existingReview.MapName // Preserve original map name

	// Validate stars if provided
	if review.Stars < 1 || review.Stars > 5 {
		return nil, routeexception.NewRouteError(nil, "Stars must be between 1 and 5", "invalid-stars", exception.CODE_BAD_REQUEST)
	}

	// Map to database model and save
	mapped := dbmapper.MapDbFields[MapReview, repository.MapReviewDB](review)
	savedReview, dbErr := repository.SaveMapReviewDB(*mapped)
	if dbErr != nil {
		return nil, routeexception.NewRouteError(dbErr, "Failed to update review", "update-review-failed", dbErr.Code)
	}

	// Map back to domain model
	result := dbmapper.MapDbFields[repository.MapReviewDB, MapReview](savedReview)
	return result, nil
}

// DeleteReview handles deleting a review
func DeleteReview(reviewId int64, cookies cookies.Cookies) routeexception.RouteError {
	currentUser, err := user.GetUserByCookies(&cookies)
	if err != nil || currentUser == nil {
		return exception.BadRequest
	}

	// Verify the review exists and belongs to the current user
	existingReview, dbErr := repository.GetMapReviewDBById(reviewId)
	if dbErr != nil {
		return routeexception.NewRouteError(dbErr, "Review not found", "review-not-found", dbErr.Code)
	}

	if existingReview.ReviewerId != currentUser.ID {
		return routeexception.NewRouteError(nil, "Cannot delete review that doesn't belong to you", "unauthorized-delete", exception.CODE_UNAUTHORIZED)
	}

	// Delete the review
	dbErr = repository.DeleteMapReviewDB(reviewId)
	if dbErr != nil {
		return routeexception.NewRouteError(dbErr, "Failed to delete review", "delete-review-failed", dbErr.Code)
	}

	return nil
}

// GetReviewsByUser returns all reviews by a specific user
func GetReviewsByUser(userId int64) ([]MapReview, routeexception.RouteError) {
	reviews, dbErr := repository.GetMapReviewDBByReviewer(userId)
	if dbErr != nil {
		return nil, routeexception.NewRouteError(dbErr, "Failed to fetch user reviews", "fetch-user-reviews-failed", dbErr.Code)
	}

	var result []MapReview
	for _, review := range reviews {
		mapped := dbmapper.MapDbFields[repository.MapReviewDB, MapReview](review)
		result = append(result, *mapped)
	}

	return result, nil
}

// GetReviewByUserAndMap returns a user's review for a specific map
func GetReviewByUserAndMap(userId int64, mapName string) (*MapReview, routeexception.RouteError) {
	review, dbErr := repository.GetMapReviewDBByMapNameAndReviewer(mapName, userId)
	if dbErr != nil {
		return nil, routeexception.NewRouteError(dbErr, "Failed to fetch review", "fetch-review-failed", dbErr.Code)
	}

	if review == nil {
		return nil, nil // No review found, but not an error
	}

	result := dbmapper.MapDbFields[repository.MapReviewDB, MapReview](*review)
	return result, nil
}
