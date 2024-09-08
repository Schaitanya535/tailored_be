package router

import (
	controller "github.com/Schaitanya535/tailored_be/m/router/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	api := app.Group("/api")
	apiV1 := api.Group("/v1")

	controller.SetupUserRoutes(apiV1)
	controller.SetupProductRoutes(apiV1)

}
