package routes

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/zakariawahyu/go-fiberavel/app/http/controller"
	admin "github.com/zakariawahyu/go-fiberavel/app/http/controller/admin"
	"github.com/zakariawahyu/go-fiberavel/app/http/middleware"
	repository "github.com/zakariawahyu/go-fiberavel/app/repository/admin"
	usecase "github.com/zakariawahyu/go-fiberavel/app/usecase/admin"
	"github.com/zakariawahyu/go-fiberavel/config"
	"github.com/zakariawahyu/go-fiberavel/internal/infrastructure/cache"
	sqlc "github.com/zakariawahyu/go-fiberavel/internal/sqlc/generated"
)

func WebRoutes(app *fiber.App, cfg *config.Config, db *sqlc.Queries, redis *cache.Storage, session *middleware.Session, validator *validator.Validate) {
	ctrlHome := controller.NewHomeController(redis, cfg.App)

	app.Get("/", ctrlHome.Index)

	// Route Backend
	repoAuth := repository.NewAuthRepository(db)
	usecaseAuth := usecase.NewAuthhUsecase(repoAuth)
	ctrlAuth := admin.NewAuthController(usecaseAuth, cfg.App, session, validator)

	app.Get("/auth/mimin", ctrlAuth.Index)
	app.Post("/auth/mimin", ctrlAuth.Login)
	app.Get("/auth/unauthorized", ctrlAuth.Unauthorized)

	mimin := app.Group("/mimin", session.Authenticated())
	repoCouple := repository.NewCoupleRepository(db)

	usecaseCouple := usecase.NewCoupleUsecase(repoCouple)

	ctrlDashboard := admin.NewDashboardController()
	ctrlCouple := admin.NewCoupleController(usecaseCouple, cfg.App, session, validator)

	mimin.Get("/logout", ctrlAuth.Logout)
	mimin.Get("/dashboard", ctrlDashboard.Index)

	mimin.Get("/couple", ctrlCouple.Index)
	mimin.Get("/couple/create", ctrlCouple.Create)
	mimin.Post("/couple/store", ctrlCouple.Store)
	mimin.Get("/couple/datatables", ctrlCouple.Datatables)
}
