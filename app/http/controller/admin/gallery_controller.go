package controller

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/zakariawahyu/go-fiberavel/app/http/middleware"
	"github.com/zakariawahyu/go-fiberavel/app/http/request"
	usecase "github.com/zakariawahyu/go-fiberavel/app/usecase/admin"
	"github.com/zakariawahyu/go-fiberavel/config"
	"github.com/zakariawahyu/go-fiberavel/internal/utils/datatables"
	"github.com/zakariawahyu/go-fiberavel/internal/utils/flash"
	"github.com/zakariawahyu/go-fiberavel/internal/utils/helper"
	"time"
)

type GalleryController struct {
	galleryUsecase usecase.GalleryUsecase
	cfgApp         config.App
	session        *middleware.Session
	validator      *validator.Validate
}

func NewGalleryController(galleryUsecase usecase.GalleryUsecase, cfgApp config.App, session *middleware.Session, validator *validator.Validate) *GalleryController {
	return &GalleryController{
		galleryUsecase: galleryUsecase,
		cfgApp:         cfgApp,
		session:        session,
		validator:      validator,
	}
}

func (c *GalleryController) Index(ctx *fiber.Ctx) error {
	build := flash.NewMessage(c.session.Store).Build()
	return ctx.Render("backend/pages/gallery/index", build.GetFlash(ctx))
}

func (c *GalleryController) Datatables(ctx *fiber.Ctx) error {
	params, err := datatables.ParseDataTableParams(ctx)
	if err != nil {
		return err
	}

	galleries, err := c.galleryUsecase.Datatables(ctx.Context(), params)
	if err != nil {
		return err
	}

	return ctx.JSON(galleries)
}

func (c *GalleryController) Create(ctx *fiber.Ctx) error {
	build := flash.NewMessage(c.session.Store).Build()
	return ctx.Render("backend/pages/gallery/create", build.GetFlash(ctx))
}

func (c *GalleryController) Store(ctx *fiber.Ctx) error {
	ctxTimeout, cancel := context.WithTimeout(ctx.Context(), c.cfgApp.Timeout*time.Second)
	defer cancel()

	var req request.CreateGalleryRequest

	if err := ctx.BodyParser(&req); err != nil {
		return flash.HandleError(ctx, c.session.Store, err, req)
	}

	image, err := helper.GetImage(ctx, "image")
	if err != nil {
		return flash.HandleError(ctx, c.session.Store, err, req)
	}

	req.File = image
	if err := c.validator.Struct(&req); err != nil {
		return flash.HandleValidationError(ctx, c.session.Store, err, req)
	}

	imageName, err := helper.UploadImage(ctx, req.File, req.ImageCaption)
	if err != nil {
		return flash.HandleError(ctx, c.session.Store, err, req)
	}

	req.Image = imageName
	_, err = c.galleryUsecase.Store(ctxTimeout, req)
	if err != nil {
		return flash.HandleError(ctx, c.session.Store, err, req)
	}

	return flash.HandleSuccess(ctx, c.session.Store, "Gallery created successfully", "/mimin/gallery")
}

func (c *GalleryController) Show(ctx *fiber.Ctx) error {
	ctxTimeout, cancel := context.WithTimeout(ctx.Context(), c.cfgApp.Timeout*time.Second)
	defer cancel()

	id, err := helper.StrToInt64(ctx.Params("id"))
	if err != nil {
		return flash.HandleError(ctx, c.session.Store, err, nil)
	}

	gallery, err := c.galleryUsecase.FindById(ctxTimeout, id)
	if err != nil {
		return flash.HandleError(ctx, c.session.Store, err, nil)
	}

	return ctx.Render("backend/pages/gallery/show", fiber.Map{
		"gallery": gallery,
	})
}

func (c *GalleryController) Edit(ctx *fiber.Ctx) error {
	ctxTimeout, cancel := context.WithTimeout(ctx.Context(), c.cfgApp.Timeout*time.Second)
	defer cancel()

	id, err := helper.StrToInt64(ctx.Params("id"))
	if err != nil {
		return flash.HandleError(ctx, c.session.Store, err, nil)
	}

	gallery, err := c.galleryUsecase.FindById(ctxTimeout, id)
	if err != nil {
		return flash.HandleError(ctx, c.session.Store, err, nil)
	}

	build := flash.NewMessage(c.session.Store).Build()
	return ctx.Render("backend/pages/gallery/edit", helper.Compact(fiber.Map{
		"gallery": gallery,
	}, build.GetFlash(ctx)))
}

func (c *GalleryController) Update(ctx *fiber.Ctx) error {
	ctxTimeout, cancel := context.WithTimeout(ctx.Context(), c.cfgApp.Timeout*time.Second)
	defer cancel()

	if ctx.FormValue("_method") == "PUT" {
		ctx.Method("PUT")
	}

	var req request.UpdateGalleryRequest

	id, err := helper.StrToInt64(ctx.Params("id"))
	if err != nil {
		return flash.HandleError(ctx, c.session.Store, err, nil)
	}

	req.ID = id
	if err := ctx.BodyParser(&req); err != nil {
		return flash.HandleError(ctx, c.session.Store, err, req)
	}

	image, err := helper.GetImage(ctx, "image")
	if err != nil {
		return flash.HandleError(ctx, c.session.Store, err, req)
	}

	req.File = image
	if err := c.validator.Struct(&req); err != nil {
		return flash.HandleValidationError(ctx, c.session.Store, err, req)
	}

	if req.File != nil {
		imageName, err := helper.UploadImage(ctx, req.File, req.ImageCaption)
		if err != nil {
			return flash.HandleError(ctx, c.session.Store, err, req)
		}

		req.Image = imageName
	}

	if err := c.galleryUsecase.Update(ctxTimeout, req); err != nil {
		return flash.HandleError(ctx, c.session.Store, err, req)
	}

	return flash.HandleSuccess(ctx, c.session.Store, "Gallery updated successfully", "/mimin/gallery")
}

func (c *GalleryController) Destroy(ctx *fiber.Ctx) error {
	ctxTimeout, cancel := context.WithTimeout(ctx.Context(), c.cfgApp.Timeout*time.Second)
	defer cancel()

	id, err := helper.StrToInt64(ctx.Params("id"))
	if err != nil {
		return flash.HandleError(ctx, c.session.Store, err, nil)
	}

	if err := c.galleryUsecase.Destroy(ctxTimeout, id); err != nil {
		return flash.HandleError(ctx, c.session.Store, err, nil)
	}

	return flash.HandleSuccess(ctx, c.session.Store, "Gallery deleted successfully", "/mimin/gallery")
}

func (c *GalleryController) Publish(ctx *fiber.Ctx) error {
	ctxTimeout, cancel := context.WithTimeout(ctx.Context(), c.cfgApp.Timeout*time.Second)
	defer cancel()

	if err := c.galleryUsecase.Publish(ctxTimeout); err != nil {
		return flash.HandleError(ctx, c.session.Store, err, nil)
	}

	return flash.HandleSuccess(ctx, c.session.Store, "Gallery published successfully", "/mimin/gallery")
}
