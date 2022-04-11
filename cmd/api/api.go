package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/spf13/cobra"
	"github.com/wendehals/bricks/api"
	"github.com/wendehals/bricks/cmd/options"
)

const (
	LIST_ID_OPT   = "list"
	LIST_ID_SOPT  = "l"
	LIST_ID_ARG   = "[-" + LIST_ID_SOPT + " LIST_ID]"
	LIST_ID_USAGE = "The list id of a user defined set list"

	SET_NUM_OPT   = "set"
	SET_NUM_SOPT  = "s"
	SET_NUM_ARG   = "[-" + SET_NUM_SOPT + " SET_NUM]"
	SET_NUM_USAGE = "The set number"
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
	ApiCmd.AddCommand(allPartListsCmd)
	ApiCmd.AddCommand(allPartsCmd)
	ApiCmd.AddCommand(allSetListsCmd)
	ApiCmd.AddCommand(allSetsCmd)

	ApiCmd.AddCommand(partListPartsCmd)
	ApiCmd.AddCommand(setListSetsCmd)
	ApiCmd.AddCommand(setPartsCmd)
	ApiCmd.AddCommand(setCmd)
	ApiCmd.AddCommand(setListCmd)

	ApiCmd.PersistentFlags().StringVarP(&credentialsFile, options.CREDENTIALS_OPT, options.CREDENTIALS_SOPT, "credentials.json",
		options.CREDENTIALS_USAGE)
	ApiCmd.PersistentFlags().StringVarP(&jsonFile, options.JSON_OUTPUT_OPT, options.JSON_OUTPUT_SOPT, "", options.JSON_OUTPUT_USAGE)
}

func checkOptionSetNum() error {
	if setNum == "" {
		return fmt.Errorf("please provide a valid set number")
	}
	return nil
}

func checkOptionListId() error {
	if listId == 0 {
		return fmt.Errorf("please provide a valid list id of a user defined set list")
	}
	return nil
}

func createClient() *http.Client {
	client := http.Client{
		Timeout: time.Second * 5,
	}
	return &client
}

func createBricksAPI() *api.BricksAPI {
	return api.NewBricksAPI(createClient(), credentials.APIKey)
}

func createUsersAPI() *api.UsersAPI {
	return api.NewUsersAPI(createClient(), credentials)
}
