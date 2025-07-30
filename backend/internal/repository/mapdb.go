package repository

import (
	"fmt"
	"strings"

	"github.com/identityofsine/fofx-go-gin-api-template/internal/components/maps/mapsearchform"
	"github.com/identityofsine/fofx-go-gin-api-template/internal/constants/exception"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/db"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/db/dao"
)

type MapDb struct {
	MapName   string `db:"map_name"`
	MapPath   string `db:"map_path"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
}

const (
	mapdb_table = "maps"
)

func GetMaps() (*[]MapDb, db.DatabaseError) {
	dbs, err := dao.SelectFromDatabaseByStruct(MapDb{}, "")
	if err != nil {
		return nil, err
	}

	return &dbs, nil
}

func GetMap(mapName string) (*MapDb, db.DatabaseError) {
	dbs, err := dao.SelectFromDatabaseByStruct(MapDb{}, "map_name = $1", mapName)
	if err != nil {
		return nil, err
	}

	if dbs == nil || len(dbs) == 0 {
		return nil, db.NewDatabaseError(
			"GetMap",
			fmt.Sprintf("Map with name '%s' not found", mapName),
			"map-not-found",
			exception.CODE_RESOURCE_NOT_FOUND,
		)
	}

	return &(dbs)[0], nil
}

func SearchMaps(form mapsearchform.MapSearchForm) (*[]MapDb, db.DatabaseError) {

	query := ""

	args := []interface{}{}
	argIndex := 1 // PostgreSQL placeholders start at $1

	if form.SearchTerm != "" {
		query += fmt.Sprintf(" AND map_name LIKE $%d", argIndex)
		args = append(args, "%"+form.SearchTerm+"%")
		argIndex++
	}
	if form.Reviewed {
		query += fmt.Sprintf(" AND reviewed = $%d", argIndex)
		args = append(args, true)
		argIndex++
	}
	if form.Unreviewed {
		query += fmt.Sprintf(" AND reviewed = $%d", argIndex)
		args = append(args, false)
		argIndex++
	}
	if len(form.Tags) > 0 {
		query += " AND map_name IN (SELECT map_name FROM map_tags WHERE lk_tag IN ("
		placeholders := []string{}
		for _, tag := range form.Tags {
			args = append(args, tag)
			placeholders = append(placeholders, fmt.Sprintf("$%d", argIndex))
			argIndex++
		}
		query += strings.Join(placeholders, ", ") + "))"
	}

	dbs, err := dao.SelectFromDatabaseByStruct(MapDb{}, query, args...)

	return &dbs, err
}
