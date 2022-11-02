package utils

import (
	"compress/gzip"
	"encoding/csv"
	"log"
	"os"
	"path/filepath"
)

func CreateFolder(folderName string) {
	err := os.MkdirAll(folderName, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
}

func FileExists(filePath string) bool {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	}

	return true
}

// SplitFileName returns the base file name and its extension.
// Example: SplitFileName("example.txt") returns ("example", ".txt")
func SplitFileName(fileName string) (string, string) {
	base := filepath.Base(fileName)
	ext := filepath.Ext(base)
	baseWoExt := base[:len(base)-len(ext)]

	return baseWoExt, ext
}

func GetBricksDir() string {
	userHome, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	bricksDir := filepath.Join(userHome, ".bricks-cli")
	if _, err := os.Stat(bricksDir); os.IsNotExist(err) {
		err = os.Mkdir(bricksDir, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
	}

	return bricksDir
}

func ColorsPath() string {
	return filepath.Join(GetBricksDir(), "colors.json")
}

func AlternatesPath() string {
	return filepath.Join(GetBricksDir(), "alternates.json")
}

func MoldsPath() string {
	return filepath.Join(GetBricksDir(), "molds.json")
}

func PrintsPath() string {
	return filepath.Join(GetBricksDir(), "prints.json")
}

func ShapesPath() string {
	return filepath.Join(GetBricksDir(), "shapes.json")
}

func ImageIndexPath() string {
	return filepath.Join(GetBricksDir(), "image_index.json")
}

func CSVReader(csvFile string) *csv.Reader {
	file, err := os.Open(csvFile)
	if err != nil {
		log.Fatal(err)
	}

	csvReader := csv.NewReader(file)
	_, err = csvReader.Read() // omit header
	if err != nil {
		log.Fatal(err)
	}

	return csvReader
}

func GzipCSVReader(csvFile string) *csv.Reader {
	file, err := os.Open(csvFile)
	if err != nil {
		log.Fatal(err)
	}

	gzReader, err := gzip.NewReader(file)
	if err != nil {
		log.Fatal(err)
	}

	csvReader := csv.NewReader(gzReader)
	_, err = csvReader.Read() // omit header
	if err != nil {
		log.Fatal(err)
	}

	return csvReader
}
