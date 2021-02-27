package main

import (
	"github.com/LILSunflower/DEKKER/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()
	routes.SetupAllRoutes(app)

	//utils.MigrateDB()
	//test := utils.GetDB()
	//fmt.Printf("TEST DB - %v\n", test)

	app.Listen(":3000")
}