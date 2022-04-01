package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
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

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
