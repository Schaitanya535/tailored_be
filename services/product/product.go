package productService

import (
	"github.com/gofiber/fiber/v2"
)

func CreateProduct(c *fiber.Ctx) error {
	return c.SendString("Create Product")
}

func GetProducts(c *fiber.Ctx) error {
	return c.SendString("Get Products")
}
