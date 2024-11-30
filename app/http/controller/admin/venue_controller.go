package controller

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/zakariawahyu/go-fiberavel/app/http/middleware"
	"github.com/zakariawahyu/go-fiberavel/app/http/request"
	usecase "github.com/zakariawahyu/go-fiberavel/app/usecase/admin"
	"github.com/zakariawahyu/go-fiberavel/config"
	"github.com/zakariawahyu/go-fiberavel/internal/utils/datatables"
	"github.com/zakariawahyu/go-fiberavel/internal/utils/flash"
	"github.com/zakariawahyu/go-fiberavel/internal/utils/helper"
	"time"
)

type VenueController struct {
	venueUsecase usecase.VenueUsecase
	cfgApp       config.App
	session      *middleware.Session
	validator    *validator.Validate
}

func NewVenueController(venueUsecase usecase.VenueUsecase, cfgApp config.App, session *middleware.Session, validator *validator.Validate) *VenueController {
	return &VenueController{
		venueUsecase: venueUsecase,
		cfgApp:       cfgApp,
		session:      session,
		validator:    validator,
	}
}

func (c *VenueController) Index(ctx *fiber.Ctx) error {
	build := flash.NewMessage(c.session.Store).Build()
	return ctx.Render("backend/pages/venue/index", build.GetFlash(ctx))
}

func (c *VenueController) Datatables(ctx *fiber.Ctx) error {
	params, err := datatables.ParseDataTableParams(ctx)
	if err != nil {
		return err
	}

	venues, err := c.venueUsecase.Datatables(ctx.Context(), params)
	if err != nil {
		return err
	}

	return ctx.JSON(venues)
}

func (c *VenueController) Create(ctx *fiber.Ctx) error {
	build := flash.NewMessage(c.session.Store).Build()
	return ctx.Render("backend/pages/venue/create", build.GetFlash(ctx))
}

func (c *VenueController) Store(ctx *fiber.Ctx) error {
	ctxTimeout, cancel := context.WithTimeout(ctx.Context(), c.cfgApp.Timeout*time.Second)
	defer cancel()

	var req request.CreateVenueRequest

	if err := ctx.BodyParser(&req); err != nil {
		return flash.HandleError(ctx, c.session.Store, err, req)
	}

	if err := c.validator.Struct(&req); err != nil {
		return flash.HandleValidationError(ctx, c.session.Store, err, req)
	}

	_, err := c.venueUsecase.Store(ctxTimeout, req)
	if err != nil {
		return flash.HandleError(ctx, c.session.Store, err, req)
	}

	return flash.HandleSuccess(ctx, c.session.Store, "Venue created successfully", "/mimin/venue")
}

func (c *VenueController) Show(ctx *fiber.Ctx) error {
	ctxTimeout, cancel := context.WithTimeout(ctx.Context(), c.cfgApp.Timeout*time.Second)
	defer cancel()

	id, err := helper.StrToInt64(ctx.Params("id"))
	if err != nil {
		return flash.HandleError(ctx, c.session.Store, err, nil)
	}

	venue, err := c.venueUsecase.FindById(ctxTimeout, id)
	if err != nil {
		return flash.HandleError(ctx, c.session.Store, err, nil)
	}

	return ctx.Render("backend/pages/venue/show", fiber.Map{
		"venue": venue,
	})
}

func (c *VenueController) Edit(ctx *fiber.Ctx) error {
	ctxTimeout, cancel := context.WithTimeout(ctx.Context(), c.cfgApp.Timeout*time.Second)
	defer cancel()

	id, err := helper.StrToInt64(ctx.Params("id"))
	if err != nil {
		return flash.HandleError(ctx, c.session.Store, err, nil)
	}

	venue, err := c.venueUsecase.FindById(ctxTimeout, id)
	if err != nil {
		return flash.HandleError(ctx, c.session.Store, err, nil)
	}

	build := flash.NewMessage(c.session.Store).Build()
	return ctx.Render("backend/pages/venue/edit", helper.Compact(fiber.Map{
		"venue": venue,
	}, build.GetFlash(ctx)))
}

func (c *VenueController) Update(ctx *fiber.Ctx) error {
	ctxTimeout, cancel := context.WithTimeout(ctx.Context(), c.cfgApp.Timeout*time.Second)
	defer cancel()

	if ctx.FormValue("_method") == "PUT" {
		ctx.Method("PUT")
	}

	var req request.UpdateVenueRequest

	id, err := helper.StrToInt64(ctx.Params("id"))
	if err != nil {
		return flash.HandleError(ctx, c.session.Store, err, nil)
	}

	req.ID = id
	if err := ctx.BodyParser(&req); err != nil {
		return flash.HandleError(ctx, c.session.Store, err, req)
	}

	if err := c.validator.Struct(&req); err != nil {
		return flash.HandleValidationError(ctx, c.session.Store, err, req)
	}

	if err := c.venueUsecase.Update(ctxTimeout, req); err != nil {
		return flash.HandleError(ctx, c.session.Store, err, req)
	}

	return flash.HandleSuccess(ctx, c.session.Store, "Venue updated successfully", "/mimin/venue")
}

func (c *VenueController) Destroy(ctx *fiber.Ctx) error {
	ctxTimeout, cancel := context.WithTimeout(ctx.Context(), c.cfgApp.Timeout*time.Second)
	defer cancel()

	id, err := helper.StrToInt64(ctx.Params("id"))
	if err != nil {
		return flash.HandleError(ctx, c.session.Store, err, nil)
	}

	if err := c.venueUsecase.Destroy(ctxTimeout, id); err != nil {
		return flash.HandleError(ctx, c.session.Store, err, nil)
	}

	return flash.HandleSuccess(ctx, c.session.Store, "Venue deleted successfully", "/mimin/venue")
}

func (c *VenueController) Publish(ctx *fiber.Ctx) error {
	ctxTimeout, cancel := context.WithTimeout(ctx.Context(), c.cfgApp.Timeout*time.Second)
	defer cancel()

	if err := c.venueUsecase.Publish(ctxTimeout); err != nil {
		return flash.HandleError(ctx, c.session.Store, err, nil)
	}

	return flash.HandleSuccess(ctx, c.session.Store, "Venue published successfully", "/mimin/venue")
}
