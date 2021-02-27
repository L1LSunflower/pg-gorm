package controller

import (
	"encoding/json"
	"fmt"
	"github.com/LILSunflower/DEKKER/models"
	"github.com/LILSunflower/DEKKER/utils"
	"github.com/gofiber/fiber/v2"
)

var (
	db = utils.GetDB()
)

func GetAllUsers(c *fiber.Ctx) error {
	listUsers := models.GetAll(db)
	result, _ := utils.ResponseWrapper(listUsers, 200, "")
	return c.Send(result)
}

func AddUser(c *fiber.Ctx) error {
	newUser := new(models.User)
	err := json.Unmarshal(c.Request().Body(), &newUser)
	if err != nil {
		return err
	}
	newUser.Create(db)
	result, _ := utils.ResponseWrapper(*newUser, 200, "User successfully created")
	// need to return access token
	return c.Send(result)
}

func UpdateUser(c *fiber.Ctx) error {
	user := new(models.User)
	err := json.Unmarshal(c.Request().Body(), &user)
	if err != nil {
		return err
	}
	fmt.Printf("User data: %v\n", user)
	user.Update(db)
	result, _ := utils.ResponseWrapper(*user, 200, "User successfully updated")
	return c.Send(result)
}

func GetUserById(c *fiber.Ctx) error {
	id := c.Params("id")
	user := new(models.User)
	user.GetByID(db, id)
	result, _ := utils.ResponseWrapper(user, 200, "")
	return c.Send(result)
}

func DeleteUser(c *fiber.Ctx) error {
	user := new(models.User)
	err := json.Unmarshal(c.Request().Body(), &user)
	if err != nil {
		return err
	}
	user.Delete(db)
	result, _ := utils.ResponseWrapper(user, 200, "User successfully deleted")
	return c.Send(result)
}