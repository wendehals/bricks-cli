package collection

import (
	"github.com/spf13/cobra"
	"github.com/wendehals/bricks/cmd/options"
)

var (
	jsonFile string

	CollectionCmd = &cobra.Command{
		Use:   "collection",
		Short: "Groups all commands for working with part collections.",
		Long:  "The collection command groups all commands for working with part collections.",

		DisableFlagsInUseLine: true,
	}
)

func init() {
	CollectionCmd.AddCommand(maxCmd)
	CollectionCmd.AddCommand(mergeCmd)
	CollectionCmd.AddCommand(sortCmd)
	CollectionCmd.AddCommand(subtractCmd)
	CollectionCmd.AddCommand(sumCmd)

	CollectionCmd.PersistentFlags().StringVarP(&jsonFile, options.JSON_OUTPUT_OPT,
		options.JSON_OUTPUT_SOPT, "", options.JSON_OUTPUT_USAGE)
}
