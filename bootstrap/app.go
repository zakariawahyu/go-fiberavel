package bootstrap

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/template/jet/v2"
	"github.com/zakariawahyu/go-fiberavel/config"
	"github.com/zakariawahyu/go-fiberavel/internal/infrastructure/cache"
	"github.com/zakariawahyu/go-fiberavel/internal/infrastructure/db"
	sqlc "github.com/zakariawahyu/go-fiberavel/internal/sqlc/generated"
	"github.com/zakariawahyu/go-fiberavel/internal/utils"
	"github.com/zakariawahyu/go-fiberavel/routes"
)

func NewApplication() *fiber.App {
	engine := jet.New("./resources/views", ".jet.html")

	app := fiber.New(fiber.Config{
		Views:        engine,
		ErrorHandler: utils.ErrorHandler,
	})

	app.Static("/assets", "./public/assets")

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	redis, err := cache.NewRedis(cfg)
	if err != nil {
		log.Fatal(err)
	}

	postgres, err := db.NewPostgres(cfg)
	if err != nil {
		log.Fatal(err)
	}

	queries := sqlc.New(postgres.Conn)

	routes.WebRoutes(app, cfg, redis)
	routes.ApiRoutes(app, cfg, queries, redis)

	log.Fatal(app.Listen(cfg.App.Port))
	return app
}
