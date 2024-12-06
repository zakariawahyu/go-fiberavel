package request

type CreateGuestRequest struct {
	Name   string `json:"name" form:"name" validate:"required,max=255"`
	Slug   string `json:"slug" `
	IsGift bool   `json:"is_gift" form:"is_gift"`
}

type UpdateGuestRequest struct {
	ID     int64  `json:"id" validate:"required"`
	Name   string `json:"name" form:"name" validate:"required,max=255"`
	Slug   string `json:"slug" `
	IsGift bool   `json:"is_gift" form:"is_gift"`
}
