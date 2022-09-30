package export

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
	EXPORT_FAILED_MSG = "exporting collection to HTML file '%s' failed: %s"
)

//go:embed resources/build_html.gotpl
//go:embed resources/parts_html.gotpl
var embeddedFS embed.FS

// ExportToHTML writes an HTML file with all parts of the collection into the given export directory.
func ExportCollectionToHTML(collection *model.Collection, folderName string, fileName string) {
	utils.CreateFolder(folderName)

	for i := range collection.Parts {
		replaceImageURL(&collection.Parts[i], folderName)
	}

	exportFilePath := filepath.FromSlash(fmt.Sprintf("%s/%s.html", folderName, fileName))
	exportTemplate("resources/parts_html.gotpl", exportFilePath, collection)

	log.Printf("Exported result to '%s'", exportFilePath)
}

// ExportToHTML writes an HTML file with the build collection into the given export directory.
func ExportBuildToHTML(buildCollection *model.BuildCollection, folderName string, fileName string) {
	utils.CreateFolder(folderName)

	for i := range buildCollection.Mapping {
		replaceImageURL(&buildCollection.Mapping[i].Original, folderName)
		for j := range buildCollection.Mapping[i].Substitutes {
			replaceImageURL(&buildCollection.Mapping[i].Substitutes[j], folderName)
		}
	}

	exportFilePath := filepath.FromSlash(fmt.Sprintf("%s/%s.html", folderName, fileName))
	exportTemplate("resources/build_html.gotpl", exportFilePath, buildCollection)

	log.Printf("Exported result to '%s'", exportFilePath)
}

func exportTemplate(templateFile string, exportFilePath string, data any) {
	templ := template.New("parts")
	templ.Funcs(template.FuncMap{
		"abs": func(i int) int {
			if i < 0 {
				return -i
			}
			return i
		},
	})

	bytes, err := embeddedFS.ReadFile(templateFile)
	if err != nil {
		log.Fatalf(EXPORT_FAILED_MSG, exportFilePath, err.Error())
	}

	templ, err = templ.Parse(string(bytes))
	if err != nil {
		log.Fatalf(EXPORT_FAILED_MSG, exportFilePath, err.Error())
	}

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

func replaceImageURL(part *model.Part, exportDir string) {
	if part.Color.ID >= 0 {
		imageUrl, err := extractImage(part.Shape.Number, part.Color.ID, exportDir)
		if err == nil {
			part.Shape.ImageURL = imageUrl
		} else if err != nil && options.Verbose {
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
	imagesFile := filepath.FromSlash(fmt.Sprintf("%s/parts_%d.zip", utils.GetBricksDir(), colorId))
	if _, err := os.Stat(imagesFile); os.IsNotExist(err) {
		return "", err
	}

	return imagesFile, nil
}
