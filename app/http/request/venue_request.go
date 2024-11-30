package request

type CreateVenueRequest struct {
	Name     string `json:"name" form:"name" validate:"required,max=255"`
	Location string `json:"location" form:"location" validate:"required,max=255"`
	Address  string `json:"address" form:"address" validate:"required,max=255"`
	DateHeld string `json:"date_held" form:"date_held" validate:"required"`
	Map      string `json:"map" form:"map" validate:"required"`
}

type UpdateVenueRequest struct {
	ID       int64  `json:"id" validate:"required"`
	Name     string `json:"name" form:"name" validate:"required,max=255"`
	Location string `json:"location" form:"location" validate:"required,max=255"`
	Address  string `json:"address" form:"address" validate:"required,max=255"`
	DateHeld string `json:"date_held" form:"date_held" validate:"required"`
	Map      string `json:"map" form:"map" validate:"required"`
}
