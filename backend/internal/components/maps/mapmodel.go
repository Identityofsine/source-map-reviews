package maps

type Map struct {
	MapName string   `db:"map_name" json:"mapName"`
	MapPath string   `db:"map_path" json:"mapPath"`
	Tags    []MapTag `dbobj:"[]maptagdb.MapTagDb" json:"mapTags"` // This is a nested object that will be mapped to a property
}
