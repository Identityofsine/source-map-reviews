package images

type Image struct {
	ImageID   int64  `json:"imageId" binding:"required" db:"image_id"`
	ImagePath string `json:"imagePath" db:"image_path"`
	Caption   string `json:"caption" db:"caption"`
}
