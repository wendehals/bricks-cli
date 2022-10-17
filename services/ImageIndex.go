package services

import (
	"archive/zip"
	"fmt"
	"log"

	"github.com/wendehals/bricks/model"
	"github.com/wendehals/bricks/utils"
	"golang.org/x/exp/maps"
)

type ImageIndex struct {
	Entries map[string]utils.StringSet `json:"entries"`
}

var imageIndex *ImageIndex

func NewImageIndex() *ImageIndex {
	return &ImageIndex{Entries: make(map[string]utils.StringSet)}
}

func CreateImageIndex() *ImageIndex {
	log.Print("Creating index for images...")

	imageIndex := NewImageIndex()

	colors := GetColors(false)

	for _, color := range colors.Colors {
		imagesFile, err := findImagesFileForColor(color.ID)
		if err != nil {
			continue
		}

		archive, err := zip.OpenReader(imagesFile)
		if err != nil {
			continue
		}
		defer archive.Close()

		for _, imageFile := range archive.File {
			part, _ := utils.SplitFileName(imageFile.Name)
			if _, found := imageIndex.Entries[part]; !found {
				imageIndex.Entries[part] = utils.NewStringSet()
			}
			imageIndex.Entries[part].Add(fmt.Sprint(color.ID))
		}
	}

	return imageIndex
}

func GetImageIndex() *ImageIndex {
	if imageIndex != nil {
		return imageIndex
	}

	imageIndexPath := utils.ImageIndexPath()
	if utils.FileExists(imageIndexPath) {
		imageIndex = model.Load(NewImageIndex(), imageIndexPath)
		return imageIndex
	}

	log.Print("no image index could be found")

	return nil
}

var COLORS_WITH_BEST_CONTRAST = []string{"71", "7", "8", "72", "1", "4", "2"}

func (i *ImageIndex) FindBestColor(part string) string {
	if set, found := i.Entries[part]; found {
		for _, color := range COLORS_WITH_BEST_CONTRAST {
			if set.Contains(color) {
				return color
			}
		}
		return maps.Keys(set.Values)[0]
	}

	return "-1"
}
