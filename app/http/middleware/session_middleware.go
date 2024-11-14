package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/fiber/v2/utils"
	"sync"
)

type FlashStore struct {
	messages map[string]fiber.Map
	mu       sync.Mutex
	Store    *session.Store
}

func InitSessionsStore(storage fiber.Storage) *FlashStore {
	store := session.New(session.Config{
		KeyLookup:      "cookie:session_id",
		KeyGenerator:   utils.UUIDv4,
		Storage:        storage,
		CookieHTTPOnly: true,
		CookieSecure:   true,
	})
	return &FlashStore{
		messages: make(map[string]fiber.Map),
		Store:    store,
	}
}

func (fs *FlashStore) SetFlash(c *fiber.Ctx, message fiber.Map) *fiber.Ctx {
	sess, err := fs.Store.Get(c)
	if err != nil {
		log.Error(err)
	}
	sessionID := sess.ID()

	fs.mu.Lock()
	fs.messages[sessionID] = message
	fs.mu.Unlock()

	if err := sess.Save(); err != nil {
		log.Error(err)
	}

	return c
}

func (fs *FlashStore) GetFlash(c *fiber.Ctx) fiber.Map {
	sess, err := fs.Store.Get(c)
	if err != nil {
		return nil
	}
	sessionID := sess.ID()

	fs.mu.Lock()
	message, exists := fs.messages[sessionID]
	if exists {
		delete(fs.messages, sessionID)
	}
	fs.mu.Unlock()

	return message
}
