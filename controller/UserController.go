package controller

import (
	"github.com/gofiber/fiber/v2"
)


func GetAllUsers(c *fiber.Ctx) error {
	//models.User{}
	//result := "Method: GET, All users,\n Your Request: " + c.Request().Body()
	return c.SendString(string(c.Request().Body()))
}

func AddUser(c *fiber.Ctx) error {
	return c.SendString("Method: POST, Add users")
}

func UpdateUser(c *fiber.Ctx) error {
	return c.SendString("Method PUT, Update users")
}

func GetUserByID(c *fiber.Ctx) error {
	return c.SendString("Method: GET, Get user by id")
}

func DeleteUser(c *fiber.Ctx) error {
	return c.SendString("Method: DELETE, Delete user by id")
}