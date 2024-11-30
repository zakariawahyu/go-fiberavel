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

type resourceRoutes struct {
	Index      fiber.Handler
	Store      fiber.Handler
	Create     fiber.Handler
	Publish    fiber.Handler
	Datatables fiber.Handler
	Show       fiber.Handler
	Update     fiber.Handler
	Edit       fiber.Handler
	Destroy    fiber.Handler
}

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
	usecaseCouple := usecase.NewCoupleUsecase(repoCouple, redis)
	ctrlCouple := admin.NewCoupleController(usecaseCouple, cfg.App, session, validator)

	repoVenue := repository.NewVenueRepository(db)
	usecaseVenue := usecase.NewVenueUsecase(repoVenue, redis)
	ctrlVenue := admin.NewVenueController(usecaseVenue, cfg.App, session, validator)

	repoGallery := repository.NewGalleryRepository(db)
	usecaseGallery := usecase.NewGalleryUsecase(repoGallery, redis)
	ctrlGallery := admin.NewGalleryController(usecaseGallery, cfg.App, session, validator)

	ctrlDashboard := admin.NewDashboardController(session)
	mimin.Get("/logout", ctrlAuth.Logout)
	mimin.Get("/dashboard", ctrlDashboard.Index)

	registerResources(mimin, "couple", resourceRoutes{
		Index:      ctrlCouple.Index,
		Store:      ctrlCouple.Store,
		Create:     ctrlCouple.Create,
		Publish:    ctrlCouple.Publish,
		Datatables: ctrlCouple.Datatables,
		Show:       ctrlCouple.Show,
		Update:     ctrlCouple.Update,
		Edit:       ctrlCouple.Edit,
		Destroy:    ctrlCouple.Destroy,
	})

	registerResources(mimin, "venue", resourceRoutes{
		Index:      ctrlVenue.Index,
		Store:      ctrlVenue.Store,
		Create:     ctrlVenue.Create,
		Publish:    ctrlVenue.Publish,
		Datatables: ctrlVenue.Datatables,
		Show:       ctrlVenue.Show,
		Update:     ctrlVenue.Update,
		Edit:       ctrlVenue.Edit,
		Destroy:    ctrlVenue.Destroy,
	})

	registerResources(mimin, "gallery", resourceRoutes{
		Index:      ctrlGallery.Index,
		Store:      ctrlGallery.Store,
		Create:     ctrlGallery.Create,
		Publish:    ctrlGallery.Publish,
		Datatables: ctrlGallery.Datatables,
		Show:       ctrlGallery.Show,
		Update:     ctrlGallery.Update,
		Edit:       ctrlGallery.Edit,
		Destroy:    ctrlGallery.Destroy,
	})
}

func registerResources(group fiber.Router, resources string, handler resourceRoutes) {
	group.Get(resources, handler.Index)
	group.Post(resources, handler.Store)
	group.Get(resources+"/create", handler.Create)
	group.Get(resources+"/publish", handler.Publish)
	group.Get(resources+"/datatables", handler.Datatables)
	group.Get(resources+"/:id", handler.Show)
	group.Post(resources+"/:id", handler.Update)
	group.Get(resources+"/:id/edit", handler.Edit)
	group.Get(resources+"/:id/delete", handler.Destroy)
}
