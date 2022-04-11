package api

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wendehals/bricks/cmd/options"
	"github.com/wendehals/bricks/model"
)

var setCmd = &cobra.Command{
	Use:   fmt.Sprintf("set %s %s %s", options.CREDENTIALS_ARG, options.JSON_OUTPUT_ARG, SET_NUM_ARG),
	Short: "Get details about a specific set",
	Long:  "The set command returns details about a specific set.",

	DisableFlagsInUseLine: true,

	PreRunE: func(cmd *cobra.Command, args []string) error {
		return checkOptionSetNum()
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return executeSet()
	},
}

func init() {
	setCmd.Flags().StringVarP(&setNum, SET_NUM_OPT, SET_NUM_SOPT, "", SET_NUM_USAGE)
}

func executeSet() error {
	set, err := createBricksAPI().GetSet(setNum)
	if err != nil {
		return err
	}

	if jsonFile == "" {
		jsonFile = fmt.Sprintf("%s_set.json", setNum)
	}

	return model.ExportToJSON(jsonFile, set)
}
