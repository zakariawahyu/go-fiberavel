package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zakariawahyu/go-fiberavel/app/http/controller"
	admin "github.com/zakariawahyu/go-fiberavel/app/http/controller/admin"
	"github.com/zakariawahyu/go-fiberavel/app/http/middleware"
	repository "github.com/zakariawahyu/go-fiberavel/app/repository/admin"
	"github.com/zakariawahyu/go-fiberavel/config"
	"github.com/zakariawahyu/go-fiberavel/internal/infrastructure/cache"
	sqlc "github.com/zakariawahyu/go-fiberavel/internal/sqlc/generated"
)

func WebRoutes(app *fiber.App, cfg *config.Config, db *sqlc.Queries, redis *cache.Storage, session *middleware.Session) {
	ctrlHome := controller.NewHomeController(redis, cfg.App)

	app.Get("/", ctrlHome.Index)

	// Route Backend
	mimin := app.Group("/mimin")

	repoAuth := repository.NewAuthRepository(db)
	ctrlDashboard := admin.NewDashboardController()
	ctrlAuth := admin.NewAuthController(repoAuth, cfg.App, session)

	mimin.Get("/login", ctrlAuth.Index)
	mimin.Post("/login", ctrlAuth.Login)
	mimin.Get("/dashboard", ctrlDashboard.Index)
}
