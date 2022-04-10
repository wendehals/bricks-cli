package api

import (
	"github.com/spf13/cobra"
	"github.com/wendehals/bricks/api"
	"github.com/wendehals/bricks/cmd/options"
)

var (
	credentialsFile string
	jsonFile        string
	setNum          string
	listId          uint

	credentials *api.Credentials

	ApiCmd = &cobra.Command{
		Use:   "api",
		Short: "Groups all commands for the Rebrickable API.",
		Long:  "The api command groups all commands for the Rebrickable API.",

		DisableFlagsInUseLine: true,

		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			var err error
			credentials, err = api.ImportCredentials(credentialsFile)
			return err
		},
	}
)

func init() {
	ApiCmd.AddCommand(allPartsCmd)
	ApiCmd.AddCommand(partListsCmd)
	ApiCmd.AddCommand(partListPartsCmd)
	ApiCmd.AddCommand(setListsCmd)
	ApiCmd.AddCommand(setListSetsCmd)
	ApiCmd.AddCommand(setPartsCmd)
	ApiCmd.AddCommand(setCmd)
	ApiCmd.AddCommand(setsCmd)

	ApiCmd.PersistentFlags().StringVarP(&credentialsFile, options.CREDENTIALS_OPT, options.CREDENTIALS_SOPT, "credentials.json",
		options.CREDENTIALS_USAGE)
	ApiCmd.PersistentFlags().StringVarP(&jsonFile, options.JSON_OUTPUT_OPT, options.JSON_OUTPUT_SOPT, "", options.JSON_OUTPUT_USAGE)
}
