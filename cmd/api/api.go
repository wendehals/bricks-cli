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
	ALL_OPT  = "all"
	ALL_SOPT = "a"
	ALL_ARG  = "-" + ALL_OPT

	SET_LIST_ID_OPT   = "setList"
	SET_LIST_ID_SOPT  = "l"
	SET_LIST_ID_ARG   = "-" + SET_LIST_ID_SOPT + " SET_LIST_ID"
	SET_LIST_ID_USAGE = "The id of a user defined set list"

	PART_LIST_ID_OPT   = "partList"
	PART_LIST_ID_SOPT  = "p"
	PART_LIST_ID_ARG   = "-" + PART_LIST_ID_SOPT + " PART_LIST_ID"
	PART_LIST_ID_USAGE = "The id of a user defined part list"

	SET_NUM_OPT   = "set"
	SET_NUM_SOPT  = "s"
	SET_NUM_ARG   = "-" + SET_NUM_SOPT + " SET_NUM"
	SET_NUM_USAGE = "The set number"
)

var (
	credentialsFile string
	jsonFile        string
	all             bool
	setNum          string
	setListId       uint
	partListId      uint

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
	ApiCmd.AddCommand(detailsCmd)
	ApiCmd.AddCommand(setListsCmd)
	ApiCmd.AddCommand(partListsCmd)
	ApiCmd.AddCommand(userSetsCmd)
	ApiCmd.AddCommand(partsCmd)
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
	if setListId == 0 {
		return fmt.Errorf("please provide a valid list id of a user defined set list")
	}
	return nil
}

func createClient() *http.Client {
	return &http.Client{
		Timeout: time.Second * 5,
	}
}

func createBricksAPI() *api.BricksAPI {
	return api.NewBricksAPI(createClient(), credentials.APIKey)
}

func createUsersAPI() *api.UsersAPI {
	return api.NewUsersAPI(createClient(), credentials)
}
