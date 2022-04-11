package api

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wendehals/bricks/api"
	"github.com/wendehals/bricks/cmd/options"
	"github.com/wendehals/bricks/model"
)

var setCmd = &cobra.Command{
	Use:   fmt.Sprintf("set %s %s -n SET_NUMBER", options.CREDENTIALS_ARG, options.JSON_OUTPUT_ARG),
	Short: "Get details for a specific set",
	Long:  "The set command returns details for a specific set.",

	DisableFlagsInUseLine: true,

	PreRunE: func(cmd *cobra.Command, args []string) error {
		return checkOptionsSet()
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return executeSet()
	},
}

func init() {
	setCmd.Flags().StringVarP(&setNum, "set", "n", "", "A set number")
}

func checkOptionsSet() error {
	if setNum == "" {
		return fmt.Errorf("please provide a set number")
	}
	return nil
}
func executeSet() error {
	bricksAPI := api.NewBricksAPI(createClient(), credentials.APIKey)
	set, err := bricksAPI.GetSet(setNum)
	if err != nil {
		return err
	}

	if jsonFile == "" {
		jsonFile = fmt.Sprintf("%s_set.json", setNum)
	}

	return model.ExportToJSON(jsonFile, set)
}
