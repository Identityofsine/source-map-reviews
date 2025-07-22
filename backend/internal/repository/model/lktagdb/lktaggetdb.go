package lktagdb

import (
	"fmt"
	"strings"

	"github.com/identityofsine/fofx-go-gin-api-template/pkg/db"
)

const (
	table = "lk_tags"
)

// GetLkTags retrieves all tag definitions from the lk_tags table
func GetLkTags() (*[]LkTagDb, db.DatabaseError) {
	dbs, err := db.Query[LkTagDb]("SELECT * from " + table)
	if err != nil {
		return nil, err
	}
	return dbs, nil
}

func GetLkTagsByLkTags(lkTags []string) (*[]LkTagDb, db.DatabaseError) {
	if len(lkTags) == 0 {
		return &[]LkTagDb{}, nil
	}

	placeholders := make([]string, len(lkTags))
	args := make([]interface{}, len(lkTags))
	for i, name := range lkTags {
		placeholders[i] = fmt.Sprintf("$%d", i+1) // PostgreSQL style
		args[i] = name
	}

	query := fmt.Sprintf("SELECT * FROM "+table+" WHERE lk_tag IN (%s)", strings.Join(placeholders, ","))

	dbs, err := db.Query[LkTagDb](query, args...)
	if err != nil {
		return nil, err
	}
	return dbs, nil
}

func GetLkTagByLkTag(lkTag string) (*LkTagDb, db.DatabaseError) {
	dbs, err := db.Query[LkTagDb]("SELECT * from "+table+" where lk_tag = ?", lkTag)
	if err != nil {
		return nil, err
	}
	if len(*dbs) == 0 {
		return nil, nil
	}
	return &(*dbs)[0], nil
}
