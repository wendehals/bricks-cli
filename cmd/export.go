package cmd

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/wendehals/bricks/api"
	"github.com/wendehals/bricks/cmd/options"
	"github.com/wendehals/bricks/model"
)

var (
	exportDir string

	exportCmd = &cobra.Command{
		Use:   fmt.Sprintf("export %s %s PARTS_FILE", options.CREDENTIALS_ARG, options.OUTPUT_DIR_ARG),
		Short: "Exports the parts input file as HTML",
		Long:  "The command exports the parts input file as an HTML file to the given output directory.",

		DisableFlagsInUseLine: true,

		Args: cobra.ExactArgs(1),

		PreRunE: func(cmd *cobra.Command, args []string) error {
			var err error
			credentials, err = api.ImportCredentials(credentialsFile)
			return err
		},
		Run: func(cmd *cobra.Command, args []string) {
			executeExport(args)
		},
	}
)

func init() {
	exportCmd.Flags().StringVarP(&credentialsFile, options.CREDENTIALS_OPT, options.CREDENTIALS_SOPT,
		"credentials.json", options.CREDENTIALS_USAGE)
	exportCmd.Flags().StringVarP(&exportDir, options.OUTPUT_DIR_OPT, options.OUTPUT_DIR_SOPT,
		"", options.OUTPUT_DIR_USAGE)
}

func executeExport(args []string) {
	if exportDir == "" {
		exportDir = options.FileNameFromArgs(args, "")
	}

	log.Printf("Exporting '%s' to directory '%s'", args[0], exportDir)

	err := os.MkdirAll(exportDir, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	collection := model.ImportCollection(args[0])
	replaceImageURLs(collection, exportDir)

	htmlFile := filepath.FromSlash(fmt.Sprintf("%s/%s.html", exportDir, exportDir))
	collection.ExportToHTML(htmlFile)
}

func replaceImageURLs(collection *model.Collection, exportDir string) {
	for _, partEntry := range collection.Parts {
		imageUrl, err := extractImage(partEntry.Part.Number, partEntry.Color.ID, exportDir)
		if err != nil {
			partEntry.Part.ImageURL = imageUrl
		} else if err != nil && options.Verbose {
			log.Print(err)
		}
	}
}

func extractImage(partNumber string, colorId uint, exportDir string) (string, error) {
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

func findImagesFileForColor(colorId uint) (string, error) {
	userHome, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	imagesFile := filepath.FromSlash(fmt.Sprintf("%s/.bricks-cli/parts_%d.zip", userHome, colorId))
	if _, err := os.Stat(imagesFile); os.IsNotExist(err) {
		return "", err
	}

	return imagesFile, nil
}
