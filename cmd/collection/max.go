package collection

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wendehals/bricks/cmd/options"
	"github.com/wendehals/bricks/model"
)

var maxCmd = &cobra.Command{
	Use:   fmt.Sprintf("max %s JSON_FILE1 JSON_FILE2 ...", options.JSON_OUTPUT_ARG),
	Short: "Calculates the maximum quantity of each part of at least two collections.",
	Long:  "The command calculates the maximum quantity of each part of at least two collections.",

	DisableFlagsInUseLine: true,

	Args: cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		executeMax(args)
	},
}

func executeMax(args []string) {
	var collections []model.Collection
	for _, filename := range args {
		collections = append(collections, *model.ImportCollection(filename))
	}

	max := model.Collection{}
	for _, collection := range collections {
		max.Max(&collection)
	}

	if jsonFile == "" {
		jsonFile = options.FileNameFromArgs(args, "_max_partss.json")
	}

	model.ExportToJSON(jsonFile, max)
}
