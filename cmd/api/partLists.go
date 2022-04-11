package api

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wendehals/bricks/api"
	"github.com/wendehals/bricks/cmd/options"
	"github.com/wendehals/bricks/model"
)

var partListsCmd = &cobra.Command{
	Use:   fmt.Sprintf("partLists %s %s", options.CREDENTIALS_ARG, options.JSON_OUTPUT_ARG),
	Short: "Get a list of all the user's part lists",
	Long:  "The partLists command returns a list of all part lists of the user.",

	DisableFlagsInUseLine: true,

	RunE: func(cmd *cobra.Command, args []string) error {
		return executePartLists()
	},
}

func executePartLists() error {
	usersAPI := api.NewUsersAPI(createClient(), credentials)
	partLists, err := usersAPI.GetPartLists()
	if err != nil {
		return err
	}

	if jsonFile == "" {
		jsonFile = "partLists.json"
	}

	return model.ExportToJSON(jsonFile, partLists)
}
