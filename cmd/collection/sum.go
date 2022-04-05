package collection

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wendehals/bricks/cmd/options"
	"github.com/wendehals/bricks/model"
)

var sumCmd = &cobra.Command{
	Use:   fmt.Sprintf("sum %s JSON_FILE1 JSON_FILE2 ...", options.JSON_OUTPUT_ARG),
	Short: "Sums up the parts of multiple collections.",
	Long: `
The command sums up all parts of multiple collections to a new collection by
merging identical parts to single lots.`,

	DisableFlagsInUseLine: true,

	Args: cobra.MinimumNArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		return executeSum(args)
	},
}

func executeSum(args []string) error {
	var collections []model.Collection
	for _, filename := range args {
		collection, err := model.ImportCollection(filename)
		if err != nil {
			return err
		}
		collections = append(collections, *collection)
	}

	sum := model.Collection{}
	for _, collection := range collections {
		sum.Add(&collection)
	}

	if jsonFile == "" {
		jsonFile = options.FileNameFromArgs(args, "_parts.json")
	}

	return model.ExportToJSON(jsonFile, sum)
}
