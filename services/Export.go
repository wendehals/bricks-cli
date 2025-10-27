package services

import (
	"archive/zip"
	"bufio"
	"embed"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"text/template"

	"github.com/wendehals/bricks/cmd/options"
	"github.com/wendehals/bricks/model"
	"github.com/wendehals/bricks/utils"
)

const (
	PARSING_FAILED_MSG = "parsing template file '%s' failed: %s"
	EXPORT_FAILED_MSG  = "exporting collection to HTML file '%s' failed: %s"
)

//go:embed resources/build.gotpl
//go:embed resources/collection.gotpl
//go:embed resources/part.gotpl
var embeddedFS embed.FS

// ExportToHTML writes an HTML file with all parts of the collection into the given export directory.
func ExportCollectionToHTML(collection *model.Collection, folderName string, fileName string) {
	utils.CreateFolder(folderName)

	for i := range collection.Parts {
		replaceImageURL(&collection.Parts[i], folderName)
	}

	exportFilePath := filepath.FromSlash(fmt.Sprintf("%s/%s.html", folderName, fileName))
	exportTemplate(collection, exportFilePath, "resources/collection.gotpl", "resources/part.gotpl")

	log.Printf("Exported result to '%s'", exportFilePath)
}

// ExportBuildCollectionToHTML writes an HTML file with the build collection into the given export directory.
func ExportBuildCollectionToHTML(buildCollection *model.BuildCollection, folderName string, fileName string) {
	utils.CreateFolder(folderName)

	for i := range buildCollection.Mapping {
		replaceImageURL(&buildCollection.Mapping[i].Original, folderName)
		for j := range buildCollection.Mapping[i].Substitutes {
			replaceImageURL(&buildCollection.Mapping[i].Substitutes[j], folderName)
		}
	}

	exportFilePath := filepath.FromSlash(fmt.Sprintf("%s/%s.html", folderName, fileName))
	exportTemplate(buildCollection, exportFilePath, "resources/build.gotpl", "resources/part.gotpl")

	log.Printf("Exported result to '%s'", exportFilePath)
}

func exportTemplate(data any, exportFilePath string, files ...string) {
	templ := parseTemplates(files...)

	file, err := os.Create(exportFilePath)
	if err != nil {
		log.Fatalf(EXPORT_FAILED_MSG, exportFilePath, err.Error())
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	err = templ.Execute(writer, data)
	if err != nil {
		log.Fatalf(EXPORT_FAILED_MSG, exportFilePath, err.Error())
	}

	writer.Flush()
}

func parseTemplates(files ...string) *template.Template {
	name, _ := utils.SplitFileName(files[0])
	templ := template.New(name)

	templ.Funcs(template.FuncMap{
		"abs": func(i int) int {
			if i < 0 {
				return -i
			}
			return i
		},
	})

	bytes, err := embeddedFS.ReadFile(files[0])
	if err != nil {
		log.Fatalf(PARSING_FAILED_MSG, files[0], err.Error())
	}

	templ, err = templ.Parse(string(bytes))
	if err != nil {
		log.Fatalf(PARSING_FAILED_MSG, files[0], err.Error())
	}

	for i := 1; i < len(files); i++ {
		name, _ := utils.SplitFileName(files[i])
		additional_templ := templ.New(name)

		bytes, err := embeddedFS.ReadFile(files[i])
		if err != nil {
			log.Fatalf(PARSING_FAILED_MSG, files[i], err.Error())
		}

		_, err = additional_templ.Parse(string(bytes))
		if err != nil {
			log.Fatalf(PARSING_FAILED_MSG, files[i], err.Error())
		}
	}

	return templ
}

func replaceImageURL(part *model.Part, exportDir string) {
	if part.Color.ID >= 0 {
		imageUrl, err := extractImage(part.Shape.Number, part.Color.ID, exportDir)
		if err == nil {
			part.Shape.ImageURL = imageUrl
		} else if options.Verbose {
			log.Print(err)
		}
	}
}

func extractImage(partNumber string, colorId int, exportDir string) (string, error) {
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

			utils.CreateFolder(filepath.Join(exportDir, strconv.Itoa(colorId)))

			imageSubPath := filepath.Join(strconv.Itoa(colorId), imageFile.Name)
			filePath := filepath.Join(exportDir, imageSubPath)

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

			return imageSubPath, nil
		}
	}

	return "", fmt.Errorf("no image file found for part number %s in color %d", partNumber, colorId)
}

func findImagesFileForColor(colorId int) (string, error) {
	imagesFile := filepath.FromSlash(fmt.Sprintf("%s/parts_%d.zip", utils.CacheDir(), colorId))
	if _, err := os.Stat(imagesFile); os.IsNotExist(err) {
		return "", err
	}

	return imagesFile, nil
}
