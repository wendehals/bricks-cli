package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wendehals/bricks/model"
)

var sumCmd = &cobra.Command{
	Use:   fmt.Sprintf("sum %s JSON_FILE...", json_output_arg),
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
	collectionCmd.AddCommand(sumCmd)
	sumCmd.Flags().StringVarP(&jsonFile, json_output_opt, json_output_sopt, "", json_output_usage)
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

	if jsonFile == "" {
		jsonFile = fileNameFromArgs(args, "_sum.json")
	}

	return model.ExportToJSON(jsonFile, sum)
}
