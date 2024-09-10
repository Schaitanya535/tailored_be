package userService

import (
	"strconv"

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
	num, err := strconv.Atoi(id)

	if err != nil {
		return c.SendString("Invalid ID")
	}

	if num > len(mapUsers) {
		return c.SendString("User not found")
	}

	user := mapUsers[num]
	return c.SendString("Get User " + user)
}

func CreateUser(c *fiber.Ctx) error {
	len := len(mapUsers)
	mapUsers[len+1] = "User " + string(rune(len+1))
	return c.JSON(mapUsers)
}
