package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zakariawahyu/go-fiberavel/app/http/controller"
	admin "github.com/zakariawahyu/go-fiberavel/app/http/controller/admin"
	"github.com/zakariawahyu/go-fiberavel/app/repository"
	"github.com/zakariawahyu/go-fiberavel/config"
	"github.com/zakariawahyu/go-fiberavel/internal/infrastructure/cache"
)

func WebRoutes(app *fiber.App, cfg *config.Config, redis *cache.Storage) {
	repoHome := repository.NewHomeRepository(redis)
	ctrlHome := controller.NewHomeController(repoHome, cfg.App)

	app.Get("/", ctrlHome.Index)

	// Route Backend
	mimin := app.Group("/mimin")

	ctrlDashboard := admin.NewDashboardController()
	mimin.Get("/dashboard", ctrlDashboard.Index)
}
