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
		Use:   fmt.Sprintf("merge %s {--color | --variant} PARTS_FILE", options.OUTPUT_FILE_ARG),
		Short: "Merges the parts of a collection by their color or by their variant",
		Long:  "The command merges all parts of the same type and color or by variants.",

		DisableFlagsInUseLine: true,

		Args: cobra.ExactArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if !color && !variant {
				return fmt.Errorf("at least one option of --color or --variant must be set")
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			executeMerge(args)
		},
	}
)

func init() {
	mergeCmd.Flags().BoolVarP(&color, "color", "", false, "merge by color")
	mergeCmd.Flags().BoolVarP(&variant, "variant", "", false, "merge by variant")
}

func executeMerge(args []string) {
	collection := model.ImportCollection(args[0])

	merged := collection
	if color {
		merged = merged.MergeByColor()
	}

	if variant {
		merged = merged.MergeByVariant()
	}

	if jsonFile == "" {
		jsonFile = options.FileNameFromArgs(args, "_merged.parts")
	}

	model.Save(merged, jsonFile)
}
