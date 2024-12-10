package controller

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/tidwall/gjson"
	"github.com/zakariawahyu/go-fiberavel/config"
	"github.com/zakariawahyu/go-fiberavel/internal/infrastructure/cache"
	"github.com/zakariawahyu/go-fiberavel/internal/utils/constants"
	"github.com/zakariawahyu/go-fiberavel/internal/utils/helper"
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

	guest := ctx.Params("guest")

	configs, err := ctrl.redis.HGetAll(c, constants.KeyConfigs)
	if err != nil {
		return err
	}

	resCouples, err := ctrl.redis.Get(constants.KeyCouples)
	if err != nil {
		return err
	}

	resVenueDetails, err := ctrl.redis.Get(constants.KeyVenues)
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

	resGuest, err := ctrl.redis.HGet(constants.KeyGuests, guest)
	if err != nil {
		return err
	}

	venues := gjson.Parse(string(resVenueDetails)).Array()
	VenueDetails := make(map[int]interface{})

	for key, value := range venues {
		jam, hari, tanggal := helper.ParseDate(value.Get("date_held").String())
		VenueDetails[key] = fiber.Map{
			"id":       value.Get("id").Int(),
			"name":     value.Get("name").String(),
			"location": value.Get("location").String(),
			"address":  value.Get("address").String(),
			"map":      value.Get("map").String(),
			"jam":      jam,
			"hari":     hari,
			"tanggal":  tanggal,
		}
	}

	event := gjson.Parse(configs["event"]).Map()
	eventDetail := fiber.Map{
		"title":         event["title"].String(),
		"description":   event["description"].String(),
		"image":         event["image"].String(),
		"image_caption": event["image_caption"].String(),
		"custom_data": fiber.Map{
			"date": helper.ParseUTC(event["custom_data"].Get("date").String()),
		},
		"is_active": event["is_active"].Bool(),
	}

	return ctx.Render("frontend/home", fiber.Map{
		"meta":          gjson.Parse(configs["meta"]).Value(),
		"cover":         gjson.Parse(configs["cover"]).Value(),
		"event":         eventDetail,
		"story":         gjson.Parse(configs["story"]).Value(),
		"venue":         gjson.Parse(configs["venue"]).Value(),
		"rsvp":          gjson.Parse(configs["rsvp"]).Value(),
		"gift":          gjson.Parse(configs["gift"]).Value(),
		"wish":          gjson.Parse(configs["wish"]).Value(),
		"thank":         gjson.Parse(configs["thank"]).Value(),
		"couples":       gjson.Parse(string(resCouples)).Value(),
		"venue_details": VenueDetails,
		"galleries":     gjson.Parse(string(resGalleries)).Value(),
		"gifts":         gjson.Parse(string(resGifts)).Value(),
		"guest":         gjson.Parse(resGuest).Value(),
		"config_app":    ctrl.cfgApp,
	})
}
