package controller

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/zakariawahyu/go-fiberavel/app/usecase"
	"github.com/zakariawahyu/go-fiberavel/config"
	"github.com/zakariawahyu/go-fiberavel/internal/infrastructure/cache"
	"sync"
	"time"
)

type HomeController struct {
	homeUsecase usecase.HomeUsecase
	redis       cache.Rueidis
	cfgApp      config.App
}

func NewHomeController(homeUsecase usecase.HomeUsecase, redis cache.Rueidis, cfgApp config.App) *HomeController {
	return &HomeController{
		homeUsecase: homeUsecase,
		redis:       redis,
		cfgApp:      cfgApp,
	}
}

func (ctrl *HomeController) Index(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), ctrl.cfgApp.Timeout*time.Second)
	defer cancel()

	guest := ctx.Params("guest")

	// Go routine to get all configs
	wg := new(sync.WaitGroup)
	chanConfig := make(chan map[string]interface{}, 1)
	chanErr := make(chan error, 1)

	go ctrl.homeUsecase.GetConfigs(c, wg, chanConfig, chanErr)
	wg.Wait()

	configs, err := <-chanConfig, <-chanErr
	if err != nil {
		return err
	}

	chanData := make(chan map[string]interface{}, 1)
	chanErrData := make(chan error, 1)
	go ctrl.homeUsecase.GetData(wg, chanData, chanErrData)
	wg.Wait()

	data, err := <-chanData, <-chanErrData
	if err != nil {
		return err
	}

	chanHget := make(chan map[string]interface{}, 1)
	chanErrHget := make(chan error, 1)
	go ctrl.homeUsecase.Hget(guest, wg, chanHget, chanErrHget)
	wg.Wait()

	guestData, err := <-chanHget, <-chanErrHget
	if err != nil {
		return err
	}

	return ctx.Render("frontend/home", fiber.Map{
		"meta":          configs["meta"],
		"cover":         configs["cover"],
		"event":         configs["event"],
		"story":         configs["story"],
		"venue":         configs["venue"],
		"rsvp":          configs["rsvp"],
		"gift":          configs["gift"],
		"wish":          configs["wish"],
		"thank":         configs["thank"],
		"couples":       data["couples"],
		"venue_details": data["venue_details"],
		"galleries":     data["galleries"],
		"gifts":         data["gifts"],
		"guest":         guestData["guest"],
		"config_app":    ctrl.cfgApp,
	})
}
