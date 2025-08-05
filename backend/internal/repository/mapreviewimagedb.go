package repository

import (
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/db"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/db/dao"
)

type MapReviewImageDB struct {
	MapReviewImageId int64 `db:"map_review_image_id" json:"mapReviewImage" dao:"pk"`
	MapReviewId      int64 `db:"map_review_id" json:"mapReviewId"`
	ImageId          int64 `db:"image_id" json:"imageId"`
}

func SaveMapReviewImageDb(reviewImage MapReviewImageDB) (*MapReviewImageDB, db.DatabaseError) {
	ptr, err := dao.InsertOrUpdate(&reviewImage,
		func(obj MapReviewImageDB) (int64, error) {
			return obj.MapReviewImageId, nil
		})
	if err != nil {
		return &reviewImage, err
	}
	return ptr, nil
}

func GetMapReviewImagesByReviewId(reviewId int64) ([]MapReviewImageDB, db.DatabaseError) {
	rows, err := dao.SelectFromDatabaseByStruct(
		MapReviewImageDB{},
		"map_review_id = $1",
		reviewId,
	)
	if err != nil {
		return nil, err
	}

	if len(rows) == 0 {
		return nil, db.NewDatabaseError("GetMapReviewImagesByReviewId", "No images found for the given review ID", "no-images-found", 404)
	}

	return rows, nil
}
