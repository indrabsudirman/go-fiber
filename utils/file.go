package utils

import (
	"fmt"
	"go-fiber/config"
	"log"

	"github.com/gofiber/fiber/v2"
)

func HandleSingleFile(ctx *fiber.Ctx) error {
	//Handle File
	file, errFile := ctx.FormFile("cover")
	if errFile != nil {
		log.Println("Error file : ", errFile)
	}

	//Check if file already
	var fileName *string
	if file != nil {
		fileName = &file.Filename

		errSaveFile := ctx.SaveFile(file, fmt.Sprintf(config.ProjectRootPath+"/public/images/%s", *fileName))
		if errSaveFile != nil {
			log.Println("Failed to store file : ", errSaveFile)
		}
	} else {
		log.Println("nothing file to be upload")
	}

	if fileName != nil {
		ctx.Locals("filename", *fileName)
	} else {
		ctx.Locals("filename", nil)
	}

	return ctx.Next()

}
