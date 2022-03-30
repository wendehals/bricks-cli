package cmd

import (
	"github.com/spf13/cobra"
	"github.com/wendehals/bricks/model"
)

var sortCmd = &cobra.Command{
	Use:   "sort [-c FILE] [-o FILE] [-j FILE] JSONFILE",
	Short: "Sorts the parts of a collection by their number.",
	Long: `
The command sorts the parts of a collection in descending order by their part number.`,

	DisableFlagsInUseLine: true,

	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return executeSort(args)
	},
}

func init() {
	rootCmd.AddCommand(sortCmd)
}

func executeSort(args []string) error {
	collection, err := model.Import(args[0])
	if err != nil {
		return err
	}

	sorted := collection.Sort()

	err = sorted.ExportToHTML(htmlFile)
	if err != nil {
		return err
	}

	if jsonFile != "" {
		err := model.ExportToJSON(jsonFile, sorted)
		if err != nil {
			return err
		}
	}

	return nil
}
