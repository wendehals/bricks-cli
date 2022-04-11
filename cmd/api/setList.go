package api

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wendehals/bricks/cmd/options"
	"github.com/wendehals/bricks/model"
)

var setListCmd = &cobra.Command{
	Use:   fmt.Sprintf("setList %s %s %s", options.CREDENTIALS_ARG, options.JSON_OUTPUT_ARG, LIST_ID_ARG),
	Short: "Get details about a specific set list",
	Long:  "The setList command returns details about a specific set list.",

	DisableFlagsInUseLine: true,

	PreRunE: func(cmd *cobra.Command, args []string) error {
		return checkOptionListId()
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return executeSetList()
	},
}

func init() {
	setListCmd.Flags().UintVarP(&listId, LIST_ID_OPT, LIST_ID_SOPT, 0, LIST_ID_USAGE)
}

func executeSetList() error {
	setList, err := createUsersAPI().GetSetList(listId)
	if err != nil {
		return err
	}

	if jsonFile == "" {
		jsonFile = fmt.Sprintf("%d_set_list.json", listId)
	}

	return model.ExportToJSON(jsonFile, setList)
}
