package reviews

import "github.com/identityofsine/fofx-go-gin-api-template/internal/components/images"

type MapImage struct {
	MapImageID string       `json:"mapImageId" binding:"required" db:"map_image_id"` // pk
	MapName    string       `json:"mapName" binding:"required" db:"map_name"`        // fk to MapName
	Image      images.Image `json:"image" binding:"required" dbobj:"images.Image"`
}
