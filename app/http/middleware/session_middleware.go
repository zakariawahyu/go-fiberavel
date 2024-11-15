package middleware

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/fiber/v2/utils"
)

const KeyAuthSession = "auth_session"

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

func (s *Session) SetAuth(ctx *fiber.Ctx, data interface{}) error {
	sess, err := s.Store.Get(ctx)
	if err != nil {
		return err
	}

	encode, err := json.Marshal(data)
	if err != nil {
		return err
	}

	sess.Set(KeyAuthSession, encode)

	return sess.Save()
}

func (s *Session) Authenticated() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		sess, err := s.Store.Get(ctx)
		if err != nil {
			return err
		}
		data := sess.Get(KeyAuthSession)
		if data == nil {
			return ctx.Redirect("/auth/unauthorized")
		}
		return ctx.Next()
	}
}
