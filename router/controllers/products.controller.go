package controller

import (
	"github.com/Schaitanya535/tailored_be/m/middlewares"
	productService "github.com/Schaitanya535/tailored_be/m/services/product"
	"github.com/gofiber/fiber/v2"
)

func SetupProductRoutes(router fiber.Router) {

	product := router.Group("/product")

	product.Post("/", middlewares.Protected, productService.CreateProduct)

	product.Get("/all", productService.GetProducts)

}
