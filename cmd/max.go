package cmd

import (
	"github.com/spf13/cobra"
	"github.com/wendehals/bricks/model"
)

var maxCmd = &cobra.Command{
	Use:   "max [-c FILE] [-o FILE] [-j FILE] JSONFILE...",
	Short: "Calculates the maximum quantity of each part of multiple collections.",
	Long: `
The command calculates the maximum quantity of each part of multiple collections.`,

	DisableFlagsInUseLine: true,

	Args: cobra.MinimumNArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		return executeMax(args)
	},
}

func init() {
	rootCmd.AddCommand(maxCmd)
}

func executeMax(args []string) error {
	var collections []model.Collection
	for _, filename := range args {
		collection, err := model.Import(filename)
		if err != nil {
			return err
		}
		collections = append(collections, *collection)
	}

	max := model.Collection{}
	for _, collection := range collections {
		max.Max(&collection)
	}

	max.ExportToHTML(htmlFile)

	if jsonFile != "" {
		model.ExportToJSON(jsonFile, max)
	}

	return nil
}
