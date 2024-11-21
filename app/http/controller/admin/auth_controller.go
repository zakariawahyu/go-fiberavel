package controller

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/zakariawahyu/go-fiberavel/app/http/middleware"
	"github.com/zakariawahyu/go-fiberavel/app/http/request"
	"github.com/zakariawahyu/go-fiberavel/app/usecase/admin"
	"github.com/zakariawahyu/go-fiberavel/config"
	"github.com/zakariawahyu/go-fiberavel/internal/utils/flash"
	"time"
)

type AuthController struct {
	authUsecase usecase.AuthUsecase
	cfgApp      config.App
	session     *middleware.Session
	validator   *validator.Validate
}

func NewAuthController(authUsecase usecase.AuthUsecase, cfgApp config.App, session *middleware.Session, validator *validator.Validate) *AuthController {
	return &AuthController{
		authUsecase: authUsecase,
		cfgApp:      cfgApp,
		session:     session,
		validator:   validator,
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
	ctxTimeout, cancel := context.WithTimeout(ctx.Context(), c.cfgApp.Timeout*time.Second)
	defer cancel()

	var req request.LoginRequest

	if err := ctx.BodyParser(&req); err != nil {
		return flash.HandleError(ctx, c.session.Store, err, req)
	}

	if err := c.validator.Struct(req); err != nil {
		return flash.HandleValidationError(ctx, c.session.Store, err, req)
	}

	auth, err := c.authUsecase.Login(ctxTimeout, req)
	if err != nil {
		return flash.HandleError(ctx, c.session.Store, err, req)
	}

	if err := c.session.SetAuth(ctx, auth); err != nil {
		return flash.HandleError(ctx, c.session.Store, err, req)
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
