package request

type CreateCoupleRequest struct {
	CoupleType        string `json:"couple_type" form:"couple_type" validate:"required,max=255"`
	Name              string `json:"name" form:"name" validate:"required,max=255"`
	ParentDescription string `json:"parent_description" form:"parent_description" validate:"required,max=255"`
	FatherName        string `json:"father_name" form:"father_name" validate:"required,max=255"`
	MotherName        string `json:"mother_name" form:"mother_name" validate:"required,max=255"`
	Image             string `json:"image" form:"image" validate:"required,max=255"`
	ImageCaption      string `json:"image_caption" form:"image_caption" validate:"required,max=255"`
	InstagramUrl      string `json:"instagram_url" form:"instagram_url" validate:"required,max=255,url"`
}
