package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zakariawahyu/go-fiberavel/app/http/middleware"
	"github.com/zakariawahyu/go-fiberavel/internal/utils/flash"
)

type DashboardController struct {
	session *middleware.Session
}

func NewDashboardController(session *middleware.Session) *DashboardController {
	return &DashboardController{
		session: session,
	}
}

func (c *DashboardController) Index(ctx *fiber.Ctx) error {
	build := flash.NewMessage(c.session.Store).Build()
	return ctx.Render("backend/pages/dashboard/index", build.GetFlash(ctx))
}
