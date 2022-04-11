package api

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wendehals/bricks/cmd/options"
	"github.com/wendehals/bricks/model"
)

var allPartsCmd = &cobra.Command{
	Use:   fmt.Sprintf("allParts %s %s", options.CREDENTIALS_ARG, options.JSON_OUTPUT_ARG),
	Short: "Get a list of all parts owned by the user",
	Long:  "The allParts command returns a list of all parts owned by the user.",

	DisableFlagsInUseLine: true,

	RunE: func(cmd *cobra.Command, args []string) error {
		return executeGetAllParts()
	},
}

func executeGetAllParts() error {
	allParts, err := createUsersAPI().GetAllParts()
	if err != nil {
		return err
	}

	if jsonFile == "" {
		jsonFile = fmt.Sprintf("%s_all_parts.json",
			options.ReplaceIllegalCharsFromFileName(credentials.UserName))
	}

	return model.ExportToJSON(jsonFile, allParts)
}
