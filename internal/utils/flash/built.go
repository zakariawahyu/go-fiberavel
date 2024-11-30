package flash

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func HandleSuccess(ctx *fiber.Ctx, store *session.Store, message string, route string) error {
	build := NewMessage(store).WithSuccess(message).Build()

	return build.SetFlash(ctx).Redirect(route)
}

func HandleValidationError(ctx *fiber.Ctx, store *session.Store, err error, data interface{}) error {
	build := NewMessage(store).WithErrorValidate(err).WithInput(data).Build()

	return build.SetFlash(ctx).RedirectBack("/")
}

func HandleError(ctx *fiber.Ctx, store *session.Store, err error, data interface{}) error {
	build := NewMessage(store).WithError(err).WithInput(data).Build()

	return build.SetFlash(ctx).RedirectBack("/mimin/dashboard")
}
