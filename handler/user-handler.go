package handler

import (
	"go-fiber/database"
	"go-fiber/model/entity"
	"log"

	"github.com/gofiber/fiber/v2"
)

func UserHandlerGetAll(ctx *fiber.Ctx) error {

	var users []entity.User
	result := database.DB.Debug().Find(&users)
	if result.Error != nil {
		log.Println(result.Error)
	}

	return ctx.JSON(users)

}
