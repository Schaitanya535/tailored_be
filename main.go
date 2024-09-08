package main

import (
	"github.com/Schaitanya535/tailored_be/m/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {

	app := fiber.New()

	app.Use(cors.New(
		cors.Config{
			AllowOrigins: "https://gofiber.io, https://gofiber.net",
			AllowHeaders: "Origin, Content-Type, Accept",
		}))

	app.Use(logger.New(
		logger.Config{
			Format:     "${time} ${status} - ${latency} ${method} ${path}\n",
			TimeFormat: "02-Jan-2006",
			TimeZone:   "Asia/Kolkata",
		},
	))

	router.SetupRoutes(app)
	// database.Connect()

	app.Listen(":8080")

}
