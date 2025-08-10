package maps

import "github.com/identityofsine/fofx-go-gin-api-template/internal/components/images"

type Map struct {
	MapName   string       `db:"map_name" json:"mapName"`
	MapPath   string       `db:"map_path" json:"mapPath"`
	Thumbnail images.Image `db:"thumbnail" json:"thumbnail"`            // This is a nested object that will be mapped to a property
	Tags      []MapTag     `dbobj:"[]maptagdb.MapTagDB" json:"mapTags"` // This is a nested object that will be mapped to a property
}
