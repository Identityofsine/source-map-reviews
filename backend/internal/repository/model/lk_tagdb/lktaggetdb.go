package lk_tagdb

import (
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
