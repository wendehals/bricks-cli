package collection

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wendehals/bricks/model"
)

var sortCmd = &cobra.Command{
	Use:   fmt.Sprintf("sort %s JSON_FILE", json_output_arg),
	Short: "Sorts the parts of a collection by their number.",
	Long:  "The command sorts the parts of a collection in descending order by their part number.",

	DisableFlagsInUseLine: true,

	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return executeSort(args)
	},
}

func init() {
	sortCmd.Flags().StringVarP(&jsonFile, json_output_opt, json_output_sopt, "", json_output_usage)
}

func executeSort(args []string) error {
	collection, err := model.Import(args[0])
	if err != nil {
		return err
	}

	if jsonFile == "" {
		jsonFile = fileNameFromArgs(args, "_sorted.json")
	}

	sorted := collection.Sort()

	return model.ExportToJSON(jsonFile, sorted)
}
