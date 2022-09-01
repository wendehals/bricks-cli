package model

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
)

const (
	REBRICKABLE_DOWNLOADS      string = "https://rebrickable.com/downloads"
	PART_RELATIONSHIPS_REG_EXP string = "<a href=\"(https://cdn\\.rebrickable\\.com/media/downloads/part_relationships\\.csv\\.gz.*)\">"
)

func DownloadPartRelationships() {
	log.Print("Loading part relationships file from rebrickable.com... ")

	downloadUrl := getPartRelationshipsURL()
	if downloadUrl == "" {
		log.Fatalln("no url for part_relationships.csv.gz found!")
	}

	response, err := http.Get(downloadUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	userHome, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	bricksDir := filepath.FromSlash(fmt.Sprintf("%s/.bricks-cli", userHome))
	if _, err := os.Stat(bricksDir); os.IsNotExist(err) {
		err = os.Mkdir(bricksDir, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
	}

	fileName := filepath.FromSlash(fmt.Sprintf("%s/part_relationships.csv.gz", bricksDir))
	if err != nil {
		log.Fatal(err)
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
		line := scanner.Text()
		match := r.FindStringSubmatch(line)
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
