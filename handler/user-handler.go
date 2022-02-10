package handler

import (
	"go-fiber/database"
	"go-fiber/model/entity"
	"log"
	"os/user"

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

func UserHandlerCreate(ctx *fiber.Ctx) error  {
	user := new(request.UserCreateRequest)
	if err := ctx.BodyParser(user); err != nil {
		return err
	}

	newUser := entity.User {
		Name: user.Name,
		Email: user.Email,
		Address: user.Address,
		Phone: user.Phone,
	}

	errCreateUser := database.DB.Create(&newUser).Error
	if errCreateUser != {
		return ctx.Status(500).JSON(fiber.Map) {
			"message" : "failed to store data"
		}
	}