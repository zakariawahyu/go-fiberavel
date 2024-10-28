package controller

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/tidwall/gjson"
	"github.com/zakariawahyu/go-fiberavel/app/repository"
	"github.com/zakariawahyu/go-fiberavel/config"
	"github.com/zakariawahyu/go-fiberavel/internal/utils/constants"
	"time"
)

type HomeController struct {
	homeRepo repository.HomeRepository
	cfgApp   config.App
}

func NewHomeController(homeRepo repository.HomeRepository, cfgApp config.App) *HomeController {
	return &HomeController{
		homeRepo: homeRepo,
		cfgApp:   cfgApp,
	}
}

func (ctrl *HomeController) Index(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), ctrl.cfgApp.Timeout*time.Second)
	defer cancel()

	configs, err := ctrl.homeRepo.HGetAll(c, constants.KeyConfigs)
	if err != nil {
		return err
	}

	resCouples, err := ctrl.homeRepo.Get(c, constants.KeyCouples)
	if err != nil {
		return err
	}

	resVenueDetails, err := ctrl.homeRepo.Get(c, constants.KeyVenueDetails)
	if err != nil {
		return err
	}

	resGalleries, err := ctrl.homeRepo.Get(c, constants.KeyGalleries)
	if err != nil {
		return err
	}

	resGifts, err := ctrl.homeRepo.Get(c, constants.KeyGift)
	if err != nil {
		return err
	}

	resGuest, err := ctrl.homeRepo.HGet(c, constants.KeyGuestList, "akbar-gustama")
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
		"couples":       gjson.Parse(resCouples).Value(),
		"venue_details": gjson.Parse(resVenueDetails).Value(),
		"galleries":     gjson.Parse(resGalleries).Value(),
		"gifts":         gjson.Parse(resGifts).Value(),
		"guest":         gjson.Parse(resGuest).Value(),
		"config_app":    ctrl.cfgApp,
	})
}
