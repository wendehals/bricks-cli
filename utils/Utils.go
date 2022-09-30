package utils

import (
	"fmt"
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
// Example: SplitFileName("example.txt") returns ("example", "txt")
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

	bricksDir := filepath.FromSlash(fmt.Sprintf("%s/.bricks-cli", userHome))
	if _, err := os.Stat(bricksDir); os.IsNotExist(err) {
		err = os.Mkdir(bricksDir, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
	}

	return bricksDir
}

func GetPartRelationshipsPath() string {
	return filepath.FromSlash(fmt.Sprintf("%s/partRelationships.json", GetBricksDir()))
}
