package controller

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/tidwall/gjson"
	"github.com/zakariawahyu/go-fiberavel/config"
	"github.com/zakariawahyu/go-fiberavel/internal/infrastructure/cache"
	"github.com/zakariawahyu/go-fiberavel/internal/utils/constants"
	"time"
)

type HomeController struct {
	redis  cache.Rueidis
	cfgApp config.App
}

func NewHomeController(redis cache.Rueidis, cfgApp config.App) *HomeController {
	return &HomeController{
		redis:  redis,
		cfgApp: cfgApp,
	}
}

func (ctrl *HomeController) Index(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), ctrl.cfgApp.Timeout*time.Second)
	defer cancel()

	configs, err := ctrl.redis.HGetAll(c, constants.KeyConfigs)
	if err != nil {
		return err
	}

	resCouples, err := ctrl.redis.Get(constants.KeyCouples)
	if err != nil {
		return err
	}

	resVenueDetails, err := ctrl.redis.Get(constants.KeyVenueDetails)
	if err != nil {
		return err
	}

	resGalleries, err := ctrl.redis.Get(constants.KeyGalleries)
	if err != nil {
		return err
	}

	resGifts, err := ctrl.redis.Get(constants.KeyGift)
	if err != nil {
		return err
	}

	resGuest, err := ctrl.redis.HGet(constants.KeyGuestList, "akbar-gustama")
	if err != nil {
		return err
	}

	return ctx.Render("frontend/home", fiber.Map{
		"meta":          gjson.Parse(configs["meta"]).Value(),
		"cover":         gjson.Parse(configs["cover"]).Value(),
		"event":         gjson.Parse(configs["event"]).Value(),
		"story":         gjson.Parse(configs["story"]).Value(),
		"venue":         gjson.Parse(configs["venue"]).Value(),
		"rsvp":          gjson.Parse(configs["rsvp"]).Value(),
		"gift":          gjson.Parse(configs["gift"]).Value(),
		"wishes":        gjson.Parse(configs["wishes"]).Value(),
		"thanks":        gjson.Parse(configs["thanks"]).Value(),
		"couples":       gjson.Parse(string(resCouples)).Value(),
		"venue_details": gjson.Parse(string(resVenueDetails)).Value(),
		"galleries":     gjson.Parse(string(resGalleries)).Value(),
		"gifts":         gjson.Parse(string(resGifts)).Value(),
		"guest":         gjson.Parse(resGuest).Value(),
		"config_app":    ctrl.cfgApp,
	})
}
