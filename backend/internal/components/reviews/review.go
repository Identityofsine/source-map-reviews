package reviews

import "time"

type MapReview struct {
	MapReviewID string `json:"mapReviewId" db:"map_review_id"` // pk
	MapName     string `json:"mapName" db:"map_name" binding:"required"`
	ReviewerID  string `json:"userId" db:"reviewer" binding:"required"`

	Images []MapImage `json:"images" binding:"required"` // array of MapImage

	Stars     int       `json:"stars" db:"stars" binding:"required"`
	Review    string    `json:"review" db:"review" binding:"required"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"` // time when the review was created
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at"` // time when the review was last updated
}
