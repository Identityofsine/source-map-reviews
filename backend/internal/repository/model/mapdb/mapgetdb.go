package mapdb

import (
	"fmt"

	"github.com/identityofsine/fofx-go-gin-api-template/internal/components/maps/model/mapsearchform"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/db"
)

const (
	table = "maps"
)

func GetMaps() (*[]MapDb, db.DatabaseError) {
	dbs, err := db.Query[MapDb]("SELECT * from " + table)
	if err != nil {
		return nil, err
	}

	return dbs, nil
}

func SearchMaps(form mapsearchform.MapSearchForm) (*[]MapDb, db.DatabaseError) {

	query := fmt.Sprintf(`
		SELECT * FROM %s
		WHERE 1=1
	`, table)

	args := []interface{}{}
	if form.SearchTerm != "" {
		query += " AND (map_name LIKE ? OR description LIKE ?)"
		args = append(args, "%"+form.SearchTerm+"%", "%"+form.SearchTerm+"%")
	}
	if form.Reviewed {
		query += " AND reviewed = ?"
		args = append(args, true)
	}
	if form.Unreviewed {
		query += " AND reviewed = ?"
		args = append(args, false)
	}
	if len(form.Tags) > 0 {
		query += " AND map_name IN (SELECT map_name FROM map_tags WHERE lk_tag IN ("
		for i := range form.Tags {
			args = append(args, form.Tags[i])
		}
		query += fmt.Sprintf("%s))", db.Placeholders(len(form.Tags)))
	}
	dbs, err := db.Query[MapDb](query, args...)
	if err != nil {
		return nil, db.NewDatabaseError(
			"SearchMaps",
			"Failed to search maps",
			"search-maps-failed",
			err.Code,
		)
	}
	return dbs, nil

}
