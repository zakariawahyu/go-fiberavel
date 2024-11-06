package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/fiber/v2/utils"
)

var Store *session.Store

func InitSessionsStore(storage fiber.Storage) {
	Store = session.New(session.Config{
		KeyLookup:      "cookie:session_id",
		KeyGenerator:   utils.UUIDv4,
		Storage:        storage,
		CookieHTTPOnly: true,
		CookieSecure:   true,
	})
}
