package reviews

import (
	"strconv"
	"time"

	"github.com/identityofsine/fofx-go-gin-api-template/internal/components/images"
	"github.com/identityofsine/fofx-go-gin-api-template/internal/constants/exception"
	"github.com/identityofsine/fofx-go-gin-api-template/internal/repository"
	"github.com/identityofsine/fofx-go-gin-api-template/internal/types/routeexception"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/db/dbmapper"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/storedlogs"
	"github.com/identityofsine/fofx-go-gin-api-template/util"
)

type MapReview struct {
	MapReviewID int64  `json:"mapReviewId" db:"map_review_id" dao:"pk"` // primary key
	MapName     string `json:"mapName" db:"map_name" binding:"required"`
	ReviewerID  int64  `json:"userId" db:"reviewer" binding:"required"`

	Images []MapReviewImage `json:"images"` // array of MapImage

	Stars     int       `json:"stars" db:"stars" binding:"required"`
	Review    string    `json:"review" db:"review" binding:"required"`
	CreatedAt time.Time `json:"createdAt" db:"created_at" dao:"omit"` // time when the review was created
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at" dao:"omit"` // time when the review was last updated
}

func (m *MapReview) Populate(imgs []MapReviewImage) routeexception.RouteError {

	if imgs == nil || len(imgs) == 0 {
		return routeexception.NewRouteError(nil, "No imaged were provided", "no-images-provided", exception.CODE_BAD_REQUEST)
	}

	imgIds := util.Map(imgs, func(img MapReviewImage) int64 {
		return img.Image.ImageID
	})

	imgMapDbs, err := repository.GetImagesByIDs(
		imgIds,
	)
	if err != nil {
		if err.Code == exception.CODE_RESOURCE_NOT_FOUND {
			return nil
		}
		return routeexception.NewRouteError(
			err,
			"Failed to get images by IDs",
			"get-images-by-ids-failed",
			exception.CODE_INTERNAL_SERVER_ERROR,
		)
	}

	imgMaps := util.MapBy(
		*dbmapper.MapAllDbFields[repository.ImageDB, images.Image](imgMapDbs),
		func(img images.Image) string {
			str := strconv.Itoa(int(img.ImageID))
			return str
		},
		func(img images.Image) *images.Image {
			return &img
		})

	for _, image := range imgs {

		imgId := imgMaps[strconv.Itoa(int(image.Image.ImageID))]
		if imgId == nil || imgId.ImageID == 0 {
			storedlogs.LogWarn("Image ID is zero, cannot populate image for review: " + strconv.Itoa(int(m.MapReviewID)))
			continue
		}

		err := image.Populate(imgId)

		if err != nil {
			if err.Code == exception.CODE_INTERNAL_SERVER_ERROR {
				return routeexception.NewRouteError(err, "Failed to populate image", "populate-image-failed", err.Code)
			}
			storedlogs.LogWarn("Failed to populate image: " + err.Error())
			continue
		}
	}

	return nil

}
