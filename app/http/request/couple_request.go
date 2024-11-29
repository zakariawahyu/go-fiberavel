package request

import "mime/multipart"

type CreateCoupleRequest struct {
	CoupleType        string                `json:"couple_type" form:"couple_type" validate:"required"`
	Name              string                `json:"name" form:"name" validate:"required,max=255"`
	ParentDescription string                `json:"parent_description" form:"parent_description" validate:"required,max=255"`
	FatherName        string                `json:"father_name" form:"father_name" validate:"required,max=255"`
	MotherName        string                `json:"mother_name" form:"mother_name" validate:"required,max=255"`
	File              *multipart.FileHeader `json:"image" form:"image" validate:"required,mime=image/jpg-image/jpeg-image/png"`
	Image             string                `json:"_"`
	ImageCaption      string                `json:"image_caption" form:"image_caption" validate:"required,max=255"`
	InstagramUrl      string                `json:"instagram_url" form:"instagram_url" validate:"required,max=255,url"`
}

type UpdateCoupleRequest struct {
	ID                int64                 `json:"id"`
	CoupleType        string                `json:"couple_type" form:"couple_type" validate:"required"`
	Name              string                `json:"name" form:"name" validate:"required,max=255"`
	ParentDescription string                `json:"parent_description" form:"parent_description" validate:"required,max=255"`
	FatherName        string                `json:"father_name" form:"father_name" validate:"required,max=255"`
	MotherName        string                `json:"mother_name" form:"mother_name" validate:"required,max=255"`
	File              *multipart.FileHeader `json:"image" form:"image" validate:"omitempty,mime=image/jpg-image/jpeg-image/png"`
	Image             string                `json:"_"`
	ImageCaption      string                `json:"image_caption" form:"image_caption" validate:"required,max=255"`
	InstagramUrl      string                `json:"instagram_url" form:"instagram_url" validate:"required,max=255,url"`
}
