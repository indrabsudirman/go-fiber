package utils

import (
	"fmt"
	"go-fiber/config"
	"log"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

const DefaultPathPublicImages = "./public/images/"

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
		log.Println("path is ", filepath.Ext(*fileName))
		extensionFile := filepath.Ext(*fileName)
		newFileName := fmt.Sprintf("profile-1%s", extensionFile)

		errSaveFile := ctx.SaveFile(file, fmt.Sprintf(config.ProjectRootPath+"/public/images/%s", newFileName))
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
			extensionFile := filepath.Ext(file.Filename)
			fileName = fmt.Sprintf("%d-%s%s", i, "image", extensionFile)

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

func HandleRemoveFile(filename string, path ...string) error {

	if len(path) > 0 {
		err := os.Remove(path[0] + filename)
		if err != nil {
			log.Println("failed to remove file", err)
			return err
		}
	} else {
		err := os.Remove(DefaultPathPublicImages + filename)
		if err != nil {
			log.Println("failed to remove file", err)
			return err
		}
	}

	return nil
}
