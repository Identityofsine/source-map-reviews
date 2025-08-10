package repository

import (
	"database/sql"
	"time"

	"github.com/identityofsine/fofx-go-gin-api-template/pkg/db"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/db/dao"
)

type MapReviewDB struct {
	MapReviewId int64          `db:"map_review_id" json:"mapReviewId" dao:"pk"`
	MapName     string         `db:"map_name" json:"mapName" binding:"required"`
	ReviewerId  int64          `db:"reviewer" json:"userId" binding:"required"` // User ID of the reviewer
	Stars       int            `db:"stars" json:"stars" binding:"required"`     // Rating given by the reviewer
	Review      sql.NullString `db:"review" json:"review" binding:"required"`   // Text of the review
	UpdatedAt   time.Time      `db:"updated_at" json:"updatedAt" dao:"omit"`    // Time when the review was last updated
	CreatedAt   time.Time      `db:"created_at" json:"createdAt" dao:"omit"`    // Time when the review was created
}

func selectMapReviewDBWrapper(whereClause string, args ...interface{}) ([]MapReviewDB, db.DatabaseError) {

	dbs, err := dao.SelectFromDatabaseByStruct(MapReviewDB{}, whereClause, args...)
	if err != nil {
		return nil, err
	}

	return dbs, nil
}

func GetMapReviewDBByMapName(mapName string) ([]MapReviewDB, db.DatabaseError) {
	return selectMapReviewDBWrapper("map_name = $1", mapName)
}

func GetMapReviewDBById(reviewId int64) (*MapReviewDB, db.DatabaseError) {
	dbs, err := selectMapReviewDBWrapper("map_review_id = $1", reviewId)
	if err != nil {
		return nil, err
	}

	if len(dbs) == 0 {
		return nil, db.NewDatabaseError("GetMapReviewDBById", "Review not found", "review-not-found", 404)
	}

	return &dbs[0], nil
}

func GetMapReviewDBByMapNameAndReviewer(mapName string, reviewerId int64) (*MapReviewDB, db.DatabaseError) {
	dbs, err := selectMapReviewDBWrapper("map_name = $1 AND reviewer = $2", mapName, reviewerId)
	if err != nil {
		return nil, err
	}

	if len(dbs) == 0 {
		return nil, nil // No review found, but not an error
	}

	return &dbs[0], nil
}

func GetMapReviewDBByReviewer(reviewerId int64) ([]MapReviewDB, db.DatabaseError) {
	dbs, err := selectMapReviewDBWrapper("reviewer = $1", reviewerId)
	if err != nil {
		return nil, err
	}

	return dbs, nil
}

// GetMapReviewDBByMapNames retrieves reviews for multiple map names efficiently
func GetMapReviewDBByMapNames(mapNames []string) (map[string][]MapReviewDB, db.DatabaseError) {
	if mapNames == nil || len(mapNames) == 0 {
		return make(map[string][]MapReviewDB), nil
	}

	whereClause := "map_name IN (" + db.Placeholders(len(mapNames)) + ")"

	// Convert to interface{} slice for variadic args
	args := make([]interface{}, len(mapNames))
	for i, mapName := range mapNames {
		args[i] = mapName
	}

	dbs, err := selectMapReviewDBWrapper(whereClause, args...)
	if err != nil {
		return nil, err
	}

	// Group reviews by map name
	reviewsByMap := make(map[string][]MapReviewDB)
	for _, review := range dbs {
		reviewsByMap[review.MapName] = append(reviewsByMap[review.MapName], review)
	}

	// Initialize empty slices for maps with no reviews
	for _, mapName := range mapNames {
		if _, exists := reviewsByMap[mapName]; !exists {
			reviewsByMap[mapName] = []MapReviewDB{}
		}
	}

	return reviewsByMap, nil
}

func DeleteMapReviewDB(reviewId int64) db.DatabaseError {
	_, err := selectMapReviewDBWrapper("map_review_id = $1", reviewId)
	if err != nil {
		return err
	}

	// Use the db.Query directly for DELETE since DAO doesn't have a delete method yet
	_, err = db.Query[interface{}]("DELETE FROM map_reviews WHERE map_review_id = $1", reviewId)
	return err
}

func SaveMapReviewDB(review MapReviewDB) (MapReviewDB, db.DatabaseError) {
	// If MapReviewId is 0, this is a new review (insert)
	if review.MapReviewId == 0 {
		// Insert new review
		insertedReview, err := dao.InsertIntoDatabaseByStruct(review)
		if err != nil {
			return MapReviewDB{}, err
		}
		return *insertedReview, nil
	} else {
		// Update existing review
		err := dao.UpdateIntoDatabaseByStruct(review)
		if err != nil {
			return MapReviewDB{}, err
		}

		// Return the updated review by querying the database
		whereClause, args, err := dao.BuildPrimaryKeyWhereClause(review)
		if err != nil {
			return MapReviewDB{}, err
		}

		results, err := selectMapReviewDBWrapper(whereClause, args...)
		if err != nil {
			return MapReviewDB{}, err
		}

		if len(results) == 0 {
			return MapReviewDB{}, db.NewDatabaseError("SaveMapReviewDB", "Review not found after update", "review-not-found", 404)
		}

		return results[0], nil
	}
}
