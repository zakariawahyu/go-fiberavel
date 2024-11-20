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
	"github.com/zakariawahyu/go-fiberavel/internal/utils/flash"
	"github.com/zakariawahyu/go-fiberavel/internal/utils/helper"
	"time"
)

type AuthController struct {
	authRepo  repository.AuthRepository
	cfgApp    config.App
	session   *middleware.Session
	validator *validator.Validate
}

func NewAuthController(authRepo repository.AuthRepository, cfgApp config.App, session *middleware.Session, validator *validator.Validate) *AuthController {
	return &AuthController{
		authRepo:  authRepo,
		cfgApp:    cfgApp,
		session:   session,
		validator: validator,
	}
}

func (c *AuthController) Index(ctx *fiber.Ctx) error {
	sess, err := c.session.Store.Get(ctx)
	if err != nil {
		return err
	}

	if sess.Get(middleware.KeyAuthSession) != nil {
		return ctx.Redirect("/mimin/dashboard")
	}

	build := flash.NewMessage(c.session.Store).Build()

	return ctx.Render("backend/pages/auth/index", build.GetFlash(ctx))
}

func (c *AuthController) Login(ctx *fiber.Ctx) error {
	context, cancel := context.WithTimeout(ctx.Context(), c.cfgApp.Timeout*time.Second)
	defer cancel()

	var req request.LoginRequest

	if err := ctx.BodyParser(&req); err != nil {
		return err
	}

	if err := c.validator.Struct(req); err != nil {
		build := flash.NewMessage(c.session.Store).WithErrorValidate(err).WithInput(req).Build()
		return build.SetFlash(ctx).RedirectBack("/")
	}

	var auth sqlc.LoginRow

	if err := smapping.FillStruct(&auth, smapping.MapFields(req)); err != nil {
		return err
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

	if err = c.session.SetAuth(ctx, result.Username); err != nil {
		return err
	}

	return ctx.Redirect("/mimin/dashboard")
}

func (c *AuthController) Logout(ctx *fiber.Ctx) error {
	sess, err := c.session.Store.Get(ctx)
	if err != nil {
		return err
	}

	if err = sess.Regenerate(); err != nil {
		return err
	}

	return ctx.Redirect("/auth/mimin")
}

func (c *AuthController) Unauthorized(ctx *fiber.Ctx) error {
	return ctx.Render("backend/pages/auth/unauthorized", nil)
}
