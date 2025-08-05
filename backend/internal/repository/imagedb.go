package repository

import (
	"fmt"

	"github.com/identityofsine/fofx-go-gin-api-template/pkg/db"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/db/dao"
	"github.com/identityofsine/fofx-go-gin-api-template/util"
)

type ImageDB struct {
	ImageID   int64  `db:"image_id" json:"imageId" dao:"pk"`
	ImagePath string `db:"image_path" json:"imagePath"`
	Caption   string `db:"caption" json:"caption"`
}

func GetImages() ([]ImageDB, db.DatabaseError) {
	return dao.SelectFromDatabaseByStruct[ImageDB](
		ImageDB{},
		"")
}

func GetImageByID(imageID int64) (*ImageDB, db.DatabaseError) {
	rows, err := dao.SelectFromDatabaseByStruct(ImageDB{}, "image_id = $1", imageID)
	if err != nil {
		return nil, err
	}

	if len(rows) == 0 {
		return nil, db.NewDatabaseError("GetImageByID", "No image found with the given ID", "no-image-found", 404)
	}

	return &rows[0], nil
}

func GetImagesByIDs(imageID []int64) ([]ImageDB, db.DatabaseError) {
	if len(imageID) == 0 {
		return []ImageDB{}, nil
	}

	query := fmt.Sprintf("image_id IN (%s)", db.Placeholders(len(imageID)))

	ids := util.ToGenericArray(imageID...)

	rows, err := dao.SelectFromDatabaseByStruct(ImageDB{}, query, ids...)
	if err != nil {
		return nil, err
	}

	if len(rows) == 0 {
		return nil, db.NewDatabaseError("GetImagesByIDs", "No images found with the given IDs", "no-images-found", 404)
	}

	return rows, nil
}

func SaveImageDb(review ImageDB) (*ImageDB, db.DatabaseError) {
	ptr, err := dao.InsertOrUpdate(&review,
		func(obj ImageDB) (int64, error) {
			return obj.ImageID, nil
		})
	if err != nil {
		return &review, err
	}
	return ptr, nil
}
