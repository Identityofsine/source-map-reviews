package images

import "mime/multipart"

type SaveImageRequest struct {
	Caption string `json:"caption" binding:"required"`
	FileExt string `json:"fileExt" binding:"required"`
}

type SaveImageForm struct {
	JSON *string                 `form:"json" binding:"required"`
	File []*multipart.FileHeader `form:"files" binding:"required"`
}
