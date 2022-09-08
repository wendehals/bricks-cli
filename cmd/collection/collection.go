package collection

import (
	"github.com/spf13/cobra"
	"github.com/wendehals/bricks/cmd/options"
)

var (
	jsonFile string

	CollectionCmd = &cobra.Command{
		Use:   "collection",
		Short: "Groups all commands for working with part collections",
		Long:  "The collection command groups all sub commands for working with part collections.",

		DisableFlagsInUseLine: true,
	}
)

func init() {
	CollectionCmd.AddCommand(maxCmd)
	CollectionCmd.AddCommand(mergeCmd)
	CollectionCmd.AddCommand(sortCmd)
	CollectionCmd.AddCommand(subtractCmd)
	CollectionCmd.AddCommand(sumCmd)

	CollectionCmd.PersistentFlags().StringVarP(&jsonFile, options.OUTPUT_FILE_OPT,
		options.OUTPUT_FILE_SOPT, "", options.OUTPUT_FILE_USAGE)
}
