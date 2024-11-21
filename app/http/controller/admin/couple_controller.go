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

	if err := c.validator.Struct(&req); err != nil {
		return flash.HandleValidationError(ctx, c.session.Store, err, req)
	}

	couple, err := c.coupleUsecase.Store(ctxTimeout, req)
	if err != nil {
		return flash.HandleError(ctx, c.session.Store, err, req)
	}

	return ctx.JSON(couple)
}
