package routes

import (
	"github.com/LILSunflower/DEKKER/controller"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func UserRoutes(app *fiber.App){

	api := app.Group("/api", logger.New())

	api.Get("/users", controller.GetAllUsers)
	api.Post("/users", controller.AddUser)
	api.Put("/users", controller.UpdateUser)
	api.Get("/users/:id", controller.GetUserByID)
	api.Delete("/users/:id", controller.DeleteUser)

}

