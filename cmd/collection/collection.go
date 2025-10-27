package collection

import (
	"github.com/spf13/cobra"
	"github.com/wendehals/bricks-cli/cmd/options"
)

var (
	outputFile string

	CollectionCmd = &cobra.Command{
		Use:   "collection",
		Short: "Groups all commands for working with part collections",
		Long:  "The collection command groups all sub commands for working with part collections.",

		DisableFlagsInUseLine: true,
	}
)

func init() {
	CollectionCmd.AddCommand(maxCmd)
	CollectionCmd.AddCommand(sortCmd)
	CollectionCmd.AddCommand(subtractCmd)
	CollectionCmd.AddCommand(sumCmd)
	CollectionCmd.AddCommand(mergeCmd)
	CollectionCmd.AddCommand(importCmd)

	CollectionCmd.PersistentFlags().StringVarP(&outputFile, options.OUTPUT_FILE_OPT,
		options.OUTPUT_FILE_SOPT, "", options.OUTPUT_FILE_USAGE)
}
