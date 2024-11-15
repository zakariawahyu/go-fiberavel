package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/fiber/v2/utils"
)

type Session struct {
	Store *session.Store
}

func InitSessionsStore(storage fiber.Storage) *Session {
	store := session.New(session.Config{
		KeyLookup:      "cookie:session_id",
		KeyGenerator:   utils.UUIDv4,
		Storage:        storage,
		CookieHTTPOnly: true,
		CookieSecure:   true,
	})
	return &Session{
		Store: store,
	}
}
