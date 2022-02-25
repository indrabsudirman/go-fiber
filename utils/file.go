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

func HandleMultipleFile(ctx *fiber.Ctx) error {
	form, errForm := ctx.MultipartForm()
	if errForm != nil {
		log.Println("Error read multipart form request, error :", errForm)
	}

	files := form.File["photos"]

	var filenames []string

	for i, file := range files {

		var fileName string
		if file != nil {
			fileName = fmt.Sprintf("%d-%s", i, file.Filename)

			errSaveFile := ctx.SaveFile(file, fmt.Sprintf(config.ProjectRootPath+"/public/images/%s", fileName))
			if errSaveFile != nil {
				log.Println("Failed to store file : ", errSaveFile)
			}
		} else {
			log.Println("nothing file to be upload")
		}

		if fileName != "" {
			filenames = append(filenames, fileName)
		}
	}

	ctx.Locals("filenames", filenames)

	return ctx.Next()
}
