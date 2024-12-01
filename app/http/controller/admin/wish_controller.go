package controller

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/zakariawahyu/go-fiberavel/app/http/middleware"
	usecase "github.com/zakariawahyu/go-fiberavel/app/usecase/admin"
	"github.com/zakariawahyu/go-fiberavel/config"
	"github.com/zakariawahyu/go-fiberavel/internal/utils/datatables"
	"github.com/zakariawahyu/go-fiberavel/internal/utils/flash"
	"github.com/zakariawahyu/go-fiberavel/internal/utils/helper"
	"time"
)

type WishController struct {
	wishUsecase usecase.WishUsecase
	cfgApp      config.App
	session     *middleware.Session
	validator   *validator.Validate
}

func NewWishController(wishUsecase usecase.WishUsecase, cfgApp config.App, session *middleware.Session, validator *validator.Validate) *WishController {
	return &WishController{
		wishUsecase: wishUsecase,
		cfgApp:      cfgApp,
		session:     session,
		validator:   validator,
	}
}

func (c *WishController) Index(ctx *fiber.Ctx) error {
	build := flash.NewMessage(c.session.Store).Build()
	return ctx.Render("backend/pages/wish/index", build.GetFlash(ctx))
}

func (c *WishController) Datatables(ctx *fiber.Ctx) error {
	params, err := datatables.ParseDataTableParams(ctx)
	if err != nil {
		return err
	}

	venues, err := c.wishUsecase.Datatables(ctx.Context(), params)
	if err != nil {
		return err
	}

	return ctx.JSON(venues)
}

func (c *WishController) Destroy(ctx *fiber.Ctx) error {
	ctxTimeout, cancel := context.WithTimeout(ctx.Context(), c.cfgApp.Timeout*time.Second)
	defer cancel()

	id, err := helper.StrToInt64(ctx.Params("id"))
	if err != nil {
		return flash.HandleError(ctx, c.session.Store, err, nil)
	}

	if err := c.wishUsecase.Destroy(ctxTimeout, id); err != nil {
		return flash.HandleError(ctx, c.session.Store, err, nil)
	}

	return flash.HandleSuccess(ctx, c.session.Store, "Wish deleted successfully", "/mimin/wish")
}

func (c *WishController) Publish(ctx *fiber.Ctx) error {
	ctxTimeout, cancel := context.WithTimeout(ctx.Context(), c.cfgApp.Timeout*time.Second)
	defer cancel()

	if err := c.wishUsecase.Publish(ctxTimeout); err != nil {
		return flash.HandleError(ctx, c.session.Store, err, nil)
	}

	return flash.HandleSuccess(ctx, c.session.Store, "Wish published successfully", "/mimin/wish")
}
