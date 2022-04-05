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
	RunE: func(cmd *cobra.Command, args []string) error {
		return executeMax(args)
	},
}

func executeMax(args []string) error {
	var collections []model.Collection
	for _, filename := range args {
		collection, err := model.ImportCollection(filename)
		if err != nil {
			return err
		}
		collections = append(collections, *collection)
	}

	max := model.Collection{}
	for _, collection := range collections {
		max.Max(&collection)
	}

	if jsonFile == "" {
		jsonFile = options.FileNameFromArgs(args, "_max_partss.json")
	}

	return model.ExportToJSON(jsonFile, max)
}
