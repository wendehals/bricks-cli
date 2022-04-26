package api

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wendehals/bricks/cmd/options"
	"github.com/wendehals/bricks/model"
)

var lostPartsCmd = &cobra.Command{
	Use:   fmt.Sprintf("lostParts %s %s", options.CREDENTIALS_ARG, options.JSON_OUTPUT_ARG),
	Short: "Get all lost parts of the user",
	Long:  "The lostParts command returns all lost parts of the user.",

	DisableFlagsInUseLine: true,

	Run: func(cmd *cobra.Command, args []string) {
		executeLostParts()
	},
}

func executeLostParts() {
	lostParts := createUsersAPI().GetLostParts()

	if jsonFile == "" {
		jsonFile = fmt.Sprintf("%s_lostParts.json",
			options.ReplaceIllegalCharsFromFileName(credentials.UserName))
	}

	model.ExportToJSON(jsonFile, lostParts)
}
