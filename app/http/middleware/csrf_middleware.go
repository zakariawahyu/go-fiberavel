package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/zakariawahyu/go-fiberavel/internal/utils/constants"
)

func CSRFMiddleware(store *session.Store) fiber.Handler {
	return csrf.New(csrf.Config{
		KeyLookup:         "header:X-Csrf-Token",
		CookieName:        "csrf_token",
		CookieSameSite:    "Lax",
		CookieHTTPOnly:    true,
		CookieSecure:      true,
		CookieSessionOnly: true,
		KeyGenerator:      utils.UUIDv4,
		Session:           store,
		ContextKey:        constants.CSRFContextKey,
		SessionKey:        "fiber.csrf.token",
		HandlerContextKey: "fiber.csrf.handler",
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return err
		},
	})
}
