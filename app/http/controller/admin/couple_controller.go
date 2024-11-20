package controller

import (
	"context"
	"github.com/gofiber/fiber/v2"
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
}

func NewCoupleController(coupleRepo repository.CoupleRepository, cfgApp config.App, session *middleware.Session) *CoupleController {
	return &CoupleController{
		coupleRepo: coupleRepo,
		cfgApp:     cfgApp,
		session:    session,
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

	var params sqlc.CreateCoupleParams

	if err := ctx.BodyParser(&params); err != nil {
		return err
	}

	if err := request.CreateCoupleValidate(params); err != nil {
		build := flash.NewMessage(c.session.Store).WithErrorValidate(err).WithInput(params).Build()
		return build.SetFlash(ctx).RedirectBack("/")
	}

	couple, err := c.coupleRepo.Insert(context, params)
	if err != nil {

	}

	return ctx.JSON(couple)
}
