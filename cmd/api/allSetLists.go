package api

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wendehals/bricks/cmd/options"
	"github.com/wendehals/bricks/model"
)

var allSetListsCmd = &cobra.Command{
	Use:   fmt.Sprintf("allSetLists %s %s", options.CREDENTIALS_ARG, options.JSON_OUTPUT_ARG),
	Short: "Get a list of all the user's set lists",
	Long:  "The allSetLists command returns a list of all set lists of the user.",

	DisableFlagsInUseLine: true,

	RunE: func(cmd *cobra.Command, args []string) error {
		return executeAllSetLists()
	},
}

func executeAllSetLists() error {
	setLists, err := createUsersAPI().GetAllSetLists()
	if err != nil {
		return err
	}

	if jsonFile == "" {
		jsonFile = fmt.Sprintf("%s_all_set_lists.json",
			options.ReplaceIllegalCharsFromFileName(credentials.UserName))
	}

	return model.ExportToJSON(jsonFile, setLists)
}
