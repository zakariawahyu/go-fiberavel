package middleware

import (
	"encoding/json"
	"errors"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/gofiber/fiber/v2"
	"github.com/tidwall/gjson"
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
			Errors:  GetError(err),
		})
	}

	return nil
}

func GetError(err error) any {
	if err, ok := err.(validation.Errors); ok {
		res, _ := json.Marshal(err)
		return gjson.Parse(string(res)).Value()
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
