package collection

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wendehals/bricks/cmd/options"
	"github.com/wendehals/bricks/model"
)

var (
	color   bool
	variant bool

	mergeCmd = &cobra.Command{
		Use:   fmt.Sprintf("merge %s {--color | --variant} JSON_FILE", options.JSON_OUTPUT_ARG),
		Short: "Merges the parts of a collection by their color or by their variant.",
		Long:  "The command merges all parts of the same type and color or by variants.",

		DisableFlagsInUseLine: true,

		Args: cobra.ExactArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if !color && !variant {
				return fmt.Errorf("at least one option of --color or --variant must be set")
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return executeMerge(args)
		},
	}
)

func init() {
	mergeCmd.Flags().BoolVarP(&color, "color", "", false, "Merge by color")
	mergeCmd.Flags().BoolVarP(&variant, "variant", "", false, "Merge by variant")
}

func executeMerge(args []string) error {
	collection, err := model.ImportCollection(args[0])
	if err != nil {
		return err
	}

	merged := collection
	if color {
		merged = merged.MergeByColor()
	}

	if variant {
		merged = merged.MergeByVariant()
	}

	if jsonFile == "" {
		jsonFile = options.FileNameFromArgs(args, "_merged_parts.json")
	}

	return model.ExportToJSON(jsonFile, merged)
}
