package services

import (
	"fmt"
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

	addColorNames(collection)
	addShapeNames(collection)
	deducePartURLs(collection)

	return collection
}

func addColorNames(collection *model.Collection) {
	names := make(map[int]string)
	for _, color := range GetColors(false).Colors {
		names[color.ID] = color.Name
	}

	for i := range collection.Parts {
		collection.Parts[i].Color.Name = names[collection.Parts[i].Color.ID]
	}
}

func addShapeNames(collection *model.Collection) {
	names := make(map[string]string)
	for _, shape := range GetShapes(false).Shapes {
		names[shape.Number] = shape.Name
	}

	for i := range collection.Parts {
		collection.Parts[i].Shape.Name = names[collection.Parts[i].Shape.Number]
	}
}

func deducePartURLs(collection *model.Collection) {
	for i := range collection.Parts {
		url := fmt.Sprintf("https://rebrickable.com/parts/%s", collection.Parts[i].Shape.Number)
		collection.Parts[i].Shape.URL = url
	}
}
