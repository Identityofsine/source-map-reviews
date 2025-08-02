package reviews

import "time"

type MapReview struct {
	MapReviewID int64  `json:"mapReviewId" db:"map_review_id" dao:"pk"` // primary key
	MapName     string `json:"mapName" db:"map_name" binding:"required"`
	ReviewerID  int64  `json:"userId" db:"reviewer" binding:"required"`

	Images []MapImage `json:"images"` // array of MapImage

	Stars     int       `json:"stars" db:"stars" binding:"required"`
	Review    string    `json:"review" db:"review" binding:"required"`
	CreatedAt time.Time `json:"createdAt" db:"created_at" dao:"omit"` // time when the review was created
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at" dao:"omit"` // time when the review was last updated
}
