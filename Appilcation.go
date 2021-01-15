package main

import (
	"fmt"
	"github.com/LILSunflower/DEKKER/routes"
	"github.com/LILSunflower/DEKKER/utils"
	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	routes.SetupAllRoutes(app)

	test := utils.GetDB()
	fmt.Printf("TEST DB - %v\n", test)

	app.Listen(":3000")
}