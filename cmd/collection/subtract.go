package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wendehals/bricks/model"
)

var subtractCmd = &cobra.Command{
	Use:   fmt.Sprintf("subtract %s JSON_FILE_MINUEND JSON_FILE_SUBTRAHEND", json_output_arg),
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

func init() {
	collectionCmd.AddCommand(subtractCmd)

	subtractCmd.Flags().StringVarP(&jsonFile, json_output_opt, json_output_sopt, "", json_output_usage)
}

func executeSubtract(args []string) error {
	minuend, err := model.Import(args[0])
	if err != nil {
		return err
	}

	subtrahend, err := model.Import(args[1])
	if err != nil {
		return err
	}

	result := minuend.Subtract(subtrahend)

	if jsonFile == "" {
		jsonFile = fileNameFromArgs(args, "_subtracted.json")
	}

	return model.ExportToJSON(jsonFile, result)
}