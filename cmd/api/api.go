package api

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wendehals/bricks/api"
	"github.com/wendehals/bricks/cmd"
)

const (
	credentials_opt   = "credentials"
	credentials_sopt  = "c"
	credentials_arg   = "[-" + credentials_sopt + " CREDENTIAL_FILE]"
	credentials_usage = "A JSON file containing the Rebrickable credentials"

	json_output_opt   = "output"
	json_output_sopt  = "o"
	json_output_arg   = "[-" + json_output_sopt + " JSON_FILE]"
	json_output_usage = "A name for the JSON output file"
)

var (
	credentialsFile string
	jsonFile        string

	credentials *api.Credentials

	apiCmd = &cobra.Command{
		Use:   "api",
		Short: "Groups all commands for the Rebrickable API.",
		Long:  "The api command groups all commands for the Rebrickable API.",

		DisableFlagsInUseLine: true,

		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return readCredentials()
		},
	}
)

func init() {
	cmd.RootCmd.AddCommand(apiCmd)

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
