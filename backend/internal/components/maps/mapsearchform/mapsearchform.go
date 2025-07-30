package mapsearchform

type MapSearchForm struct {
	SearchTerm string   `json:"searchTerm" binding:"omitempty"`
	Tags       []string `json:"tags" binding:"omitempty"`
	Reviewed   bool     `json:"reviewed" binding:"omitempty" default:"true"`
	Unreviewed bool     `json:"unreviewed" binding:"omitempty" default:"false"`
}
