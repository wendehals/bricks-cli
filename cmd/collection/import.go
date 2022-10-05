package collection

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/wendehals/bricks/cmd/options"
	"github.com/wendehals/bricks/model"
)

var (
	importCmd = &cobra.Command{
		Use:   fmt.Sprintf("import CSV_FILE %s", options.OUTPUT_FILE_ARG),
		Short: "Imports parts from a CSV file",
		Long:  "The command imports a collection of parts from a CSV file.",

		DisableFlagsInUseLine: true,

		Args: cobra.ExactArgs(1),

		Run: func(cmd *cobra.Command, args []string) {
			executeImport(args)
		},
	}
)

func executeImport(args []string) {
	if outputFile == "" {
		outputFile = options.FileNameFromArgs(args, ".parts")
	}

	log.Printf("Importing parts from '%s'", args[0])

	collection := model.ImportCSVCollection(args[0])
	addColorNames(collection)
	addShapeNames(collection)
	deducePartURL(collection)

	model.Save(collection, outputFile)
}

func addColorNames(collection *model.Collection) {
	names := make(map[int]string)
	for _, color := range model.GetColors(false).Colors {
		names[color.ID] = color.Name
	}

	for i := range collection.Parts {
		collection.Parts[i].Color.Name = names[collection.Parts[i].Color.ID]
	}
}

func addShapeNames(collection *model.Collection) {
	names := make(map[string]string)
	for _, shape := range model.GetShapes(false).Shapes {
		names[shape.Number] = shape.Name
	}

	for i := range collection.Parts {
		collection.Parts[i].Shape.Name = names[collection.Parts[i].Shape.Number]
	}
}

func deducePartURL(collection *model.Collection) {
	for i := range collection.Parts {
		url := fmt.Sprintf("https://rebrickable.com/parts/%s", collection.Parts[i].Shape.Number)
		collection.Parts[i].Shape.URL = url
	}
}
