package controller

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/tidwall/gjson"
	"github.com/zakariawahyu/go-fiberavel/config"
	"github.com/zakariawahyu/go-fiberavel/internal/infrastructure/cache"
	"github.com/zakariawahyu/go-fiberavel/internal/utils/constants"
	"sync"
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
	log.Info("Start the proccess")

	type task struct {
		key   string
		value interface{}
		err   error
	}

	tasks := []struct {
		name string
		work func() (interface{}, error)
	}{
		{"configs", func() (interface{}, error) {
			return ctrl.redis.HGetAll(c, constants.KeyConfigs)
		}},
		{"couples", func() (interface{}, error) {
			return ctrl.redis.Get(constants.KeyCouples)
		}},
		{"venue_details", func() (interface{}, error) {
			return ctrl.redis.Get(constants.KeyVenues)
		}},
		{"galleries", func() (interface{}, error) {
			return ctrl.redis.Get(constants.KeyGalleries)
		}},
		{"gifts", func() (interface{}, error) {
			return ctrl.redis.Get(constants.KeyGift)
		}},
		{"guest", func() (interface{}, error) {
			return ctrl.redis.HGet(constants.KeyGuests, guest)
		}},
	}

	// Worker pool setup
	numWorkers := 3
	taskChan := make(chan struct {
		name string
		work func() (interface{}, error)
	})
	resultChan := make(chan task, len(tasks))

	var wg sync.WaitGroup

	// Start workers
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for t := range taskChan {
				log.Infof("Processing task: %s in workers %v", t.name, i)
				value, err := t.work()
				resultChan <- task{key: t.name, value: value, err: err}
			}
		}()
	}

	// Assign tasks
	go func() {
		for _, t := range tasks {
			taskChan <- t
		}
		close(taskChan)
	}()

	// Wait for workers to complete
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	// Collect results
	results := make(map[string]interface{})
	for r := range resultChan {
		if r.err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, r.err.Error())
		}
		results[r.key] = r.value
	}

	// Parsing JSON
	configs, ok := results["configs"].(map[string]string)
	if !ok || configs == nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to fetch configs")
	}

	return ctx.Render("frontend/home", fiber.Map{
		"meta":          gjson.Parse(configs["meta"]).Value(),
		"cover":         gjson.Parse(configs["cover"]).Value(),
		"event":         gjson.Parse(configs["event"]).Value(),
		"story":         gjson.Parse(configs["story"]).Value(),
		"venue":         gjson.Parse(configs["venue"]).Value(),
		"rsvp":          gjson.Parse(configs["rsvp"]).Value(),
		"gift":          gjson.Parse(configs["gift"]).Value(),
		"wish":          gjson.Parse(configs["wish"]).Value(),
		"thank":         gjson.Parse(configs["thank"]).Value(),
		"couples":       ctrl.safeParseString(results["couples"]),
		"venue_details": ctrl.safeParseString(results["venue_details"]),
		"galleries":     ctrl.safeParseString(results["galleries"]),
		"gifts":         ctrl.safeParseString(results["gifts"]),
		"guest":         ctrl.safeParseString(results["guest"]),
		"config_app":    ctrl.cfgApp,
	})
}

func (ctrl *HomeController) safeParseString(val interface{}) interface{} {
	str, ok := val.(string)
	if !ok || str == "" {
		return nil
	}
	return gjson.Parse(str).Value()
}
