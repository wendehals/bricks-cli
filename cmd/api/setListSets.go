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

var setListSetsCmd = &cobra.Command{
	Use:   fmt.Sprintf("setListSets %s %s -l LIST_ID", options.CREDENTIALS_ARG, options.JSON_OUTPUT_ARG),
	Short: "Get a list of all the sets in a specific set list",
	Long:  "The setListSets command returns a list of all the sets in a specific set list.",

	DisableFlagsInUseLine: true,

	PreRunE: func(cmd *cobra.Command, args []string) error {
		return checkOptionsSetListSets()
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return executeSetListSets()
	},
}

func init() {
	setListSetsCmd.Flags().UintVarP(&listId, "listId", "l", 0,
		"The list id of a user defined set list")
}

func checkOptionsSetListSets() error {
	if listId == 0 {
		return fmt.Errorf("please provide a list id of a user defined set list")
	}
	return nil
}

func executeSetListSets() error {
	client := http.Client{
		Timeout: time.Second * 5,
	}

	usersAPI := api.NewUsersAPI(&client, credentials)
	sets, err := usersAPI.GetSetListSets(listId)
	if err != nil {
		return err
	}

	if jsonFile == "" {
		jsonFile = fmt.Sprintf("%d_sets.json", listId)
	}

	return model.ExportToJSON(jsonFile, sets)
}
