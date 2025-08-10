package repository

import (
	"fmt"
	"strings"

	"github.com/identityofsine/fofx-go-gin-api-template/internal/components/maps/mapsearchform"
	"github.com/identityofsine/fofx-go-gin-api-template/internal/constants/exception"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/db"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/db/dao"
)

type MapDB struct {
	MapName   string `db:"map_name"`
	MapPath   string `db:"map_path"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
}

const (
	mapdb_table = "maps"
)

func GetMaps() (*[]MapDB, db.DatabaseError) {
	dbs, err := dao.SelectFromDatabaseByStruct(MapDB{}, "")
	if err != nil {
		return nil, err
	}

	return &dbs, nil
}

func GetMap(mapName string) (*MapDB, db.DatabaseError) {
	// Use exact case-sensitive match for map name lookup
	dbs, err := dao.SelectFromDatabaseByStruct(MapDB{}, "map_name = $1", mapName)
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

func SearchMaps(form mapsearchform.MapSearchForm) (*[]MapDB, db.DatabaseError) {
	var conditions []string
	args := []interface{}{}
	argIndex := 1

	// Case-insensitive search term with multiple search patterns using ILIKE
	if form.SearchTerm != "" {
		searchConditions := []string{
			fmt.Sprintf("map_name ILIKE $%d", argIndex),   // Case-insensitive substring match
			fmt.Sprintf("map_name ILIKE $%d", argIndex+1), // Case-insensitive start of string match
			fmt.Sprintf("map_name ILIKE $%d", argIndex+2), // Case-insensitive end of string match
		}

		// Add the search term variations
		args = append(args, "%"+form.SearchTerm+"%") // Contains
		args = append(args, form.SearchTerm+"%")     // Starts with
		args = append(args, "%"+form.SearchTerm)     // Ends with
		argIndex += 3

		// Combine search conditions with OR for more flexible matching
		conditions = append(conditions, "("+strings.Join(searchConditions, " OR ")+")")
	}

	// Handle review status with proper subquery
	if form.Reviewed && !form.Unreviewed {
		// Only reviewed maps - maps that have at least one review
		conditions = append(conditions, `
			map_name IN (
				SELECT DISTINCT map_name 
				FROM map_reviews
			)`)
	} else if form.Unreviewed && !form.Reviewed {
		// Only unreviewed maps - maps with no reviews
		conditions = append(conditions, `
			map_name NOT IN (
				SELECT DISTINCT map_name 
				FROM map_reviews
			)`)
	}
	// If both or neither are true, don't filter by review status

	// Handle tags with optimized subquery
	if len(form.Tags) > 0 {
		tagPlaceholders := make([]string, len(form.Tags))
		for i, tag := range form.Tags {
			tagPlaceholders[i] = fmt.Sprintf("$%d", argIndex)
			args = append(args, tag)
			argIndex++
		}

		// Use subquery to find maps that have ALL specified tags
		tagCondition := fmt.Sprintf(`
			map_name IN (
				SELECT map_name 
				FROM map_tags 
				WHERE lk_tag IN (%s)
				GROUP BY map_name 
				HAVING COUNT(DISTINCT lk_tag) = $%d
			)`, strings.Join(tagPlaceholders, ", "), argIndex)

		args = append(args, len(form.Tags))
		argIndex++
		conditions = append(conditions, tagCondition)
	}

	// Build WHERE clause
	whereClause := "1=1"
	if len(conditions) > 0 {
		whereClause += " AND " + strings.Join(conditions, " AND ")
	}

	// Add ordering for consistent results and better user experience
	whereClause += " ORDER BY map_name ASC"

	dbs, err := dao.SelectFromDatabaseByStruct(MapDB{}, whereClause, args...)

	return &dbs, err
}
