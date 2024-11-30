package routes

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/zakariawahyu/go-fiberavel/app/http/controller"
	"github.com/zakariawahyu/go-fiberavel/app/repository"
	"github.com/zakariawahyu/go-fiberavel/app/usecase"
	"github.com/zakariawahyu/go-fiberavel/config"
	"github.com/zakariawahyu/go-fiberavel/internal/infrastructure/cache"
	sqlc "github.com/zakariawahyu/go-fiberavel/internal/sqlc/generated"
)

func ApiRoutes(app *fiber.App, cfg *config.Config, db *sqlc.Queries, redis *cache.Storage, validator *validator.Validate) {
	repoWish := repository.NewWishesRepository(db)
	usecaseWish := usecase.NewWishUsecase(repoWish, redis)
	ctrlWishes := controller.NewWishController(usecaseWish, cfg.App, validator)

	repoRsvp := repository.NewRsvpRepository(db)
	ctrlRsvp := controller.NewRsvpController(repoRsvp, cfg.App)

	api := app.Group("/api")
	api.Post("/wish", ctrlWishes.Store)
	api.Get("/wish", ctrlWishes.GetAll)
	api.Post("/rsvp", ctrlRsvp.CreateRsvp)
}
