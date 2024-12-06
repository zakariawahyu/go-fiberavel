package helper

import "github.com/gofiber/fiber/v2"

func Compact(base fiber.Map, extra interface{}) fiber.Map {
	if extra != nil {
		for key, value := range extra.(map[string]interface{}) {
			base[key] = value
		}
	}

	return base
}

func InArray(data map[string]string, key string) bool {
	for k, _ := range data {
		if k == key {
			return true
		}
	}

	return false
}
