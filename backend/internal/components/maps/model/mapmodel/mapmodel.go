package mapmodel

type Map struct {
	MapName string `db:"map_name" json:"mapName"`
	MapPath string `db:"map_path" json:"mapPath"`
}
