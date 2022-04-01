package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/wendehals/bricks/cmd/api"
	"github.com/wendehals/bricks/cmd/collection"
)

var RootCmd = &cobra.Command{
	Use:   "bricks",
	Short: "A command line interface to Rebrickable and bricks collections",
	Long: `bricks is command line interface to the Rebrickable.com API. It
enables you to access the Rebrickable database to retrieve for example a list
of all bricks of a certain set or user specific brick collections.

bricks is also able to merge, sort, add, or subtract collections of bricks
to new collections.`,
	Version: "0.1.0",

	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true},
}

func init() {
	RootCmd.AddCommand(api.ApiCmd)
	RootCmd.AddCommand(collection.CollectionCmd)
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
