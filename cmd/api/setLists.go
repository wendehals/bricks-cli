package api

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
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
	log.Printf("Retrieving set lists of user %s\n", credentials.UserName)
	setLists, err := createUsersAPI().GetSetLists()
	if err != nil {
		return err
	}

	if jsonFile == "" {
		jsonFile = fmt.Sprintf("%s_setLists.json",
			options.ReplaceIllegalCharsFromFileName(credentials.UserName))
	}

	return model.ExportToJSON(jsonFile, setLists)
}
