package controller

import "github.com/gofiber/fiber/v2"

type DashboardController struct {
}

func NewDashboardController() *DashboardController {
	return &DashboardController{}
}

func (ctrl *DashboardController) Index(ctx *fiber.Ctx) error {
	return ctx.Render("backend/pages/dashboard/index", fiber.Map{})
}
