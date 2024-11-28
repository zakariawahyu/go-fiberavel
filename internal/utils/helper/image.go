package helper

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gosimple/slug"
	"github.com/valyala/fasthttp"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
)

func GetImage(ctx *fiber.Ctx, key string) (*multipart.FileHeader, error) {
	image, err := ctx.FormFile(key)
	if err != nil && !errors.Is(err, fasthttp.ErrMissingFile) {
		return nil, err
	}

	return image, nil
}

func UploadImage(ctx *fiber.Ctx, image *multipart.FileHeader, caption string) (string, error) {
	ext := filepath.Ext(image.Filename)
	timeStamp := time.Now().Unix()
	fileName := slug.Make(caption)
	imageName := fmt.Sprintf("%s-%d%s", fileName, timeStamp, ext)

	uploadDir := "./public/images"
	if err := createDirIfNotExists(uploadDir); err != nil {
		return "", err
	}

	filePath := filepath.Join(uploadDir, imageName)

	if err := ctx.SaveFile(image, filePath); err != nil {
		return "", err
	}

	return imageName, nil
}

func createDirIfNotExists(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return os.MkdirAll(dir, 0755)
	}
	return nil
}
