package services

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/wendehals/bricks/model"
	"github.com/wendehals/bricks/utils"
)

const (
	REBRICKABLE_DOWNLOADS      string = "https://rebrickable.com/downloads"
	COLORS_REG_EXP             string = "<a href=\"(https://cdn\\.rebrickable\\.com/media/downloads/colors\\.csv\\.gz.*)\">"
	PART_RELATIONSHIPS_REG_EXP string = "<a href=\"(https://cdn\\.rebrickable\\.com/media/downloads/part_relationships\\.csv\\.gz.*)\">"
	PART_IMAGES_REG_EXP        string = "<a href=\"(https://cdn\\.rebrickable\\.com/media/downloads/ldraw/parts_\\d+.zip)\">"
	SHAPES_REG_EXP             string = "<a href=\"(https://cdn\\.rebrickable\\.com/media/downloads/parts\\.csv\\.gz.*)\">"
)

func DownloadColors() *model.Colors {
	log.Print("Loading colors file from rebrickable.com... ")

	colors := model.NewColors()

	csvFile := downloadColorsCSVFile()
	csvReader := utils.GzipCSVReader(csvFile)
	for {
		record, err := csvReader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}

		color := model.Color{}
		color.ID, err = strconv.Atoi(record[0])
		if err != nil {
			log.Fatal(err)
		}
		color.Name = record[1]
		colors.Colors = append(colors.Colors, color)
	}

	model.Save(colors, utils.ColorsPath())

	os.Remove(csvFile)

	return colors
}

func DownloadPartRelationships() (*model.Alternates, *model.Molds, *model.Prints) {
	log.Print("Loading part relationships file from rebrickable.com... ")

	alternates := model.NewAlternates()
	molds := model.NewMolds()
	prints := model.NewPrints()

	csvFile := downloadPartRelationshipsCSVFile()
	csvReader := utils.GzipCSVReader(csvFile)
	for {
		record, err := csvReader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}

		switch record[0] {
		case "A":
			alternates.Add(record[1], record[2])
		case "M":
			molds.Add(record[1], record[2])
		case "P":
			prints.Add(record[1], record[2])
		}
	}

	alternates.TransitiveClosure()
	molds.TransitiveClosure()

	model.Save(alternates, utils.AlternatesPath())
	model.Save(molds, utils.MoldsPath())
	model.Save(prints, utils.PrintsPath())

	os.Remove(csvFile)

	return alternates, molds, prints
}

func DownloadShapes() *model.Shapes {
	log.Print("Loading parts file from rebrickable.com... ")

	shapes := model.NewShapes()

	csvFile := downloadShapesCSVFile()
	csvReader := utils.GzipCSVReader(csvFile)
	for {
		record, err := csvReader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}

		shape := model.Shape{}
		shape.Number = record[0]
		shape.Name = record[1]
		shapes.Shapes = append(shapes.Shapes, shape)
	}
	model.Save(shapes, utils.ShapesPath())
	os.Remove(csvFile)

	return shapes
}

func DownloadPartImages(update bool) {
	log.Print("Loading part images file from rebrickable.com...")

	partImagesURLs := findDownloadURLs(PART_IMAGES_REG_EXP)
	if len(partImagesURLs) == 0 {
		log.Fatal("no URLs for part images found!")
	}

	bricksDir := utils.GetBricksDir()

	for _, partImagesURL := range partImagesURLs {
		log.Printf("Loading %s...", partImagesURL)

		response, err := http.Get(partImagesURL)
		if err != nil {
			log.Fatal(err)
		}
		defer response.Body.Close()

		splitURL := strings.Split(partImagesURL, "/")
		fileName := filepath.FromSlash(fmt.Sprintf("%s/%s", bricksDir, splitURL[len(splitURL)-1]))

		if utils.FileExists(fileName) && !update {
			continue
		}

		file, err := os.Create(fileName)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		_, err = io.Copy(file, response.Body)
		if err != nil {
			log.Fatal(err)
		}

		log.Println("done")
	}

	log.Println("finished")
}

func downloadColorsCSVFile() string {
	downloadUrls := findDownloadURLs(COLORS_REG_EXP)
	if len(downloadUrls) == 0 {
		log.Fatal("no URL for colors.csv.gz found!")
	}

	return downloadFileFromURL(downloadUrls[0], "colors.csv.gz")
}

func downloadShapesCSVFile() string {
	downloadUrls := findDownloadURLs(SHAPES_REG_EXP)
	if len(downloadUrls) == 0 {
		log.Fatal("no URL for parts.csv.gz found!")
	}

	return downloadFileFromURL(downloadUrls[0], "parts.csv.gz")
}

func downloadPartRelationshipsCSVFile() string {
	downloadUrls := findDownloadURLs(PART_RELATIONSHIPS_REG_EXP)
	if len(downloadUrls) == 0 {
		log.Fatal("no URL for part_relationships.csv.gz found!")
	}

	return downloadFileFromURL(downloadUrls[0], "part_relationships.csv.gz")
}

func findDownloadURLs(regex string) []string {
	response, err := http.Get(REBRICKABLE_DOWNLOADS)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	var downloadUrls []string
	r := regexp.MustCompile(regex)
	scanner := bufio.NewScanner(response.Body)
	for scanner.Scan() {
		match := r.FindStringSubmatch(scanner.Text())
		if match != nil {
			downloadUrls = append(downloadUrls, match[1])
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return downloadUrls
}

func downloadFileFromURL(url, fileName string) string {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	path := filepath.FromSlash(fmt.Sprintf("%s/%s", os.TempDir(), fileName))
	file, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("done")

	return path
}
