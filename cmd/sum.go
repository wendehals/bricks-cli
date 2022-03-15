/*
Copyright © 2022 Lothar Wendehals

*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/wendehals/bricks/model"
)

var sumCmd = &cobra.Command{
	Use:   "sum [-c FILE] [-o FILE] [-j FILE] JSONFILE...",
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

func init() {
	rootCmd.AddCommand(sumCmd)
}

func executeSum(args []string) error {
	var collections []model.Collection
	for _, filename := range args {
		collection, err := model.Import(filename)
		if err != nil {
			return err
		}
		collections = append(collections, *collection)
	}

	sum := model.Collection{}
	for _, collection := range collections {
		sum.Add(&collection)
	}

	sum.ExportToHTML(htmlFile)

	if jsonFile != "" {
		sum.Export(jsonFile)
	}

	return nil
}
