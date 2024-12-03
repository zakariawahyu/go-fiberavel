package request

type ConfigCoverRequest struct {
	Type        string `json:"type" form:"type" validate:"required"`
	Title       string `json:"title" form:"title" validate:"required,max=255"`
	Subtitle    string `json:"subtitle" form:"subtitle" validate:"required,max=255"`
	Description string `json:"description" form:"description" validate:"required"`
}

type ConfigVenueRequest struct {
	Type        string `json:"type" form:"type" validate:"required"`
	Title       string `json:"title" form:"title" validate:"required,max=255"`
	Description string `json:"description" form:"description" validate:"required"`
}

type ConfigGiftRequest struct {
	Type        string `json:"type" form:"type" validate:"required"`
	Title       string `json:"title" form:"title" validate:"required,max=255"`
	Description string `json:"description" form:"description" validate:"required"`
}

type ConfiWishRequest struct {
	Type        string `json:"type" form:"type" validate:"required"`
	Title       string `json:"title" form:"title" validate:"required,max=255"`
	Description string `json:"description" form:"description" validate:"required"`
}
