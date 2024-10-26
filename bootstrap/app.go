package bootstrap

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/zakariawahyu/go-fiberavel/config"
	"github.com/zakariawahyu/go-fiberavel/internal/infrastructure/cache"
	"github.com/zakariawahyu/go-fiberavel/internal/utils"
	"github.com/zakariawahyu/go-fiberavel/routes"
)

func NewApplication() *fiber.App {
	app := fiber.New(fiber.Config{
		ErrorHandler: utils.ErrorHandler,
	})

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	redis, err := cache.NewRedis(cfg)
	if err != nil {
		log.Fatal(err)
	}

	routes.WebRoutes(app, cfg, redis)

	log.Fatal(app.Listen(cfg.App.Port))
	return app
}
