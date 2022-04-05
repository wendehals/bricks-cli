package collection

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wendehals/bricks/cmd/options"
	"github.com/wendehals/bricks/model"
)

var sortCmd = &cobra.Command{
	Use:   fmt.Sprintf("sort %s JSON_FILE", options.JSON_OUTPUT_ARG),
	Short: "Sorts the parts of a collection by their number.",
	Long:  "The command sorts the parts of a collection in descending order by their part number.",

	DisableFlagsInUseLine: true,

	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return executeSort(args)
	},
}

func executeSort(args []string) error {
	collection, err := model.ImportCollection(args[0])
	if err != nil {
		return err
	}

	if jsonFile == "" {
		jsonFile = options.FileNameFromArgs(args, "_sorted_parts.json")
	}

	sorted := collection.Sort()

	return model.ExportToJSON(jsonFile, sorted)
}
