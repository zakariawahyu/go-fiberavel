package request

import "mime/multipart"

type CreateGalleryRequest struct {
	File         *multipart.FileHeader `json:"image" form:"image" validate:"required,mime=image/jpg-image/jpeg-image/png"`
	Image        string                `json:"_"`
	ImageCaption string                `json:"image_caption" form:"image_caption" validate:"required,max=255"`
}

type UpdateGalleryRequest struct {
	ID           int64                 `json:"id"`
	File         *multipart.FileHeader `json:"image" form:"image" validate:"omitempty,mime=image/jpg-image/jpeg-image/png"`
	Image        string                `json:"_"`
	ImageCaption string                `json:"image_caption" form:"image_caption" validate:"required,max=255"`
}
