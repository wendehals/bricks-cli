package model

import (
	"io"
	"log"
	"os"
	"strconv"

	"github.com/wendehals/bricks/utils"
)

type Colors struct {
	Colors []Color `json:"colors"`
}

var (
	colors *Colors
)

func GetColors() *Colors {
	if colors != nil {
		return colors
	}

	colorsPath := utils.ColorsPath()
	if utils.FileExists(colorsPath) {
		colors = Load(&Colors{}, colorsPath)
		return colors
	}

	colors := &Colors{}
	colors.Colors = []Color{}

	csvFile := utils.DownloadColors()
	csvReader := utils.GzipCSVReader(csvFile)
	for {
		record, err := csvReader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}

		color := Color{}
		color.ID, err = strconv.Atoi(record[0])
		if err != nil {
			log.Fatal(err)
		}
		color.Name = record[1]
		colors.Colors = append(colors.Colors, color)
	}
	Save(colors, colorsPath)
	os.Remove(csvFile)

	return colors
}
