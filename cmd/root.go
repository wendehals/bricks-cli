package cmd

import (
	"github.com/spf13/cobra"
	"github.com/wendehals/bricks/cmd/api"
	"github.com/wendehals/bricks/cmd/collection"
	"github.com/wendehals/bricks/cmd/download"
	"github.com/wendehals/bricks/cmd/options"
)

var RootCmd = &cobra.Command{
	Use:   "bricks",
	Short: "A command line interface to Rebrickable and bricks collections",
	Long: `bricks is a command line interface to the rebrickable.com API. It
enables you to access the Rebrickable database to retrieve for example a list
of all parts of a certain set or user specific part collections.

bricks is also able to merge, sort, add, or subtract collections of parts
to new collections.`,
	Version: "0.6.0",

	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true},
}

func init() {
	RootCmd.AddCommand(api.ApiCmd)
	RootCmd.AddCommand(collection.CollectionCmd)
	RootCmd.AddCommand(buildCmd)
	RootCmd.AddCommand(exportCmd)
	RootCmd.AddCommand(scriptCmd)
	RootCmd.AddCommand(download.DownloadCmd)

	RootCmd.PersistentFlags().BoolVarP(&options.Verbose, "verbose", "", false, "verbose output")
}
