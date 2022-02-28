package main

import (
	"go-fiber/database"
	"go-fiber/migration"
	"go-fiber/route"
	"log"
	"os"

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

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	errListen := app.Listen(":" + port)
	if errListen != nil {
		log.Println("Failed to listen Go fiber server")
		os.Exit(1)
	}
}
