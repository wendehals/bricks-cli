package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func FileExists(filePath string) bool {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	}

	return true
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
