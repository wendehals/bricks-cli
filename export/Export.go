package export

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/wendehals/bricks/cmd/options"
	"github.com/wendehals/bricks/utils"
)

func ExtractImage(partNumber string, colorId int, exportDir string) (string, error) {
	imagesFile, err := findImagesFileForColor(colorId)
	if err != nil {
		return "", err
	}

	archive, err := zip.OpenReader(imagesFile)
	if err != nil {
		return "", err
	}
	defer archive.Close()

	imageFileName := fmt.Sprintf("%s.png", partNumber)
	for _, imageFile := range archive.File {
		if imageFile.Name == imageFileName {
			if options.Verbose {
				log.Printf("Unzipping image file '%s'", imageFile.Name)
			}

			filePath := filepath.Join(exportDir, imageFile.Name)
			dstFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, imageFile.Mode())
			if err != nil {
				log.Fatal(err)
			}

			fileInArchive, err := imageFile.Open()
			if err != nil {
				log.Fatal(err)
			}

			if _, err := io.Copy(dstFile, fileInArchive); err != nil {
				log.Fatal(err)
			}

			dstFile.Close()
			fileInArchive.Close()

			return imageFile.Name, nil
		}
	}

	return "", fmt.Errorf("no image file found for part number %s in color %d", partNumber, colorId)
}

func findImagesFileForColor(colorId int) (string, error) {
	imagesFile := filepath.FromSlash(fmt.Sprintf("%s/parts_%d.zip", utils.GetBricksDir(), colorId))
	if _, err := os.Stat(imagesFile); os.IsNotExist(err) {
		return "", err
	}

	return imagesFile, nil
}
