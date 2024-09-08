package controller

import (
	"github.com/Schaitanya535/tailored_be/m/middlewares"
	userService "github.com/Schaitanya535/tailored_be/m/services/user"
	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(router fiber.Router) {

	user := router.Group("/user")

	user.Post("/", middlewares.Protected, userService.CreateUser)
	user.Get("/all", userService.GetUsers)

}
