package request

import (
	"mime/multipart"
)

type ConfigRequest struct {
	Type         string `json:"type"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	Image        string `json:"image"`
	ImageCaption string `json:"image_caption"`
	CustomData   []byte `json:"custom_data"`
}

type ConfigCoverRequest struct {
	Title       string `json:"title" form:"title" validate:"required,max=255"`
	Subtitle    string `json:"subtitle" form:"subtitle" validate:"required,max=255"`
	Description string `json:"description" form:"description" validate:"required"`
}

type ConfigVenueRequest struct {
	Title       string `json:"title" form:"title" validate:"required,max=255"`
	Description string `json:"description" form:"description" validate:"required"`
}

type ConfigGiftRequest struct {
	Title       string `json:"title" form:"title" validate:"required,max=255"`
	Description string `json:"description" form:"description" validate:"required"`
}

type ConfigWishRequest struct {
	Title       string `json:"title" form:"title" validate:"required,max=255"`
	Description string `json:"description" form:"description" validate:"required"`
}

type ConfigEventRequest struct {
	Title        string                `json:"title" form:"title" validate:"required,max=255"`
	Description  string                `json:"description" form:"description" validate:"required"`
	File         *multipart.FileHeader `json:"image" form:"image" validate:"omitempty,required,mime=image/jpg-image/jpeg-image/png"`
	Image        string                `json:"_"`
	ImageCaption string                `json:"image_caption" form:"image_caption" validate:"required,max=255"`
	Date         string                `json:"date" form:"date" validate:"required"`
}

type ConfigRsvpRequest struct {
	Title        string                `json:"title" form:"title" validate:"required,max=255"`
	Description  string                `json:"description" form:"description" validate:"required"`
	File         *multipart.FileHeader `json:"image" form:"image" validate:"omitempty,required,mime=image/jpg-image/jpeg-image/png"`
	Image        string                `json:"_"`
	ImageCaption string                `json:"image_caption" form:"image_caption" validate:"required,max=255"`
}

type ConfigStoryRequest struct {
	Title        string                `json:"title" form:"title" validate:"required,max=255"`
	Description  string                `json:"description" form:"description" validate:"required"`
	File         *multipart.FileHeader `json:"image" form:"image" validate:"omitempty,required,mime=image/jpg-image/jpeg-image/png"`
	Image        string                `json:"_"`
	ImageCaption string                `json:"image_caption" form:"image_caption" validate:"required,max=255"`
}

type ConfigThankRequest struct {
	Title        string                `json:"title" form:"title" validate:"required,max=255"`
	Description  string                `json:"description" form:"description" validate:"required"`
	File         *multipart.FileHeader `json:"image" form:"image" validate:"omitempty,required,mime=image/jpg-image/jpeg-image/png"`
	Image        string                `json:"_"`
	ImageCaption string                `json:"image_caption" form:"image_caption" validate:"required,max=255"`
}

type ConfigMetaRequest struct {
	Title       string                `json:"title" form:"title" validate:"required,min=30,max=65"`
	Description string                `json:"description" form:"description" validate:"required,min=120,max=320"`
	Keywords    string                `json:"keywords" form:"keywords" validate:"required,max=255"`
	Author      string                `json:"author" form:"author" validate:"required,max=255"`
	FileImage   *multipart.FileHeader `json:"image" form:"image" validate:"omitempty,required,mime=image/jpg-image/jpeg-image/png"`
	FileIcon    *multipart.FileHeader `json:"icon" form:"icon" validate:"omitempty,required,mime=image/jpg-image/jpeg-image/png"`
	Image       string                `json:"_"`
	Icon        string                `json:"_"`
}
