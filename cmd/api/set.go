package api

import (
	"fmt"
	"net/http"
	"time"

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

	RunE: func(cmd *cobra.Command, args []string) error {
		return executeSet()
	},
}

func init() {
	setCmd.Flags().StringVarP(&setNum, "set", "n", "", "A set number")
}

func executeSet() error {
	client := http.Client{
		Timeout: time.Second * 5,
	}

	bricksAPI := api.NewBricksAPI(&client, credentials.APIKey)
	set, err := bricksAPI.GetSet(setNum)
	if err != nil {
		return err
	}

	if jsonFile == "" {
		jsonFile = fmt.Sprintf("%s_set.json", setNum)
	}

	return model.ExportToJSON(jsonFile, set)
}
