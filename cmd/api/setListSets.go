package api

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wendehals/bricks/cmd/options"
	"github.com/wendehals/bricks/model"
)

var setListSetsCmd = &cobra.Command{
	Use:   fmt.Sprintf("setListSets %s %s %s", options.CREDENTIALS_ARG, options.JSON_OUTPUT_ARG, LIST_ID_ARG),
	Short: "Get a list of all sets in a specific set list",
	Long:  "The setListSets command returns a list of all sets in a specific set list.",

	DisableFlagsInUseLine: true,

	PreRunE: func(cmd *cobra.Command, args []string) error {
		return checkOptionListId()
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return executeSetListSets()
	},
}

func init() {
	setListSetsCmd.Flags().UintVarP(&listId, LIST_ID_OPT, LIST_ID_SOPT, 0, LIST_ID_USAGE)
}

func executeSetListSets() error {
	sets, err := createUsersAPI().GetSetListSets(listId)
	if err != nil {
		return err
	}

	if jsonFile == "" {
		jsonFile = fmt.Sprintf("%d_sets.json", listId)
	}

	return model.ExportToJSON(jsonFile, sets)
}
