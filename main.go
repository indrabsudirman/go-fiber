package main

import (
	"go-fiber/route"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	//Initial Route
	route.RouteInit(app)

	app.Listen(":3000")
}
