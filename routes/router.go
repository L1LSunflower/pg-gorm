package routes

import (
	"github.com/gofiber/fiber/v2"
)

func SetupAllRoutes(app *fiber.App){
	UserRoutes(app)
}
