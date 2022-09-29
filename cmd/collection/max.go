package collection

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/wendehals/bricks/cmd/options"
	"github.com/wendehals/bricks/model"
)

var maxCmd = &cobra.Command{
	Use:   fmt.Sprintf("max %s PARTS_FILE1 PARTS_FILE2 ...", options.OUTPUT_FILE_ARG),
	Short: "Calculates the maximum quantity of each part of at least two collections",
	Long:  "The command calculates the maximum quantity of each part of at least two collections.",

	DisableFlagsInUseLine: true,

	Args: cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		executeMax(args)
	},
}

func executeMax(args []string) {
	log.Print("Calculating maximum of each part of the given collections")

	var collections []model.Collection
	for _, filename := range args {
		collections = append(collections, *model.Load(&model.Collection{}, filename))
	}

	max := model.Collection{}
	for _, collection := range collections {
		max.Max(&collection)
	}

	if outputFile == "" {
		outputFile = options.FileNameFromArgs(args, "_max.parts")
	}

	model.Save(max, outputFile)
}
