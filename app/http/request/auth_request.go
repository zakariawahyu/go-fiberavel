package request

import (
	validation "github.com/go-ozzo/ozzo-validation"
	sqlc "github.com/zakariawahyu/go-fiberavel/internal/sqlc/generated"
)

func LoginValidate(r sqlc.LoginRow) error {
	return validation.ValidateStruct(
		&r,
		validation.Field(&r.Username, validation.Required),
		validation.Field(&r.Password, validation.Required),
	)
}
