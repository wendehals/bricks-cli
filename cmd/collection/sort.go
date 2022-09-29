package collection

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/wendehals/bricks/cmd/options"
	"github.com/wendehals/bricks/model"
)

var sortCmd = &cobra.Command{
	Use:   fmt.Sprintf("sort %s PARTS_FILE", options.OUTPUT_FILE_ARG),
	Short: "Sorts the parts of a collection by their color and name",
	Long:  "The command sorts the parts of a collection in descending order by their color and name.",

	DisableFlagsInUseLine: true,

	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		executeSort(args)
	},
}

func executeSort(args []string) {
	log.Print("Sorting the parts of the given collection")

	collection := model.Load(&model.Collection{}, args[0])

	if outputFile == "" {
		outputFile = options.FileNameFromArgs(args, "_sorted.parts")
	}

	model.Save(collection.Sort(), outputFile)
}
