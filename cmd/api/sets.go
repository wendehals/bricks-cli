package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/spf13/cobra"
	"github.com/wendehals/bricks/api"
	"github.com/wendehals/bricks/model"
)

var setsCmd = &cobra.Command{
	Use:   fmt.Sprintf("sets %s %s", credentials_arg, json_output_arg),
	Short: "Get a list of all the user's sets",
	Long:  "The sets command returns a list of all sets of the user.",

	DisableFlagsInUseLine: true,

	RunE: func(cmd *cobra.Command, args []string) error {
		return executeSets()
	},
}

func init() {
	apiCmd.AddCommand(setsCmd)
}

func executeSets() error {
	client := http.Client{
		Timeout: time.Second * 5,
	}

	usersAPI := api.NewUsersAPI(&client, credentials)
	sets, err := usersAPI.GetSets()
	if err != nil {
		return err
	}

	if jsonFile == "" {
		jsonFile = "sets.json"
	}

	return model.ExportToJSON(jsonFile, sets)
}
