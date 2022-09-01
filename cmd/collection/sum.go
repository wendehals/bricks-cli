package collection

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wendehals/bricks/cmd/options"
	"github.com/wendehals/bricks/model"
)

var sumCmd = &cobra.Command{
	Use:   fmt.Sprintf("sum %s JSON_FILE1 JSON_FILE2 ...", options.JSON_OUTPUT_ARG),
	Short: "Sums up the parts of multiple collections",
	Long: `
The command sums up all parts of multiple collections to a new collection by
merging identical parts to single lots.`,

	DisableFlagsInUseLine: true,

	Args: cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		executeSum(args)
	},
}

func executeSum(args []string) {
	var collections []model.Collection
	for _, filename := range args {
		collections = append(collections, *model.ImportCollection(filename))
	}

	sum := model.Collection{}
	for _, collection := range collections {
		sum.Add(&collection)
	}

	if jsonFile == "" {
		jsonFile = options.FileNameFromArgs(args, "_sum.parts")
	}

	model.Save(sum, jsonFile)
}
