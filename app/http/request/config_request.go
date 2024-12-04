package request

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

type ConfiWishRequest struct {
	Title       string `json:"title" form:"title" validate:"required,max=255"`
	Description string `json:"description" form:"description" validate:"required"`
}
