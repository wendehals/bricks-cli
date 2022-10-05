package utils

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

const (
	REBRICKABLE_DOWNLOADS      string = "https://rebrickable.com/downloads"
	COLORS_REG_EXP             string = "<a href=\"(https://cdn\\.rebrickable\\.com/media/downloads/colors\\.csv\\.gz.*)\">"
	PART_RELATIONSHIPS_REG_EXP string = "<a href=\"(https://cdn\\.rebrickable\\.com/media/downloads/part_relationships\\.csv\\.gz.*)\">"
	PART_IMAGES_REG_EXP        string = "<a href=\"(https://cdn\\.rebrickable\\.com/media/downloads/ldraw/parts_\\d+.zip)\">"
	SHAPES_REG_EXP             string = "<a href=\"(https://cdn\\.rebrickable\\.com/media/downloads/parts\\.csv\\.gz.*)\">"
)

func DownloadColors() string {
	log.Print("Loading colors file from rebrickable.com... ")

	downloadUrls := findDownloadURLs(COLORS_REG_EXP)
	if len(downloadUrls) == 0 {
		log.Fatal("no URL for colors.csv.gz found!")
	}

	return downloadFileFromURL(downloadUrls[0], "colors.csv.gz")
}

func DownloadPartRelationships() string {
	log.Print("Loading part relationships file from rebrickable.com... ")

	downloadUrls := findDownloadURLs(PART_RELATIONSHIPS_REG_EXP)
	if len(downloadUrls) == 0 {
		log.Fatal("no URL for part_relationships.csv.gz found!")
	}

	return downloadFileFromURL(downloadUrls[0], "part_relationships.csv.gz")
}

func DownloadShapes() string {
	log.Print("Loading parts file from rebrickable.com... ")

	downloadUrls := findDownloadURLs(SHAPES_REG_EXP)
	if len(downloadUrls) == 0 {
		log.Fatal("no URL for parts.csv.gz found!")
	}

	return downloadFileFromURL(downloadUrls[0], "parts.csv.gz")
}

func DownloadPartImages(update bool) {
	log.Print("Loading part images file from rebrickable.com...")

	partImagesURLs := findDownloadURLs(PART_IMAGES_REG_EXP)
	if len(partImagesURLs) == 0 {
		log.Fatal("no URLs for part images found!")
	}

	bricksDir := GetBricksDir()

	for _, partImagesURL := range partImagesURLs {
		log.Printf("Loading %s...", partImagesURL)

		response, err := http.Get(partImagesURL)
		if err != nil {
			log.Fatal(err)
		}
		defer response.Body.Close()

		splitURL := strings.Split(partImagesURL, "/")
		fileName := filepath.FromSlash(fmt.Sprintf("%s/%s", bricksDir, splitURL[len(splitURL)-1]))

		if FileExists(fileName) && !update {
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
