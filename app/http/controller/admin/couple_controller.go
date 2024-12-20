package controller

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/zakariawahyu/go-fiberavel/app/http/middleware"
	"github.com/zakariawahyu/go-fiberavel/app/http/request"
	usecase "github.com/zakariawahyu/go-fiberavel/app/usecase/admin"
	"github.com/zakariawahyu/go-fiberavel/config"
	"github.com/zakariawahyu/go-fiberavel/internal/utils/constants"
	"github.com/zakariawahyu/go-fiberavel/internal/utils/datatables"
	"github.com/zakariawahyu/go-fiberavel/internal/utils/flash"
	"github.com/zakariawahyu/go-fiberavel/internal/utils/helper"
	"time"
)

type CoupleController struct {
	coupleUsecase usecase.CoupleUsecase
	cfgApp        config.App
	session       *middleware.Session
	validator     *validator.Validate
}

func NewCoupleController(coupleUsecase usecase.CoupleUsecase, cfgApp config.App, session *middleware.Session, validator *validator.Validate) *CoupleController {
	return &CoupleController{
		coupleUsecase: coupleUsecase,
		cfgApp:        cfgApp,
		session:       session,
		validator:     validator,
	}
}

func (c *CoupleController) Index(ctx *fiber.Ctx) error {
	build := flash.NewMessage(c.session.Store).Build()
	return ctx.Render("backend/pages/couple/index", build.GetFlash(ctx))
}

func (c *CoupleController) Datatables(ctx *fiber.Ctx) error {
	ctxTimeout, cancel := context.WithTimeout(ctx.Context(), c.cfgApp.Timeout*time.Second)
	defer cancel()

	params, err := datatables.ParseDataTableParams(ctx)
	if err != nil {
		return err
	}

	couples, err := c.coupleUsecase.Datatables(ctxTimeout, params)
	if err != nil {
		return err
	}

	return ctx.JSON(couples)
}

func (c *CoupleController) Create(ctx *fiber.Ctx) error {
	build := flash.NewMessage(c.session.Store).Build()
	coupleTypes := constants.CoupleTypes

	return ctx.Render("backend/pages/couple/create", helper.Compact(fiber.Map{
		"coupleTypes": coupleTypes,
	}, build.GetFlash(ctx)))
}

func (c *CoupleController) Store(ctx *fiber.Ctx) error {
	ctxTimeout, cancel := context.WithTimeout(ctx.Context(), c.cfgApp.Timeout*time.Second)
	defer cancel()

	var req request.CreateCoupleRequest

	if err := ctx.BodyParser(&req); err != nil {
		return flash.HandleError(ctx, c.session.Store, err, req)
	}

	image, err := helper.GetImage(ctx, "image")
	if err != nil {
		return flash.HandleError(ctx, c.session.Store, err, req)
	}

	req.File = image
	if err := c.validator.Struct(&req); err != nil {
		return flash.HandleValidationError(ctx, c.session.Store, err, req)
	}

	imageName, err := helper.UploadImage(ctx, req.File, req.ImageCaption)
	if err != nil {
		return flash.HandleError(ctx, c.session.Store, err, req)
	}

	req.Image = imageName
	_, err = c.coupleUsecase.Store(ctxTimeout, req)
	if err != nil {
		return flash.HandleError(ctx, c.session.Store, err, req)
	}

	return flash.HandleSuccess(ctx, c.session.Store, "Couple successfully created", "/mimin/couple")
}

func (c *CoupleController) Show(ctx *fiber.Ctx) error {
	id, err := helper.StrToInt64(ctx.Params("id"))
	if err != nil {
		return flash.HandleError(ctx, c.session.Store, err, nil)
	}

	couple, err := c.coupleUsecase.FindByID(ctx.Context(), id)
	if err != nil {
		return flash.HandleError(ctx, c.session.Store, err, nil)
	}

	coupleTypes := constants.CoupleTypes

	return ctx.Render("backend/pages/couple/show", fiber.Map{
		"couple":      couple,
		"coupleTypes": coupleTypes,
	})
}

func (c *CoupleController) Edit(ctx *fiber.Ctx) error {
	ctxTimeout, cancel := context.WithTimeout(ctx.Context(), c.cfgApp.Timeout*time.Second)
	defer cancel()

	id, err := helper.StrToInt64(ctx.Params("id"))
	if err != nil {
		return flash.HandleError(ctx, c.session.Store, err, nil)
	}

	couple, err := c.coupleUsecase.FindByID(ctxTimeout, id)
	if err != nil {
		return flash.HandleError(ctx, c.session.Store, err, nil)
	}

	build := flash.NewMessage(c.session.Store).Build()
	coupleTypes := constants.CoupleTypes

	return ctx.Render("backend/pages/couple/edit", helper.Compact(fiber.Map{
		"couple":      couple,
		"coupleTypes": coupleTypes,
	}, build.GetFlash(ctx)))
}

func (c *CoupleController) Update(ctx *fiber.Ctx) error {
	ctxTimeout, cancel := context.WithTimeout(ctx.Context(), c.cfgApp.Timeout*time.Second)
	defer cancel()

	if ctx.FormValue("_method") == "PUT" {
		ctx.Method("PUT")
	}

	var req request.UpdateCoupleRequest

	id, err := helper.StrToInt64(ctx.Params("id"))
	if err != nil {
		return flash.HandleError(ctx, c.session.Store, err, req)
	}

	req.ID = id
	if err := ctx.BodyParser(&req); err != nil {
		return flash.HandleError(ctx, c.session.Store, err, req)
	}

	image, err := helper.GetImage(ctx, "image")
	if err != nil {
		return flash.HandleError(ctx, c.session.Store, err, req)
	}

	req.File = image
	if err := c.validator.Struct(&req); err != nil {
		return flash.HandleValidationError(ctx, c.session.Store, err, req)
	}

	if req.File != nil {
		imageName, err := helper.UploadImage(ctx, req.File, req.ImageCaption)
		if err != nil {
			return flash.HandleError(ctx, c.session.Store, err, req)
		}

		req.Image = imageName
	}

	if err = c.coupleUsecase.Update(ctxTimeout, req); err != nil {
		return flash.HandleError(ctx, c.session.Store, err, req)
	}

	return flash.HandleSuccess(ctx, c.session.Store, "Couple successfully updated", "/mimin/couple")
}

func (c *CoupleController) Destroy(ctx *fiber.Ctx) error {
	ctxTimeout, cancel := context.WithTimeout(ctx.Context(), c.cfgApp.Timeout*time.Second)
	defer cancel()

	id, err := helper.StrToInt64(ctx.Params("id"))
	if err != nil {
		return flash.HandleError(ctx, c.session.Store, err, nil)
	}

	if err = c.coupleUsecase.Destroy(ctxTimeout, id); err != nil {
		return flash.HandleError(ctx, c.session.Store, err, nil)
	}

	return flash.HandleSuccess(ctx, c.session.Store, "Couple successfully deleted", "/mimin/couple")
}

func (c *CoupleController) Publish(ctx *fiber.Ctx) error {
	ctxTimeout, cancel := context.WithTimeout(ctx.Context(), c.cfgApp.Timeout*time.Second)
	defer cancel()

	if err := c.coupleUsecase.Publish(ctxTimeout); err != nil {
		return flash.HandleError(ctx, c.session.Store, err, nil)
	}

	return flash.HandleSuccess(ctx, c.session.Store, "Couple successfully published", "/mimin/couple")
}
