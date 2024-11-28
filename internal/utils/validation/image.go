package validation

import (
	"github.com/go-playground/validator/v10"
	"mime/multipart"
	"net/http"
	"strings"
)

func MimeType(fl validator.FieldLevel) bool {
	// Check if the field is a file
	image, ok := fl.Field().Interface().(multipart.FileHeader)
	if !ok {
		return false
	}

	// Open the file
	file, err := image.Open()
	if err != nil {
		return false
	}
	defer file.Close()

	// Read the first 512 bytes of the file to determine mime type
	buffer := make([]byte, 512)
	if _, err := file.Read(buffer); err != nil {
		return false
	}

	// Determine the mime type
	mimeType := http.DetectContentType(buffer)

	// Get the allowed mime types from the tag struct
	allowedMimeTypes := strings.Split(fl.Param(), "-")

	// Check if the mime type is allowed
	for _, allowedMimeType := range allowedMimeTypes {
		if mimeType == allowedMimeType {
			return true
		}
	}

	return false
}
