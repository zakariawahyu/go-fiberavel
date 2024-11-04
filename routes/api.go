package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zakariawahyu/go-fiberavel/app/http/controller"
	"github.com/zakariawahyu/go-fiberavel/app/repository"
	"github.com/zakariawahyu/go-fiberavel/config"
	"github.com/zakariawahyu/go-fiberavel/internal/infrastructure/cache"
	sqlc "github.com/zakariawahyu/go-fiberavel/internal/sqlc/generated"
)

func ApiRoutes(app *fiber.App, cfg *config.Config, db *sqlc.Queries, redis *cache.Storage) {
	repoWishes := repository.NewWishesRepository(db, redis)
	ctrlWishes := controller.NewWishController(repoWishes, cfg.App)

	repoRsvp := repository.NewRsvpRepository(db)
	ctrlRsvp := controller.NewRsvpController(repoRsvp, cfg.App)

	api := app.Group("/api")
	api.Post("/wish", ctrlWishes.CreateWish)
	api.Get("/wish", ctrlWishes.GetAllWishes)
	api.Post("/rsvp", ctrlRsvp.CreateRsvp)
}
