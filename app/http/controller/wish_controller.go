package controller

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/tidwall/gjson"
	"github.com/zakariawahyu/go-fiberavel/app/http/request"
	"github.com/zakariawahyu/go-fiberavel/app/usecase"
	"github.com/zakariawahyu/go-fiberavel/config"
	"time"
)

type WishController struct {
	wishUsecase usecase.WishUsecase
	cfgApp      config.App
	validator   *validator.Validate
}

func NewWishController(wishUsecase usecase.WishUsecase, cfgApp config.App, validator *validator.Validate) *WishController {
	return &WishController{
		wishUsecase: wishUsecase,
		cfgApp:      cfgApp,
		validator:   validator,
	}
}

func (c *WishController) GetAll(ctx *fiber.Ctx) error {
	result, err := c.wishUsecase.GetAll()
	if err != nil {
		return err
	}

	return ctx.Render("frontend/partials/wishes-data", fiber.Map{
		"wishes": gjson.Parse(string(result)).Value(),
	})
}

func (c *WishController) Store(ctx *fiber.Ctx) error {
	ctxTimeout, cancel := context.WithTimeout(ctx.Context(), c.cfgApp.Timeout*time.Second)
	defer cancel()

	var req request.CreateWishParams

	if err := ctx.BodyParser(&req); err != nil {
		return err
	}

	if err := c.validator.Struct(req); err != nil {
		return err
	}

	path := ctx.Path()
	result, err := c.wishUsecase.Store(ctxTimeout, req, path)
	if err != nil {
		return err
	}

	return ctx.JSON(result)
}
