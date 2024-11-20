package controller

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/mashingan/smapping"
	"github.com/zakariawahyu/go-fiberavel/app/http/middleware"
	"github.com/zakariawahyu/go-fiberavel/app/http/request"
	repository "github.com/zakariawahyu/go-fiberavel/app/repository/admin"
	"github.com/zakariawahyu/go-fiberavel/config"
	sqlc "github.com/zakariawahyu/go-fiberavel/internal/sqlc/generated"
	"github.com/zakariawahyu/go-fiberavel/internal/utils/constants"
	"github.com/zakariawahyu/go-fiberavel/internal/utils/flash"
	"github.com/zakariawahyu/go-fiberavel/internal/utils/helper"
	"time"
)

type CoupleController struct {
	coupleRepo repository.CoupleRepository
	cfgApp     config.App
	session    *middleware.Session
	validator  *validator.Validate
}

func NewCoupleController(coupleRepo repository.CoupleRepository, cfgApp config.App, session *middleware.Session, validator *validator.Validate) *CoupleController {
	return &CoupleController{
		coupleRepo: coupleRepo,
		cfgApp:     cfgApp,
		session:    session,
		validator:  validator,
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
	context, cancel := context.WithTimeout(ctx.Context(), c.cfgApp.Timeout*time.Second)
	defer cancel()

	var req request.CreateCoupleRequest

	if err := ctx.BodyParser(&req); err != nil {
		return err
	}

	if err := c.validator.Struct(&req); err != nil {
		build := flash.NewMessage(c.session.Store).WithErrorValidate(err).Build()

		return build.SetFlash(ctx).RedirectBack("/")
	}

	var couple sqlc.CreateCoupleParams

	if err := smapping.FillStruct(&couple, smapping.MapFields(&req)); err != nil {
		return err
	}

	result, err := c.coupleRepo.Insert(context, couple)
	if err != nil {
		build := flash.NewMessage(c.session.Store).WithError(err).WithInput(couple).Build()
		return build.SetFlash(ctx).RedirectBack("/")
	}

	return ctx.JSON(result)
}
