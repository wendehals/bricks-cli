package cmd

import (
	"github.com/spf13/cobra"
	"github.com/wendehals/bricks/cmd/api"
	"github.com/wendehals/bricks/cmd/collection"
)

var RootCmd = &cobra.Command{
	Use:   "bricks",
	Short: "A command line interface to Rebrickable and bricks collections",
	Long: `bricks is a command line interface to the rebrickable.com API. It
enables you to access the Rebrickable database to retrieve for example a list
of all bricks of a certain set or user specific brick collections.

bricks is also able to merge, sort, add, or subtract collections of parts
to new collections.`,
	Version: "0.3.0",

	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true},
}

func init() {
	RootCmd.AddCommand(api.ApiCmd)
	RootCmd.AddCommand(collection.CollectionCmd)
	RootCmd.AddCommand(exportCmd)
}
