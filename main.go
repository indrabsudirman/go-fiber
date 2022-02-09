package main

import (
	"go-fiber/database"
	"go-fiber/migration"
	"go-fiber/route"

	"github.com/gofiber/fiber/v2"
)

func main() {

	//Initial Database
	database.DatabaseInit()

	//Initial migration
	migration.RunMigration()

	app := fiber.New()

	//Initial Route
	route.RouteInit(app)

	app.Listen(":3000")
}
