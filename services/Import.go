package services

import (
	"io"
	"log"
	"strconv"

	"github.com/wendehals/bricks/model"
	"github.com/wendehals/bricks/utils"
)

func ImportCSVCollection(csvFile string) *model.Collection {
	log.Printf("Importing parts from '%s'", csvFile)

	collection := model.NewCollection()

	csvReader := utils.CSVReader(csvFile)
	for {
		record, err := csvReader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}

		part := model.NewPart()
		part.Quantity, _ = strconv.Atoi(record[2])
		part.Shape.Number = record[0]
		part.Color.ID, _ = strconv.Atoi(record[1])

		collection.Parts = append(collection.Parts, *part)
	}

	addMissingData(collection)

	return collection
}

func addMissingData(collection *model.Collection) {
	shapes := GetShapes(false)
	colorNames := make(map[int]string)

	for _, color := range GetColors(false).Colors {
		colorNames[color.ID] = color.Name
	}

	for i := range collection.Parts {
		shape, found := shapes.GetShape(collection.Parts[i].Shape.Number)
		if found {
			collection.Parts[i].Color.Name = colorNames[collection.Parts[i].Color.ID]
			collection.Parts[i].Shape.Name = shape.Name
			collection.Parts[i].Shape.URL = shape.URL
			collection.Parts[i].Shape.ImageURL = shape.ImageURL
		}
	}
}
