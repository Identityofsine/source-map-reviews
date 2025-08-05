package reviews

import (
	"github.com/identityofsine/fofx-go-gin-api-template/internal/components/images"
	"github.com/identityofsine/fofx-go-gin-api-template/internal/constants/exception"
	"github.com/identityofsine/fofx-go-gin-api-template/internal/types/routeexception"
)

type MapReviewImage struct {
	MapReviewImageId int64        `json:"mapReviewImageId" db:"map_review_image_id"`
	MapReviewId      int64        `json:"mapReviewId" db:"map_review_id"`
	Image            images.Image `json:"image"` // Assuming Image is a struct defined in the images package
}

func (m *MapReviewImage) Populate(img *images.Image) routeexception.RouteError {

	if img == nil || img.ImageID == 0 {
		return routeexception.NewRouteError(
			nil,
			"Image ID is zero, cannot populate image",
			"populate-map-review-image-failed",
			exception.CODE_BAD_REQUEST,
		)
		// No image to populate
	}

	m.Image = *img

	return nil
}
