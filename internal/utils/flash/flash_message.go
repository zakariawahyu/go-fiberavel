package flash

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/tidwall/gjson"
	"github.com/zakariawahyu/go-fiberavel/app/http/middleware"
)

var keyFlashMessage = "flash_message"

type FlashMessage struct {
	Store   *session.Store
	Message map[string]interface{}
}

type FlashBuilder struct {
	store    *session.Store
	messages map[string]interface{}
}

func NewMessage(store *session.Store) *FlashBuilder {
	return &FlashBuilder{
		store:    store,
		messages: make(map[string]interface{}),
	}
}

func (f *FlashMessage) GetFlash(c *fiber.Ctx) interface{} {
	sess, err := f.Store.Get(c)
	if err != nil {
		return err
	}

	data := sess.Get(keyFlashMessage)
	if data != nil {
		sess.Delete(keyFlashMessage)
		sess.Save()
	}

	str, ok := data.(string)
	if ok {
		return gjson.Parse(str).Value()
	}

	return nil
}

func (f *FlashMessage) SetFlash(c *fiber.Ctx) *fiber.Ctx {
	sess, err := f.Store.Get(c)
	if err != nil {
		log.Error(err)
	}

	data, err := json.Marshal(f.Message)
	if err != nil {
		log.Error(err)
	}

	sess.Set(keyFlashMessage, string(data))

	if err := sess.Save(); err != nil {
		log.Error(err)
	}

	return c
}

func (f *FlashBuilder) WithErrorValidate(err error) *FlashBuilder {
	f.messages["errValidate"] = middleware.GetError(err)
	return f
}

func (f *FlashBuilder) WithError(err error) *FlashBuilder {
	f.messages["error"] = err.Error()
	return f
}

func (f *FlashBuilder) WithSuccess(message string) *FlashBuilder {
	f.messages["success"] = message
	return f
}

func (f *FlashBuilder) WithInput(data interface{}) *FlashBuilder {
	f.messages["old"] = data
	return f
}

func (f *FlashBuilder) Build() FlashMessage {
	return FlashMessage{
		Store:   f.store,
		Message: f.messages,
	}
}
