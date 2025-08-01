package images

type Image struct {
	ImageID   string `json:"imageId" binding:"required" db:"image_id"`
	ImagePath string `json:"imagePath" binding:"required" db:"image_path"`
	Caption   string `json:"caption" binding:"required" db:"caption"`
}
