package userService

import (
	"github.com/gofiber/fiber/v2"
)

var mapUsers = map[int]string{
	1: "User 1",
	2: "User 2",
	3: "User 3",
}

func GetUsers(c *fiber.Ctx) error {
	return c.JSON(mapUsers)
}

func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	return c.SendString("Get User " + id)
}

func CreateUser(c *fiber.Ctx) error {
	len := len(mapUsers)
	mapUsers[len+1] = "User " + string(rune(len+1))
	return c.JSON(mapUsers)
}
