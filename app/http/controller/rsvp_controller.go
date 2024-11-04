package controller

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/zakariawahyu/go-fiberavel/app/repository"
	"github.com/zakariawahyu/go-fiberavel/config"
	sqlc "github.com/zakariawahyu/go-fiberavel/internal/sqlc/generated"
	"time"
)

type RsvpController struct {
	rsvpRepo repository.RsvpRepository
	cfgApp   config.App
}

func NewRsvpController(rsvpRepo repository.RsvpRepository, cfgApp config.App) *RsvpController {
	return &RsvpController{
		rsvpRepo: rsvpRepo,
		cfgApp:   cfgApp,
	}
}

func (ctrl *RsvpController) CreateRsvp(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), ctrl.cfgApp.Timeout*time.Second)
	defer cancel()

	var request sqlc.CreateRsvpParams

	if err := ctx.BodyParser(&request); err != nil {
		return err
	}

	result, err := ctrl.rsvpRepo.CreateRsvp(c, request)
	if err != nil {
		return err
	}

	return ctx.JSON(result)
}
