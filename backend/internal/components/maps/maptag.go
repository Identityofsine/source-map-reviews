package maps

type TagLk struct {
	TagLk               string `db:"lk_tag" json:"tagLk"`
	TagDescription      string `db:"description" json:"tagDescription"`
	TagDescriptionShort string `db:"short_description" json:"tagDescriptionShort"`
	CreatedAt           string `db:"created_at" json:"createdAt"`
	UpdatedAt           string `db:"updated_at" json:"updatedAt"`
}

type MapTag struct {
	TagName             string `db:"lk_tag" json:"tagName"`
	TagDescription      string `db:"tag_description" json:"tagDescription"`
	TagDescriptionShort string `db:"tag_description_short" json:"tagDescriptionShort"`
	CreatedAt           string `db:"created_at" json:"createdAt"`
	UpdatedAt           string `db:"updated_at" json:"updatedAt"`
}
