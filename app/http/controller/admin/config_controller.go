package controller

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/zakariawahyu/go-fiberavel/app/http/middleware"
	"github.com/zakariawahyu/go-fiberavel/app/http/request"
	usecase "github.com/zakariawahyu/go-fiberavel/app/usecase/admin"
	"github.com/zakariawahyu/go-fiberavel/config"
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
	config, err := c.configUsecase.FindByType(ctx.Context(), "cover")
	if err != nil {
		return flash.HandleError(ctx, c.session.Store, err, nil)
	}

	build := flash.NewMessage(c.session.Store).Build()

	return ctx.Render("backend/pages/config/cover", helper.Compact(fiber.Map{
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
