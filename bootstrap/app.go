package bootstrap

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/template/jet/v2"
	"github.com/zakariawahyu/go-fiberavel/app/http/middleware"
	"github.com/zakariawahyu/go-fiberavel/config"
	"github.com/zakariawahyu/go-fiberavel/internal/infrastructure/cache"
	"github.com/zakariawahyu/go-fiberavel/internal/infrastructure/db"
	sqlc "github.com/zakariawahyu/go-fiberavel/internal/sqlc/generated"
	"github.com/zakariawahyu/go-fiberavel/routes"
)

func NewApplication() *fiber.App {
	// Load configuration from .env file
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	// Initialize Redis Client
	redis, err := cache.NewRedis(cfg)
	if err != nil {
		log.Fatal(err)
	}

	// Initialize Postgres Connection
	postgres, err := db.NewPostgres(cfg)
	if err != nil {
		log.Fatal(err)
	}
	// Register SQLC Queries
	queries := sqlc.New(postgres.Conn)

	// Initialize Jet Template Engine
	engine := jet.New("./resources/views", ".jet.html")

	// Initialize Fiber App
	app := fiber.New(fiber.Config{
		Views:        engine,
		ErrorHandler: middleware.ErrorHandler,
	})

	// Register Static File
	app.Static("/assets", "./public/assets")

	// Initialize Redis Store
	cfg.Redis.SelectDB = 1
	redisStore, err := cache.NewRedis(cfg)
	if err != nil {
		log.Fatal(err)
	}

	// Initialize Session Store and Register CSRF Middleware
	middleware.InitSessionsStore(redisStore)
	app.Use(middleware.CSRFMiddleware(middleware.Store))

	// Initialize Cache Middleware
	app.Use(middleware.CacheMiddleware(redisStore))

	// Register Routes
	routes.WebRoutes(app, cfg, redis)
	routes.ApiRoutes(app, cfg, queries, redis)

	// Start Fiber App
	log.Fatal(app.Listen(cfg.App.Port))

	return app
}
