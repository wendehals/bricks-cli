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

func CacheDir() string {
	userCacheDir, err := os.UserCacheDir()
	if err != nil {
		log.Fatal(err)
	}

	bricksDir := filepath.Join(userCacheDir, "bricks-cli")
	if !FileExists(bricksDir) {
		err = os.Mkdir(bricksDir, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
	}

	return bricksDir
}

func ColorsPath() string {
	return filepath.Join(CacheDir(), "colors.json")
}

func AlternatesPath() string {
	return filepath.Join(CacheDir(), "alternates.json")
}

func MoldsPath() string {
	return filepath.Join(CacheDir(), "molds.json")
}

func PrintsPath() string {
	return filepath.Join(CacheDir(), "prints.json")
}

func ShapesPath() string {
	return filepath.Join(CacheDir(), "shapes.json")
}

func SetsPath() string {
	setsPath := filepath.Join(CacheDir(), "sets")
	if !FileExists(setsPath) {
		err := os.Mkdir(setsPath, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
	}

	return setsPath
}

func CredentialsDefaultPath() string {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	return filepath.Join(userHomeDir, ".bricks-credentials.json")
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
