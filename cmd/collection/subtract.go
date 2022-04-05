package collection

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wendehals/bricks/cmd/options"
	"github.com/wendehals/bricks/model"
)

var subtractCmd = &cobra.Command{
	Use:   fmt.Sprintf("subtract %s JSON_FILE_MINUEND JSON_FILE_SUBTRAHEND", options.JSON_OUTPUT_ARG),
	Short: "Subtracts one collection of parts from another.",
	Long: `
The command subtracts the second collection of parts given in the args from the
first collection.

If the quantity of a certain part is zero the lot of this part is removed from
the resulting collection. If the quantity is negative the lot will be added to
the list of missing parts.`,

	DisableFlagsInUseLine: true,

	Args: cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		return executeSubtract(args)
	},
}

func executeSubtract(args []string) error {
	minuend, err := model.ImportCollection(args[0])
	if err != nil {
		return err
	}

	subtrahend, err := model.ImportCollection(args[1])
	if err != nil {
		return err
	}

	result := minuend.Subtract(subtrahend).RemoveQuantityZero()

	if jsonFile == "" {
		jsonFile = options.FileNameFromArgs(args, "_subtracted_parts.json")
	}

	return model.ExportToJSON(jsonFile, result)
}
