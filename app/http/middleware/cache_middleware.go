package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"strings"
	"time"
)

func CacheMiddleware(storage fiber.Storage) fiber.Handler {
	return cache.New(cache.Config{
		Expiration:   15 * time.Minute,
		CacheControl: true,
		Storage:      storage,
		Next: func(c *fiber.Ctx) bool {
			return strings.HasPrefix(c.Path(), "/mimin")
		},
	})
}
