/*
Copyright © 2022 Lothar Wendehals

*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/wendehals/bricks/model"
)

var subtractCmd = &cobra.Command{
	Use:   "subtract [-c FILE] [-o FILE] [-j FILE] JSONFILE_MINUEND JSONFILE_SUBTRAHEND",
	Short: "Subtracts one collection of parts from another.",
	Long: `The command subtracts the second collection of parts given in the args from
the first collection.

If the quantity of a certain part is zero the lot of this part is removed
from the resulting collection. If the quantity is negative the lot will be added to
the list of missing parts.`,

	DisableFlagsInUseLine: true,

	Args: cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		return executeSubtract(args)
	},
}

func init() {
	rootCmd.AddCommand(subtractCmd)
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

	result.ExportToHTML(htmlFile)

	if jsonFile != "" {
		result.Export(jsonFile)
	}

	return nil
}
