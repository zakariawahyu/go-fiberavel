package request

import (
	validation "github.com/go-ozzo/ozzo-validation"
	sqlc "github.com/zakariawahyu/go-fiberavel/internal/sqlc/generated"
)

func CreateCoupleValidate(r sqlc.CreateCoupleParams) error {
	return validation.ValidateStruct(
		&r,
		validation.Field(&r.CoupleType, validation.Required),
		validation.Field(&r.Name, validation.Required),
		validation.Field(&r.ParentDescription, validation.Required, validation.Max(255)),
		validation.Field(&r.FatherName, validation.Required, validation.Max(255)),
		validation.Field(&r.MotherName, validation.Required, validation.Max(255)),
		validation.Field(&r.InstagramUrl, validation.Required, validation.Max(255)),
		validation.Field(&r.Image, validation.Required),
		validation.Field(&r.ImageCaption, validation.Required, validation.Max(255)),
	)
}
