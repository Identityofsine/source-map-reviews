package lk_tagdb

// LkTagDb represents a row in the lk_tags lookup table for tag definitions
// Matches schema in 000000013_create_lk_tags_table.sql
// Columns: lk_tag (PK), description, short_description, created_at, updated_at
// Timestamps are included as per migration options

type LkTagDb struct {
	LkTag            string `db:"lk_tag"`
	Description      string `db:"description"`
	ShortDescription string `db:"short_description"`
	CreatedAt        string `db:"created_at"`
	UpdatedAt        string `db:"updated_at"`
}
