package reviews

import (
	"github.com/identityofsine/fofx-go-gin-api-template/internal/components/user"
	"github.com/identityofsine/fofx-go-gin-api-template/internal/constants/exception"
	"github.com/identityofsine/fofx-go-gin-api-template/internal/repository"
	"github.com/identityofsine/fofx-go-gin-api-template/internal/types/routeexception"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/cookies"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/db/dbmapper"
)

func SaveReview(review MapReview, cookies cookies.Cookies) (*MapReview, routeexception.RouteError) {
	currentUser, err := user.GetUserByCookies(&cookies)
	if err != nil || currentUser == nil {
		return nil, exception.BadRequest
	}

	if review.MapName == "" {
		return nil, routeexception.NewRouteError(nil, "Map name is required", "map-name-required", exception.CODE_BAD_REQUEST)
	}

	reviewerId := review.ReviewerID

	if reviewerId != 0 && reviewerId != currentUser.ID {
		return nil, routeexception.NewRouteError(nil, "Reviewer ID does not match current user", "reviewer-id-mismatch", exception.CODE_UNAUTHORIZED)
	}

	review.ReviewerID = currentUser.ID

	mapped := dbmapper.MapDbFields[MapReview, repository.MapReviewDB](review)
	if review, err := repository.SaveMapReviewDB(*mapped); err != nil {
		return nil, routeexception.NewRouteError(err, "Failed to save review", "save-review-failed", err.Code)
	} else {
		return dbmapper.MapDbFields[repository.MapReviewDB, MapReview](review), nil
	}
}
