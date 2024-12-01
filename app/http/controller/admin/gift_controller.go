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

type GiftController struct {
	giftUsecase usecase.GiftUsecase
	cfgApp      config.App
	session     *middleware.Session
	validator   *validator.Validate
}

func NewGiftController(giftUsecase usecase.GiftUsecase, cfgApp config.App, session *middleware.Session, validator *validator.Validate) *GiftController {
	return &GiftController{
		giftUsecase: giftUsecase,
		cfgApp:      cfgApp,
		session:     session,
		validator:   validator,
	}
}

func (c *GiftController) Index(ctx *fiber.Ctx) error {
	build := flash.NewMessage(c.session.Store).Build()
	return ctx.Render("backend/pages/gift/index", build.GetFlash(ctx))
}

func (c *GiftController) Datatables(ctx *fiber.Ctx) error {
	params, err := datatables.ParseDataTableParams(ctx)
	if err != nil {
		return err
	}

	gifts, err := c.giftUsecase.Datatables(ctx.Context(), params)
	if err != nil {
		return err
	}

	return ctx.JSON(gifts)
}

func (c *GiftController) Create(ctx *fiber.Ctx) error {
	build := flash.NewMessage(c.session.Store).Build()
	banks := constants.Banks

	return ctx.Render("backend/pages/gift/create", helper.Compact(fiber.Map{
		"banks": banks,
	}, build.GetFlash(ctx)))
}

func (c *GiftController) Store(ctx *fiber.Ctx) error {
	ctxTimeout, cancel := context.WithTimeout(ctx.Context(), c.cfgApp.Timeout*time.Second)
	defer cancel()

	var req request.CreateGiftRequest

	if err := ctx.BodyParser(&req); err != nil {
		return flash.HandleError(ctx, c.session.Store, err, req)
	}

	if err := c.validator.Struct(&req); err != nil {
		return flash.HandleValidationError(ctx, c.session.Store, err, req)
	}

	_, err := c.giftUsecase.Store(ctxTimeout, req)
	if err != nil {
		return flash.HandleError(ctx, c.session.Store, err, req)
	}

	return flash.HandleSuccess(ctx, c.session.Store, "Gift created successfully", "/mimin/gift")
}

func (c *GiftController) Show(ctx *fiber.Ctx) error {
	ctxTimeout, cancel := context.WithTimeout(ctx.Context(), c.cfgApp.Timeout*time.Second)
	defer cancel()

	id, err := helper.StrToInt64(ctx.Params("id"))
	if err != nil {
		return flash.HandleError(ctx, c.session.Store, err, nil)
	}

	gift, err := c.giftUsecase.FindById(ctxTimeout, id)
	if err != nil {
		return flash.HandleError(ctx, c.session.Store, err, nil)
	}

	return ctx.Render("backend/pages/gift/show", fiber.Map{
		"gift": gift,
	})
}

func (c *GiftController) Edit(ctx *fiber.Ctx) error {
	ctxTimeout, cancel := context.WithTimeout(ctx.Context(), c.cfgApp.Timeout*time.Second)
	defer cancel()

	id, err := helper.StrToInt64(ctx.Params("id"))
	if err != nil {
		return flash.HandleError(ctx, c.session.Store, err, nil)
	}

	gift, err := c.giftUsecase.FindById(ctxTimeout, id)
	if err != nil {
		return flash.HandleError(ctx, c.session.Store, err, nil)
	}

	build := flash.NewMessage(c.session.Store).Build()
	banks := constants.Banks

	return ctx.Render("backend/pages/gift/edit", helper.Compact(fiber.Map{
		"gift":  gift,
		"banks": banks,
	}, build.GetFlash(ctx)))
}

func (c *GiftController) Update(ctx *fiber.Ctx) error {
	ctxTimeout, cancel := context.WithTimeout(ctx.Context(), c.cfgApp.Timeout*time.Second)
	defer cancel()

	if ctx.FormValue("_method") == "PUT" {
		ctx.Method("PUT")
	}

	var req request.UpdateGiftRequest

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

	if err := c.giftUsecase.Update(ctxTimeout, req); err != nil {
		return flash.HandleError(ctx, c.session.Store, err, req)
	}

	return flash.HandleSuccess(ctx, c.session.Store, "Gift updated successfully", "/mimin/gift")
}

func (c *GiftController) Destroy(ctx *fiber.Ctx) error {
	ctxTimeout, cancel := context.WithTimeout(ctx.Context(), c.cfgApp.Timeout*time.Second)
	defer cancel()

	id, err := helper.StrToInt64(ctx.Params("id"))
	if err != nil {
		return flash.HandleError(ctx, c.session.Store, err, nil)
	}

	if err := c.giftUsecase.Destroy(ctxTimeout, id); err != nil {
		return flash.HandleError(ctx, c.session.Store, err, nil)
	}

	return flash.HandleSuccess(ctx, c.session.Store, "Gift deleted successfully", "/mimin/gift")
}

func (c *GiftController) Publish(ctx *fiber.Ctx) error {
	ctxTimeout, cancel := context.WithTimeout(ctx.Context(), c.cfgApp.Timeout*time.Second)
	defer cancel()

	if err := c.giftUsecase.Publish(ctxTimeout); err != nil {
		return flash.HandleError(ctx, c.session.Store, err, nil)
	}

	return flash.HandleSuccess(ctx, c.session.Store, "Gift published successfully", "/mimin/gift")
}
