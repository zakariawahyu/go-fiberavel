package controller

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/zakariawahyu/go-fiberavel/app/http/middleware"
	"github.com/zakariawahyu/go-fiberavel/app/http/request"
	repository "github.com/zakariawahyu/go-fiberavel/app/repository/admin"
	"github.com/zakariawahyu/go-fiberavel/config"
	sqlc "github.com/zakariawahyu/go-fiberavel/internal/sqlc/generated"
	"github.com/zakariawahyu/go-fiberavel/internal/utils/helper"
	"time"
)

type AuthController struct {
	authRepo repository.AuthRepository
	cfgApp   config.App
	session  *middleware.FlashStore
}

func NewAuthController(authRepo repository.AuthRepository, cfgApp config.App, session *middleware.FlashStore) *AuthController {
	return &AuthController{
		authRepo: authRepo,
		cfgApp:   cfgApp,
		session:  session,
	}
}

func (c *AuthController) Index(ctx *fiber.Ctx) error {
	return ctx.Render("backend/pages/auth/index", c.session.GetFlash(ctx))
}

func (c *AuthController) Login(ctx *fiber.Ctx) error {
	context, cancel := context.WithTimeout(ctx.Context(), c.cfgApp.Timeout*time.Second)
	defer cancel()

	var auth sqlc.LoginRow

	if err := ctx.BodyParser(&auth); err != nil {
		return err
	}

	if err := request.LoginValidate(auth); err != nil {
		fm := fiber.Map{
			"errValidate": middleware.GetError(err),
		}
		return c.session.SetFlash(ctx, fm).RedirectBack("/")
	}

	result, err := c.authRepo.Login(context, auth.Username)
	if err != nil {
		fm := fiber.Map{
			"errors": err.Error(),
		}
		return c.session.SetFlash(ctx, fm).RedirectBack("/")
	}

	if err := helper.ComparePassword(result.Password, auth.Password); err != nil {
		fm := fiber.Map{
			"errors": middleware.ErrPasswordNotMatch,
		}
		return c.session.SetFlash(ctx, fm).RedirectBack("/")
	}

	return ctx.JSON(result)
}
