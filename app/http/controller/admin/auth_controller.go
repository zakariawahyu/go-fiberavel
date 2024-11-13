package controller

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/zakariawahyu/go-fiberavel/app/http/request"
	repository "github.com/zakariawahyu/go-fiberavel/app/repository/admin"
	"github.com/zakariawahyu/go-fiberavel/config"
	sqlc "github.com/zakariawahyu/go-fiberavel/internal/sqlc/generated"
	"time"
)

type AuthController struct {
	authRepo repository.AuthRepository
	cfgApp   config.App
}

func NewAuthController(authRepo repository.AuthRepository, cfgApp config.App) *AuthController {
	return &AuthController{
		authRepo: authRepo,
		cfgApp:   cfgApp,
	}
}

func (c *AuthController) Index(ctx *fiber.Ctx) error {
	return ctx.Render("backend/pages/auth/index", fiber.Map{})
}

func (c *AuthController) Login(ctx *fiber.Ctx) error {
	context, cancel := context.WithTimeout(ctx.Context(), c.cfgApp.Timeout*time.Second)
	defer cancel()

	var auth sqlc.LoginRow

	if err := ctx.BodyParser(&auth); err != nil {
		return err
	}

	if err := request.LoginValidate(auth); err != nil {
		return err
	}

	result, err := c.authRepo.Login(context, auth.Username)
	if err != nil {
		return err
	}

	return ctx.JSON(result)
}
