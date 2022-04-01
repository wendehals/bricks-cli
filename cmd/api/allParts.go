package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/spf13/cobra"
	"github.com/wendehals/bricks/api"
	"github.com/wendehals/bricks/model"
)

var allPartsCmd = &cobra.Command{
	Use:   fmt.Sprintf("allParts %s %s", credentials_arg, json_output_arg),
	Short: "Get all parts owned by the user",
	Long:  "The allParts command returns a list of all parts owned by the user.",

	DisableFlagsInUseLine: true,

	RunE: func(cmd *cobra.Command, args []string) error {
		return executeGetAllParts()
	},
}

func init() {
	apiCmd.AddCommand(allPartsCmd)
}

func executeGetAllParts() error {
	client := http.Client{
		Timeout: time.Second * 5,
	}

	usersAPI := api.NewUsersAPI(&client, credentials)
	allParts, err := usersAPI.GetAllParts()
	if err != nil {
		return err
	}

	if jsonFile == "" {
		jsonFile = "allParts.json"
	}

	return model.ExportToJSON(jsonFile, allParts)
}
