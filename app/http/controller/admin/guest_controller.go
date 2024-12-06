package controller

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gosimple/slug"
	"github.com/zakariawahyu/go-fiberavel/app/http/middleware"
	"github.com/zakariawahyu/go-fiberavel/app/http/request"
	usecase "github.com/zakariawahyu/go-fiberavel/app/usecase/admin"
	"github.com/zakariawahyu/go-fiberavel/config"
	"github.com/zakariawahyu/go-fiberavel/internal/utils/datatables"
	"github.com/zakariawahyu/go-fiberavel/internal/utils/flash"
	"github.com/zakariawahyu/go-fiberavel/internal/utils/helper"
	"time"
)

type GuestController struct {
	guestUsecase usecase.GuestUsecase
	cfgApp       config.App
	session      *middleware.Session
	validator    *validator.Validate
}

func NewGuestController(guestUsecase usecase.GuestUsecase, cfgApp config.App, session *middleware.Session, validator *validator.Validate) *GuestController {
	return &GuestController{
		guestUsecase: guestUsecase,
		cfgApp:       cfgApp,
		session:      session,
		validator:    validator,
	}
}

func (c *GuestController) Index(ctx *fiber.Ctx) error {
	build := flash.NewMessage(c.session.Store).Build()
	return ctx.Render("backend/pages/guest/index", build.GetFlash(ctx))
}

func (c *GuestController) Datatables(ctx *fiber.Ctx) error {
	params, err := datatables.ParseDataTableParams(ctx)
	if err != nil {
		return err
	}

	guests, err := c.guestUsecase.Datatables(ctx.Context(), params)
	if err != nil {
		return err
	}

	return ctx.JSON(guests)
}

func (c *GuestController) Create(ctx *fiber.Ctx) error {
	build := flash.NewMessage(c.session.Store).Build()
	return ctx.Render("backend/pages/guest/create", build.GetFlash(ctx))
}

func (c *GuestController) Store(ctx *fiber.Ctx) error {
	ctxTimeout, cancel := context.WithTimeout(ctx.Context(), c.cfgApp.Timeout*time.Second)
	defer cancel()

	var req request.CreateGuestRequest

	if err := ctx.BodyParser(&req); err != nil {
		return flash.HandleError(ctx, c.session.Store, err, req)
	}

	if err := c.validator.Struct(&req); err != nil {
		return flash.HandleValidationError(ctx, c.session.Store, err, req)
	}

	req.Slug = slug.Make(req.Name)
	if err := c.guestUsecase.Store(ctxTimeout, req); err != nil {
		return flash.HandleError(ctx, c.session.Store, err, req)
	}

	return flash.HandleSuccess(ctx, c.session.Store, "Guest created successfully", "/mimin/guest")
}

func (c *GuestController) Show(ctx *fiber.Ctx) error {
	ctxTimeout, cancel := context.WithTimeout(ctx.Context(), c.cfgApp.Timeout*time.Second)
	defer cancel()

	id, err := helper.StrToInt64(ctx.Params("id"))
	if err != nil {
		return flash.HandleError(ctx, c.session.Store, err, nil)
	}

	guest, err := c.guestUsecase.FindById(ctxTimeout, id)
	if err != nil {
		return flash.HandleError(ctx, c.session.Store, err, nil)
	}

	return ctx.Render("backend/pages/guest/show", fiber.Map{
		"guest": guest,
	})
}

func (c *GuestController) Edit(ctx *fiber.Ctx) error {
	ctxTimeout, cancel := context.WithTimeout(ctx.Context(), c.cfgApp.Timeout*time.Second)
	defer cancel()

	id, err := helper.StrToInt64(ctx.Params("id"))
	if err != nil {
		return flash.HandleError(ctx, c.session.Store, err, nil)
	}

	guest, err := c.guestUsecase.FindById(ctxTimeout, id)
	if err != nil {
		return flash.HandleError(ctx, c.session.Store, err, nil)
	}

	build := flash.NewMessage(c.session.Store).Build()
	return ctx.Render("backend/pages/guest/edit", helper.Compact(fiber.Map{
		"guest": guest,
	}, build.GetFlash(ctx)))
}

func (c *GuestController) Update(ctx *fiber.Ctx) error {
	ctxTimeout, cancel := context.WithTimeout(ctx.Context(), c.cfgApp.Timeout*time.Second)
	defer cancel()

	if ctx.FormValue("_method") == "PUT" {
		ctx.Method("PUT")
	}

	var req request.UpdateGuestRequest

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

	req.Slug = slug.Make(req.Name)
	if err := c.guestUsecase.Update(ctxTimeout, req); err != nil {
		return flash.HandleError(ctx, c.session.Store, err, req)
	}

	return flash.HandleSuccess(ctx, c.session.Store, "Guest updated successfully", "/mimin/guest")
}

func (c *GuestController) Destroy(ctx *fiber.Ctx) error {
	ctxTimeout, cancel := context.WithTimeout(ctx.Context(), c.cfgApp.Timeout*time.Second)
	defer cancel()

	id, err := helper.StrToInt64(ctx.Params("id"))
	if err != nil {
		return flash.HandleError(ctx, c.session.Store, err, nil)
	}

	if err := c.guestUsecase.Destroy(ctxTimeout, id); err != nil {
		return flash.HandleError(ctx, c.session.Store, err, nil)
	}

	return flash.HandleSuccess(ctx, c.session.Store, "Guest deleted successfully", "/mimin/guest")
}

func (c *GuestController) Publish(ctx *fiber.Ctx) error {
	ctxTimeout, cancel := context.WithTimeout(ctx.Context(), c.cfgApp.Timeout*time.Second)
	defer cancel()

	if err := c.guestUsecase.Publish(ctxTimeout); err != nil {
		return flash.HandleError(ctx, c.session.Store, err, nil)
	}

	return flash.HandleSuccess(ctx, c.session.Store, "Guest published successfully", "/mimin/guest")
}
