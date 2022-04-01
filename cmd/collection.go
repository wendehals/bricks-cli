package cmd

import (
	"github.com/spf13/cobra"
)

var collectionCmd = &cobra.Command{
	Use:   "collection",
	Short: "Groups all commands for working with bricks collections.",
	Long:  "The collection command groups all commands for working with bricks collections.",

	DisableFlagsInUseLine: true,
}

func init() {
	rootCmd.AddCommand(collectionCmd)
}
