package controller

import (
	"context"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/zakariawahyu/go-fiberavel/app/http/middleware"
	"github.com/zakariawahyu/go-fiberavel/app/http/request"
	usecase "github.com/zakariawahyu/go-fiberavel/app/usecase/admin"
	"github.com/zakariawahyu/go-fiberavel/config"
	"github.com/zakariawahyu/go-fiberavel/internal/utils/constants"
	"github.com/zakariawahyu/go-fiberavel/internal/utils/flash"
	"github.com/zakariawahyu/go-fiberavel/internal/utils/helper"
	"time"
)

type ConfigController struct {
	configUsecase usecase.ConfigUsecase
	cfgApp        config.App
	session       *middleware.Session
	validator     *validator.Validate
}

func NewConfigController(configUsecase usecase.ConfigUsecase, cfgApp config.App, session *middleware.Session, validator *validator.Validate) *ConfigController {
	return &ConfigController{
		configUsecase: configUsecase,
		cfgApp:        cfgApp,
		session:       session,
		validator:     validator,
	}
}

func (c *ConfigController) Index(ctx *fiber.Ctx) error {
	type_ := ctx.Params("type")

	configs := constants.Configuration
	if !helper.InArray(configs, type_) {
		return flash.HandleError(ctx, c.session.Store, errors.New("Configuration not found"), nil)
	}

	config, err := c.configUsecase.FindByType(ctx.Context(), type_)
	if err != nil {
		return flash.HandleError(ctx, c.session.Store, err, nil)
	}

	build := flash.NewMessage(c.session.Store).Build()

	return ctx.Render("backend/pages/config/"+type_, helper.Compact(fiber.Map{
		"config": config,
	}, build.GetFlash(ctx)))
}

func (c *ConfigController) StoreCover(ctx *fiber.Ctx) error {
	ctxTimeout, cancel := context.WithTimeout(ctx.Context(), c.cfgApp.Timeout*time.Second)
	defer cancel()

	var req request.ConfigCoverRequest

	if err := ctx.BodyParser(&req); err != nil {
		return flash.HandleError(ctx, c.session.Store, err, req)
	}

	req.Type = "cover"
	if err := c.validator.Struct(req); err != nil {
		return flash.HandleValidationError(ctx, c.session.Store, err, req)
	}

	if err := c.configUsecase.StoreCover(ctxTimeout, req); err != nil {
		return flash.HandleError(ctx, c.session.Store, err, req)
	}

	return flash.HandleSuccess(ctx, c.session.Store, "Cover has been updated", "/mimin/config/cover")
}

func (c *ConfigController) StoreVenue(ctx *fiber.Ctx) error {
	ctxTimeout, cancel := context.WithTimeout(ctx.Context(), c.cfgApp.Timeout*time.Second)
	defer cancel()

	var req request.ConfigVenueRequest

	if err := ctx.BodyParser(&req); err != nil {
		return flash.HandleError(ctx, c.session.Store, err, req)
	}

	req.Type = "venue"
	if err := c.validator.Struct(req); err != nil {
		return flash.HandleValidationError(ctx, c.session.Store, err, req)
	}

	if err := c.configUsecase.StoreVenue(ctxTimeout, req); err != nil {
		return flash.HandleError(ctx, c.session.Store, err, req)
	}

	return flash.HandleSuccess(ctx, c.session.Store, "Venue has been updated", "/mimin/config/venue")
}

func (c *ConfigController) StoreGift(ctx *fiber.Ctx) error {
	ctxTimeout, cancel := context.WithTimeout(ctx.Context(), c.cfgApp.Timeout*time.Second)
	defer cancel()

	var req request.ConfigGiftRequest

	if err := ctx.BodyParser(&req); err != nil {
		return flash.HandleError(ctx, c.session.Store, err, req)
	}

	req.Type = "gift"
	if err := c.validator.Struct(req); err != nil {
		return flash.HandleValidationError(ctx, c.session.Store, err, req)
	}

	if err := c.configUsecase.StoreGift(ctxTimeout, req); err != nil {
		return flash.HandleError(ctx, c.session.Store, err, req)
	}

	return flash.HandleSuccess(ctx, c.session.Store, "Gift has been updated", "/mimin/config/gift")
}

func (c *ConfigController) StoreWish(ctx *fiber.Ctx) error {
	ctxTimeout, cancel := context.WithTimeout(ctx.Context(), c.cfgApp.Timeout*time.Second)
	defer cancel()

	var req request.ConfiWishRequest

	if err := ctx.BodyParser(&req); err != nil {
		return flash.HandleError(ctx, c.session.Store, err, req)
	}

	req.Type = "wish"
	if err := c.validator.Struct(req); err != nil {
		return flash.HandleValidationError(ctx, c.session.Store, err, req)
	}

	if err := c.configUsecase.StoreWish(ctxTimeout, req); err != nil {
		return flash.HandleError(ctx, c.session.Store, err, req)
	}

	return flash.HandleSuccess(ctx, c.session.Store, "Wish has been updated", "/mimin/config/wish")
}
