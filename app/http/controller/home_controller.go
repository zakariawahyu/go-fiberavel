package controller

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/zakariawahyu/go-fiberavel/app/repository"
	"github.com/zakariawahyu/go-fiberavel/internal/utils/constants"
	"time"
)

type HomeController struct {
	homeRepo       repository.HomeRepository
	contextTimeout time.Duration
}

func NewHomeController(homeRepo repository.HomeRepository, contextTimeout time.Duration) *HomeController {
	return &HomeController{
		homeRepo:       homeRepo,
		contextTimeout: contextTimeout,
	}
}

func (ctrl *HomeController) Index(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), ctrl.contextTimeout)
	defer cancel()

	config, err := ctrl.homeRepo.Get(c, constants.KeyConfigs)
	if err != nil {
		return err
	}

	return ctx.JSON(fiber.Map{
		"data": config,
	})
}
