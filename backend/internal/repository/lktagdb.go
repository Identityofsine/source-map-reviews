package repository

import (
	"database/sql"

	"github.com/identityofsine/fofx-go-gin-api-template/pkg/db"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/db/dao"
	"github.com/identityofsine/fofx-go-gin-api-template/util"
)

// LkTagDB represents a row in the lk_tags lookup table for tag definitions
// Matches schema in 000000013_create_lk_tags_table.sql
// Columns: lk_tag (PK), description, short_description, created_at, updated_at
// Timestamps are included as per migration options

type LkTagDB struct {
	LkTag            string         `db:"lk_tag" json:"tagLk"`
	Description      sql.NullString `db:"description" json:"description"`
	ShortDescription sql.NullString `db:"short_description" json:"shortDescription"`
	CreatedAt        string         `db:"created_at" json:"createdAt" dao:"omit"`
	UpdatedAt        string         `db:"updated_at" json:"updatedAt" dao:"omit"`
}

const (
	lktags_table = "lk_tags"
)

// GetLkTags retrieves all tag definitions from the lk_tags table
func GetLkTags() (*[]LkTagDB, db.DatabaseError) {
	dbs, err := dao.SelectFromDatabaseByStruct(LkTagDB{}, "")
	if err != nil {
		return nil, err
	}
	return &dbs, nil
}

func GetLkTagsByLkTags(lkTags []string) (*[]LkTagDB, db.DatabaseError) {
	if len(lkTags) == 0 {
		return &[]LkTagDB{}, nil
	}

	whereClause := "lk_tag IN (" + db.Placeholders(len(lkTags)) + ")"
	args := util.ToGenericArray(lkTags...)

	dbs, err := dao.SelectFromDatabaseByStruct(LkTagDB{}, whereClause, args...)
	if err != nil {
		return nil, err
	}
	return &dbs, nil
}

func GetLkTagByLkTag(lkTag string) (*LkTagDB, db.DatabaseError) {
	dbs, err := dao.SelectFromDatabaseByStruct(LkTagDB{}, "lk_tag = $1", lkTag)
	if err != nil {
		return nil, err
	}
	if len(dbs) == 0 {
		return nil, nil
	}
	return &dbs[0], nil
}

// SaveLkTagDb handles both insert and update for tag lookups
func SaveLkTagDb(tagLk LkTagDB) (*LkTagDB, db.DatabaseError) {
	ptr, err := dao.InsertOrUpdate(&tagLk,
		func(obj LkTagDB) (int64, error) {
			// For LkTag, we use a different strategy since lk_tag is string
			// We'll check if it exists first
			existing, dbErr := GetLkTagByLkTag(obj.LkTag)
			if dbErr != nil {
				return 0, dbErr
			}
			if existing != nil {
				return 1, nil // Indicates update
			}
			return 0, nil // Indicates insert
		})
	if err != nil {
		return &tagLk, err
	}
	return ptr, nil
}

// DeleteLkTagDb deletes a tag lookup by lk_tag
func DeleteLkTagDb(lkTag string) db.DatabaseError {
	_, err := db.Delete("DELETE FROM "+lktags_table+" WHERE lk_tag = $1", lkTag)
	return err
}
