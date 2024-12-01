package request

type CreateWishParams struct {
	Name            string `json:"name" form:"name" validate:"required"`
	WishDescription string `json:"wish_description" validate:"required"`
}
