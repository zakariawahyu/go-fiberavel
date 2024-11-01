package utils

import (
	"errors"
	"github.com/gofiber/fiber/v2"
)

type ErrorResponse struct {
	Success bool        `json:"success"`
	Code    int         `json:"code"`
	Errors  interface{} `json:"errors"`
}

var (
	ErrNotFound = errors.New("Your requested item is not found")
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
			Errors:  err.Error(),
		})
	}

	return nil
}

func getStatusCode(err error) int {
	if err == nil {
		return fiber.StatusOK
	}

	switch err {
	case ErrNotFound:
		return fiber.StatusNotFound
	default:
		return fiber.StatusInternalServerError
	}
}
