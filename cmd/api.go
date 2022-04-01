package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wendehals/bricks/api"
)

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Groups all commands for the Rebrickable API.",
	Long:  "The api command groups all commands for the Rebrickable API.",

	DisableFlagsInUseLine: true,

	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return readCredentials()
	},
}

func init() {
	rootCmd.AddCommand(apiCmd)

	apiCmd.PersistentFlags().StringVarP(&credentialsFile, credentials_opt, credentials_sopt, "credentials.json",
		credentials_usage)
	apiCmd.PersistentFlags().StringVarP(&jsonFile, json_output_opt, json_output_sopt, "", json_output_usage)
}

func readCredentials() error {
	var err error
	credentials, err = api.NewCredentials(credentialsFile)
	if err != nil {
		return fmt.Errorf("please provide a valid JSON file containing the Rebrickable credentials:\n   %s", err)
	}

	return nil
}
