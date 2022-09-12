package collection

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wendehals/bricks/cmd/options"
	"github.com/wendehals/bricks/model"
)

var sortCmd = &cobra.Command{
	Use:   fmt.Sprintf("sort %s PARTS_FILE", options.OUTPUT_FILE_ARG),
	Short: "Sorts the parts of a collection by their number",
	Long:  "The command sorts the parts of a collection in descending order by their part number.",

	DisableFlagsInUseLine: true,

	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		executeSort(args)
	},
}

func executeSort(args []string) {
	collection := model.Load(&model.Collection{}, args[0])

	if jsonFile == "" {
		jsonFile = options.FileNameFromArgs(args, "_sorted.parts")
	}

	model.Save(collection.Sort(), jsonFile)
}
