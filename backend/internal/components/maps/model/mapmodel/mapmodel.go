package mapmodel

import "github.com/identityofsine/fofx-go-gin-api-template/internal/components/maps/model/maptags"

type Map struct {
	MapName string            `db:"map_name" json:"mapName"`
	MapPath string            `db:"map_path" json:"mapPath"`
	Tags    []maptags.MapTags `dbobj:"MapTagDb" json:"mapTags"` // This is a nested object that will be mapped to a property
}
