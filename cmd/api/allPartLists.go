package api

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wendehals/bricks/cmd/options"
	"github.com/wendehals/bricks/model"
)

var allPartListsCmd = &cobra.Command{
	Use:   fmt.Sprintf("allPartLists %s %s", options.CREDENTIALS_ARG, options.JSON_OUTPUT_ARG),
	Short: "Get a list of all the user's part lists",
	Long:  "The allPartLists command returns a list of all part lists of the user.",

	DisableFlagsInUseLine: true,

	RunE: func(cmd *cobra.Command, args []string) error {
		return executeAllPartLists()
	},
}

func executeAllPartLists() error {
	partLists, err := createUsersAPI().GetAllPartLists()
	if err != nil {
		return err
	}

	if jsonFile == "" {
		jsonFile = fmt.Sprintf("%s_all_part_lists.json",
			options.ReplaceIllegalCharsFromFileName(credentials.UserName))
	}

	return model.ExportToJSON(jsonFile, partLists)
}
