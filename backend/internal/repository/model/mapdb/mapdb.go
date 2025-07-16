package mapdb

type MapDb struct {
	MapName   string `db:"map_name"`
	MapPath   string `db:"map_path"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
}
