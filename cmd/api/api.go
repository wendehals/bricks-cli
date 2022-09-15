package api

import (
	"net/http"
	"time"

	"github.com/spf13/cobra"
	"github.com/wendehals/bricks/api"
	"github.com/wendehals/bricks/cmd/options"
)

const (
	ALL_OPT  = "all"
	ALL_SOPT = "a"
	ALL_ARG  = "-" + ALL_SOPT

	SET_NUM_OPT   = "set"
	SET_NUM_SOPT  = "s"
	SET_NUM_ARG   = "-" + SET_NUM_SOPT + " SET_NUMBER"
	SET_NUM_USAGE = "the set number"

	SET_LIST_ID_OPT   = "setList"
	SET_LIST_ID_SOPT  = "l"
	SET_LIST_ID_ARG   = "-" + SET_LIST_ID_SOPT + " SET_LIST_ID"
	SET_LIST_ID_USAGE = "the id of a user defined set list"

	PART_LIST_ID_OPT   = "partList"
	PART_LIST_ID_SOPT  = "p"
	PART_LIST_ID_ARG   = "-" + PART_LIST_ID_SOPT + " PART_LIST_ID"
	PART_LIST_ID_USAGE = "the id of a user defined part list"
)

var (
	credentialsFile string
	outputFile      string
	all             bool
	setNum          string
	setListId       uint
	partListId      uint

	credentials *api.Credentials

	ApiCmd = &cobra.Command{
		Use:   "api",
		Short: "Groups all commands for the Rebrickable API",
		Long:  "The api command groups all sub commands for the Rebrickable API.",

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
	ApiCmd.AddCommand(listsCmd)
	ApiCmd.AddCommand(setsCmd)
	ApiCmd.AddCommand(partsCmd)

	ApiCmd.PersistentFlags().StringVarP(&credentialsFile, options.CREDENTIALS_OPT, options.CREDENTIALS_SOPT, "credentials.json",
		options.CREDENTIALS_USAGE)
	ApiCmd.PersistentFlags().StringVarP(&outputFile, options.OUTPUT_FILE_OPT, options.OUTPUT_FILE_SOPT, "", options.OUTPUT_FILE_USAGE)
}

func CreateBricksAPI() *api.BricksAPI {
	return api.NewBricksAPI(createClient(), credentials.APIKey, options.Verbose)
}

func CreateUsersAPI() *api.UsersAPI {
	return api.NewUsersAPI(createClient(), credentials, options.Verbose)
}

func createClient() *http.Client {
	return &http.Client{
		Timeout: time.Second * 5,
	}
}
