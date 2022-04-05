package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/spf13/cobra"
	"github.com/wendehals/bricks/api"
	"github.com/wendehals/bricks/cmd/options"
	"github.com/wendehals/bricks/model"
)

var setListsCmd = &cobra.Command{
	Use:   fmt.Sprintf("setLists %s %s", options.CREDENTIALS_ARG, options.JSON_OUTPUT_ARG),
	Short: "Get a list of all the user's set lists",
	Long:  "The setLists command returns a list of all set lists of the user.",

	DisableFlagsInUseLine: true,

	RunE: func(cmd *cobra.Command, args []string) error {
		return executeSetLists()
	},
}

func executeSetLists() error {
	client := http.Client{
		Timeout: time.Second * 5,
	}

	usersAPI := api.NewUsersAPI(&client, credentials)
	setLists, err := usersAPI.GetSetLists()
	if err != nil {
		return err
	}

	if jsonFile == "" {
		jsonFile = "setLists.json"
	}

	return model.ExportToJSON(jsonFile, setLists)
}
