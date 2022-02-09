package main

import (
	"go-fiber/database"
	"go-fiber/route"

	"github.com/gofiber/fiber/v2"
)

func main() {

	//Initial Database
	database.DatabaseInit()

	app := fiber.New()

	//Initial Route
	route.RouteInit(app)

	app.Listen(":3000")
}
