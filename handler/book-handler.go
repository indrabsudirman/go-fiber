package handler

import (
	"fmt"
	"go-fiber/config"
	"go-fiber/database"
	"go-fiber/model/entity"
	"go-fiber/model/request"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func BookHandlerCreate(ctx *fiber.Ctx) error {
	book := new(request.BookCreateRequest)
	if err := ctx.BodyParser(book); err != nil {
		return err
	}

	//Validation Request
	validate := validator.New()
	errValidate := validate.Struct(book)
	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error":   errValidate.Error(),
		})
	}

	//Handle File
	file, errFile := ctx.FormFile("cover")
	if errFile != nil {
		log.Println("Error file : ", errFile)
	}

	//Check if file already
	var fileName string
	if file != nil {
		fileName = file.Filename

		errSaveFile := ctx.SaveFile(file, fmt.Sprintf(config.ProjectRootPath+"/public/images/%s", fileName))
		if errSaveFile != nil {
			log.Println("Failed to store file : ", errSaveFile)
		}
	} else {
		log.Println("nothing file to be upload")
	}

	newBook := entity.Book{
		Title:  book.Title,
		Author: book.Author,
		Cover:  fileName,
	}

	errCreateBook := database.DB.Create(&newBook).Error
	if errCreateBook != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "failed to store data",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    newBook,
	})

}
