package handler

import (
	"go-fiber/database"
	"go-fiber/model/entity"
	"go-fiber/model/request"
	"go-fiber/utils"
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

	// Validation required image
	filenames := ctx.Locals("filenames")
	log.Println("filename = :", filenames)

	if filenames == nil {
		return ctx.Status(422).JSON(fiber.Map{
			"message": "image is required",
		})
	} else {

		filenameData := filenames.([]string)
		for _, filename := range filenameData {
			newPhoto := entity.Photo{
				Image:      filename,
				CategoryID: photo.CategoryId,
			}

			errCreatePhoto := database.DB.Create(&newPhoto).Error
			if errCreatePhoto != nil {
				log.Println("some not saved properly", errCreatePhoto)
			}
		}
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
	})
}

func PhotoHandlerDelete(ctx *fiber.Ctx) error {
	photoId := ctx.Params("id")

	var photo entity.Photo

	//Available photo in DB
	err := database.DB.Debug().First(&photo, "id = ?", photoId).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "photo not found",
		})
	}

	//delete files on public/images
	errDeletePhoto := utils.HandleRemoveFile(photo.Image)
	if errDeletePhoto != nil {
		log.Println("failed to delete some photos")

	}

	//delete files on database
	errDelete := database.DB.Debug().Delete(&photo).Error
	if errDelete != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "photo was deleted",
	})

}
