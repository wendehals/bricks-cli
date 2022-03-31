/*
Copyright © 2022 Lothar Wendehals
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/wendehals/bricks/utils"
)

var (
	credentialsFile string
	htmlFile        string
	jsonFile        string

	credentials *utils.Credentials

	rootCmd = &cobra.Command{
		Use:   "bricks",
		Short: "A command line interface to Rebrickable and bricks collections",
		Long: `bricks is command line interface to the Rebrickable.com API. It
enables you to access the Rebrickable database to retrieve for example a list
of all bricks of a certain set or user specific brick collections.

bricks is also able to merge, sort, add, or subtract collections of bricks
to new collections.`,
		Version: "0.0.1-alpha1",

		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd: true},

		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return readCredentials()
		},
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&credentialsFile, "credentials", "c", "credentials.json",
		"A JSON file containing the Rebrickable credentials")
	rootCmd.PersistentFlags().StringVarP(&htmlFile, "html", "o", "output.html",
		"A name for the HTML output file")
	rootCmd.PersistentFlags().StringVarP(&jsonFile, "json", "j", "",
		"A name for the JSON output file")
}

func readCredentials() error {
	var err error
	credentials, err = utils.NewCredentials(credentialsFile)
	if err != nil {
		return fmt.Errorf("please provide a valid JSON file containing the Rebrickable credentials:\n   %s", err)
	}

	return nil
}
