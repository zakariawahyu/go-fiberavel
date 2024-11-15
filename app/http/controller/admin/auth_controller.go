package controller

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/zakariawahyu/go-fiberavel/app/http/middleware"
	"github.com/zakariawahyu/go-fiberavel/app/http/request"
	repository "github.com/zakariawahyu/go-fiberavel/app/repository/admin"
	"github.com/zakariawahyu/go-fiberavel/config"
	sqlc "github.com/zakariawahyu/go-fiberavel/internal/sqlc/generated"
	"github.com/zakariawahyu/go-fiberavel/internal/utils/flash"
	"github.com/zakariawahyu/go-fiberavel/internal/utils/helper"
	"time"
)

type AuthController struct {
	authRepo repository.AuthRepository
	cfgApp   config.App
	session  *middleware.Session
}

func NewAuthController(authRepo repository.AuthRepository, cfgApp config.App, session *middleware.Session) *AuthController {
	return &AuthController{
		authRepo: authRepo,
		cfgApp:   cfgApp,
		session:  session,
	}
}

func (c *AuthController) Index(ctx *fiber.Ctx) error {
	build := flash.NewMessage(c.session.Store).Build()

	return ctx.Render("backend/pages/auth/index", build.GetFlash(ctx))
}

func (c *AuthController) Login(ctx *fiber.Ctx) error {
	context, cancel := context.WithTimeout(ctx.Context(), c.cfgApp.Timeout*time.Second)
	defer cancel()

	var auth sqlc.LoginRow

	if err := ctx.BodyParser(&auth); err != nil {
		return err
	}

	if err := request.LoginValidate(auth); err != nil {
		build := flash.NewMessage(c.session.Store).WithErrorValidate(err).WithInput(auth).Build()
		return build.SetFlash(ctx).RedirectBack("/")
	}

	result, err := c.authRepo.Login(context, auth.Username)
	if err != nil {
		build := flash.NewMessage(c.session.Store).WithError(err).WithInput(auth).Build()
		return build.SetFlash(ctx).RedirectBack("/")
	}

	if err := helper.ComparePassword(result.Password, auth.Password); err != nil {
		build := flash.NewMessage(c.session.Store).WithError(middleware.ErrPasswordNotMatch).WithInput(auth).Build()
		return build.SetFlash(ctx).RedirectBack("/")
	}

	return ctx.JSON(result)
}