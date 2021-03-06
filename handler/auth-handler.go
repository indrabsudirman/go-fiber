package handler

import (
	"go-fiber/database"
	"go-fiber/model/entity"
	"go-fiber/model/request"
	"go-fiber/utils"
	"log"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func LoginHandler(ctx *fiber.Ctx) error {

	loginRequest := new(request.LoginRequest)
	if err := ctx.BodyParser(loginRequest); err != nil {
		return err
	}

	log.Println(loginRequest)

	//Validation Request
	validate := validator.New()
	errValidate := validate.Struct(loginRequest)
	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error":   errValidate.Error(),
		})
	}

	// Check available user

	var user entity.User

	err := database.DB.First(&user, "email = ?", loginRequest.Email).Error
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "wrong credential",
		})
	}

	//Check validation password
	isValid := utils.CheckPasswordHash(user.Password, loginRequest.Password)
	if !isValid {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "wrong credential",
		})
	}

	//Generate JWT
	claims := jwt.MapClaims{}
	claims["name"] = user.Name
	claims["email"] = user.Email
	claims["address"] = user.Address
	claims["phone"] = user.Phone
	claims["exp"] = time.Now().Add(time.Minute * 5).Unix()

	if user.Email == "indrasudirman@gmail.com" {
		claims["role"] = "admin"
	} else {
		claims["role"] = "user"
	}

	token, errGenerateToken := utils.GenerateToken(&claims)
	if errGenerateToken != nil {

		log.Println(errGenerateToken)
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "wrong credential",
		})
	}

	return ctx.JSON(fiber.Map{
		"token": token,
	})
}
