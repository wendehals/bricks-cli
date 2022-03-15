/*
Copyright © 2022 Lothar Wendehals

*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/wendehals/bricks/model"
)

var addCmd = &cobra.Command{
	Use:   "add [-c FILE] [-o FILE] [-j FILE] JSONFILE...",
	Short: "Sums up the bricks of multiple collections.",
	Long: `add sums up all bricks of multiple collections to a new collection of bricks
by merging identical parts to single lots.`,

	DisableFlagsInUseLine: true,

	Args: cobra.MinimumNArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		return executeAdd(args)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}

func executeAdd(args []string) error {
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
