package cmd

import (
	"fmt"
	"net/http"
	"time"

	"github.com/spf13/cobra"
	"github.com/wendehals/bricks/api"
	"github.com/wendehals/bricks/model"
)

var partListsCmd = &cobra.Command{
	Use:   fmt.Sprintf("partLists %s %s", credentials_arg, json_output_arg),
	Short: "Get a list of all the user's part lists",
	Long:  "The partLists command returns a list of all part lists of the user.",

	DisableFlagsInUseLine: true,

	RunE: func(cmd *cobra.Command, args []string) error {
		return executePartLists()
	},
}

func init() {
	apiCmd.AddCommand(partListsCmd)
}

func executePartLists() error {
	client := http.Client{
		Timeout: time.Second * 5,
	}

	usersAPI := api.NewUsersAPI(&client, credentials)
	partLists, err := usersAPI.GetPartLists()
	if err != nil {
		return err
	}

	if jsonFile == "" {
		jsonFile = "partLists.json"
	}

	return model.ExportToJSON(jsonFile, partLists)
}
