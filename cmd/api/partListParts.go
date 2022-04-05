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

var (
	listId string

	partListPartsCmd = &cobra.Command{
		Use:   fmt.Sprintf("partListParts %s %s -l LIST_ID", options.CREDENTIALS_ARG, options.JSON_OUTPUT_ARG),
		Short: "Get all parts of a user defined part list.",
		Long:  "The partListParts command returns all parts of a user defined part list.",

		DisableFlagsInUseLine: true,

		PreRunE: func(cmd *cobra.Command, args []string) error {
			return checkOptionsPartListParts()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return executePartListParts()
		},
	}
)

func init() {
	partListPartsCmd.Flags().StringVarP(&listId, "list", "l", "", "The list id of a user defined part list")
}

func checkOptionsPartListParts() error {
	if listId == "" {
		return fmt.Errorf("please provide a list id of a user defined part list")
	}
	return nil
}

func executePartListParts() error {
	client := http.Client{
		Timeout: time.Second * 5,
	}

	usersAPI := api.NewUsersAPI(&client, credentials)
	partLists, err := usersAPI.GetPartListParts(listId)
	if err != nil {
		return err
	}

	if jsonFile == "" {
		jsonFile = listId + "_parts.json"
	}

	return model.ExportToJSON(jsonFile, partLists)
}
