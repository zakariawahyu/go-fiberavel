package controller

import (
	"context"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/redis/rueidis"
	"github.com/tidwall/gjson"
	"github.com/zakariawahyu/go-fiberavel/app/repository"
	"github.com/zakariawahyu/go-fiberavel/config"
	sqlc "github.com/zakariawahyu/go-fiberavel/internal/sqlc/generated"
	"github.com/zakariawahyu/go-fiberavel/internal/utils/constants"
	"time"
)

type WishController struct {
	wishRepo repository.WishesRepository
	cfgApp   config.App
}

func NewWishController(wishRepo repository.WishesRepository, cfgApp config.App) *WishController {
	return &WishController{
		wishRepo: wishRepo,
		cfgApp:   cfgApp,
	}
}

func (ctrl *WishController) CreateWish(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), ctrl.cfgApp.Timeout*time.Second)
	defer cancel()

	var request sqlc.CreateWishParams

	if err := ctx.BodyParser(&request); err != nil {
		return err
	}

	result, err := ctrl.wishRepo.CreateWish(c, request)
	if err != nil {
		return err
	}

	wishes, err := ctrl.wishRepo.GetAllWishes(c)
	if err != nil {
		return err
	}

	wishesBytes, err := json.Marshal(wishes)
	if err != nil {
		return err
	}

	err = ctrl.wishRepo.SetRedis(c, constants.KeyWishes, rueidis.BinaryString(wishesBytes))
	if err != nil {
		return err
	}

	return ctx.JSON(result)
}

func (ctrl *WishController) GetAllWishes(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), ctrl.cfgApp.Timeout*time.Second)
	defer cancel()

	result, err := ctrl.wishRepo.GetRedis(c, constants.KeyWishes)
	if err != nil {
		return err
	}

	return ctx.Render("frontend/partials/wishes-data", fiber.Map{
		"wishes": gjson.Parse(result).Value(),
	})
}
