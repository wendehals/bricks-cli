package model

import (
	"io"
	"log"
	"os"

	"github.com/wendehals/bricks/utils"
)

type Shapes struct {
	Shapes []Shape `json:"shapes"`
}

var shapes *Shapes

func GetShapes() *Shapes {
	if shapes != nil {
		return shapes
	}

	shapesPath := utils.ShapesPath()
	if utils.FileExists(shapesPath) {
		shapes = Load(&Shapes{}, shapesPath)
		return shapes
	}

	shapes := &Shapes{}
	shapes.Shapes = []Shape{}

	csvFile := utils.DownloadShapes()
	csvReader := utils.GzipCSVReader(csvFile)
	for {
		record, err := csvReader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}

		shape := Shape{}
		shape.Number = record[0]
		shape.Name = record[1]
		shapes.Shapes = append(shapes.Shapes, shape)
	}
	Save(shapes, shapesPath)
	os.Remove(csvFile)

	return shapes
}
