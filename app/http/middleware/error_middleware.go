package middleware

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type ErrorResponse struct {
	Success bool        `json:"success"`
	Code    int         `json:"code"`
	Errors  interface{} `json:"errors"`
}

var (
	ErrNotFound         = errors.New("Your requested item is not found")
	ErrPasswordNotMatch = errors.New("Password not match")
	ErrLogin            = errors.New("Username not found")
)

var ErrorHandler = func(ctx *fiber.Ctx, err error) error {
	code := getStatusCode(err)

	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}

	if err != nil {
		return ctx.Status(code).JSON(ErrorResponse{
			Success: false,
			Code:    code,
			Errors:  ExtractErrorsToMap(err),
		})
	}

	return nil
}

func ExtractErrorsToMap(err error) interface{} {
	errors := make(map[string]interface{})

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, err := range validationErrors {
			switch err.Tag() {
			case "required":
				errors[err.Field()] = fmt.Sprintf("Field %s can not empty!", err.Field())
			case "max":
				errors[err.Field()] = fmt.Sprintf("Field %s must have a maximum of %s characters", err.Field(), err.Param())
			case "url":
				errors[err.Field()] = fmt.Sprintf("Field %s must be a valid URL", err.Field())
			}
		}
		return errors
	}

	return err.Error()
}

func getStatusCode(err error) int {
	if err == nil {
		return fiber.StatusOK
	}

	switch err {
	case ErrNotFound:
		return fiber.StatusNotFound
	case ErrPasswordNotMatch:
		return fiber.StatusBadRequest
	case ErrLogin:
		return fiber.StatusBadRequest
	default:
		return fiber.StatusInternalServerError
	}
}
