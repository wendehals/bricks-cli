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
	PART_RELATIONSHIPS_REG_EXP string = "<a href=\"(https://cdn\\.rebrickable\\.com/media/downloads/part_relationships\\.csv\\.gz.*)\">"
	PART_IMAGES_REG_EXP        string = "<a href=\"(https://cdn\\.rebrickable\\.com/media/downloads/ldraw/parts_\\d+.zip)\">"
)

func DownloadPartRelationships() string {
	log.Print("Loading part relationships file from rebrickable.com... ")

	downloadUrl := getPartRelationshipsURL()
	if downloadUrl == "" {
		log.Fatalln("no URL for part_relationships.csv.gz found!")
	}

	response, err := http.Get(downloadUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	fileName := filepath.FromSlash(fmt.Sprintf("%s/part_relationships.csv.gz", os.TempDir()))
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

	return fileName
}

func getPartRelationshipsURL() string {
	response, err := http.Get(REBRICKABLE_DOWNLOADS)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	downloadUrl := ""
	r := regexp.MustCompile(PART_RELATIONSHIPS_REG_EXP)
	scanner := bufio.NewScanner(response.Body)
	for scanner.Scan() {
		match := r.FindStringSubmatch(scanner.Text())
		if len(match) > 0 {
			downloadUrl = match[1]
			break
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return downloadUrl
}

func DownloadPartImages(update bool) {
	log.Print("Loading part images file from rebrickable.com...")

	partImagesURLs := getPartImagesURLs()
	if len(partImagesURLs) == 0 {
		log.Fatalln("no URLs for part images found!")
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

func getPartImagesURLs() []string {
	response, err := http.Get(REBRICKABLE_DOWNLOADS)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	var downloadUrls []string
	r := regexp.MustCompile(PART_IMAGES_REG_EXP)
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
