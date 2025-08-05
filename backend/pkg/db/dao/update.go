package dao

import (
	"github.com/identityofsine/fofx-go-gin-api-template/internal/constants/exception"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/db"
)

type KeySupplier[T interface{}] func(obj T) (int64, error)

func InsertOrUpdate[T interface{}](obj *T, keySupplier KeySupplier[T]) (*T, db.DatabaseError) {

	key, err := keySupplier(*obj)
	if err != nil {
		return nil, db.NewDatabaseError("InsertOrUpdate", "Failed to get key for object", "key-supplier-error", exception.CODE_INTERNAL_SERVER_ERROR)
	}

	// If MapReviewId is 0, this is a new review (insert)
	if key == 0 {
		// Insert new review
		inserted, err := InsertIntoDatabaseByStruct(*obj)
		if err != nil {
			return nil, err
		}
		return inserted, nil
	} else {
		// Update existing review
		err := UpdateIntoDatabaseByStruct(*obj)
		if err != nil {
			return nil, err
		}

		// Return the updated review by querying the database
		whereClause, args, err := BuildPrimaryKeyWhereClause(obj)
		if err != nil {
			return nil, err
		}

		results, err := SelectFromDatabaseByStruct(obj, whereClause, args...)
		if err != nil {
			return nil, err
		}

		if len(results) == 0 {
			return nil, db.NewDatabaseError("SaveMapReviewDB", "Review not found after update", "review-not-found", 404)
		}

		return results[0], nil
	}
}
