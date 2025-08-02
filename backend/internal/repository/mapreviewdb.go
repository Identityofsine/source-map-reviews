package repository

import (
	"database/sql"
	"time"

	"github.com/identityofsine/fofx-go-gin-api-template/pkg/db"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/db/dao"
)

type MapReviewDB struct {
	MapReviewId int64          `db:"map_review_id" json:"mapReviewId" dao:"omit"`
	MapName     string         `db:"map_name" json:"mapName" binding:"required"`
	ReviewerId  int64          `db:"reviewer" json:"userId" binding:"required"` // User ID of the reviewer
	Stars       int            `db:"stars" json:"stars" binding:"required"`     // Rating given by the reviewer
	Review      sql.NullString `db:"review" json:"review" binding:"required"`   // Text of the review
	UpdatedAt   time.Time      `db:"updated_at" json:"updatedAt" dao:"omit"`    // Time when the review was last updated
	CreatedAt   time.Time      `db:"created_at" json:"createdAt" dao:"omit"`    // Time when the review was created
}

func GetMapReviewDBByMapName(mapName string) ([]MapReviewDB, db.DatabaseError) {
	dbs, err := dao.SelectFromDatabaseByStruct(MapReviewDB{}, "map_name = $1", mapName)
	if err != nil {
		return nil, err
	}

	return dbs, nil
}

func SaveMapReviewDB(review MapReviewDB) (MapReviewDB, db.DatabaseError) {
	// Insert or update the review in the database
	err := dao.InsertIntoDatabaseByStruct(review)
	if err != nil {
		return MapReviewDB{}, err
	}

	return review, nil
}
