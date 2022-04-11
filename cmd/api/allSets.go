package api

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wendehals/bricks/cmd/options"
	"github.com/wendehals/bricks/model"
)

var allSetsCmd = &cobra.Command{
	Use:   fmt.Sprintf("allSets %s %s", options.CREDENTIALS_ARG, options.JSON_OUTPUT_ARG),
	Short: "Get a list of all sets owned by the user",
	Long:  "The allSets command returns a list of all sets owned by the user.",

	DisableFlagsInUseLine: true,

	RunE: func(cmd *cobra.Command, args []string) error {
		return executeAllSets()
	},
}

func executeAllSets() error {
	sets, err := createUsersAPI().GetAllSets()
	if err != nil {
		return err
	}

	if jsonFile == "" {
		jsonFile = fmt.Sprintf("%s_all_sets.json",
			options.ReplaceIllegalCharsFromFileName(credentials.UserName))
	}

	return model.ExportToJSON(jsonFile, sets)
}
