package maptagdb

// MapTagDb represents a row in the map_tags table (junction table linking maps to tags)
type MapTagDb struct {
	LkTag     string `db:"lk_tag"`
	MapName   string `db:"map_name"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
}
