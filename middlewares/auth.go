package middlewares

import (
	"github.com/gofiber/fiber/v2"
)

func Protected(c *fiber.Ctx) error {
	if c.Get("Authorization") == "Bearer mytoken" {
		return c.Next()
	}
	return c.SendString("Protected")
}
