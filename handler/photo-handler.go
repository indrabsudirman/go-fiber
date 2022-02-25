package handler

import (
	"fmt"
	"go-fiber/model/request"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func PhotoHandlerCreate(ctx *fiber.Ctx) error {
	photo := new(request.PhotoCreateRequest)
	if err := ctx.BodyParser(photo); err != nil {
		return err
	}

	//Validation Request
	validate := validator.New()
	errValidate := validate.Struct(photo)
	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error":   errValidate.Error(),
		})
	}

	var filenameString string

	// Validation required image
	filenames := ctx.Locals("filenames")

	if filenames == nil {
		return ctx.Status(422).JSON(fiber.Map{
			"message": "image is required",
		})
	} else {
		filenameString = fmt.Sprintf("%v", filenames)
	}

	log.Println(filenameString)

	// newPhoto := entity.Photo{
	// 	Image:  filenames,
	// 	CategoryID: 1,
	// }

	// errCreatePhoto := database.DB.Create(&newPhoto).Error
	// if errCreatePhoto != nil {
	// 	log.Println("ada file yang gagal", errCreatePhoto)
	// 	// return ctx.Status(500).JSON(fiber.Map{
	// 	// 	"message": "failed to store data",
	// 	})
	// }

	return ctx.JSON(fiber.Map{
		"message": "success",
		// "data":    newPhoto,
	})
}
