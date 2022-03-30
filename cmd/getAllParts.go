package cmd

import (
	"net/http"
	"time"

	"github.com/spf13/cobra"
	"github.com/wendehals/bricks/api"
	"github.com/wendehals/bricks/model"
)

var getAllPartsCmd = &cobra.Command{
	Use:   "getAllParts [-c FILE] [-o FILE] [-j FILE]",
	Short: "Get all parts owned by the user",
	Long: `
The command returns a list of all parts owned by the user.
`,

	DisableFlagsInUseLine: true,

	RunE: func(cmd *cobra.Command, args []string) error {
		return executeGetAllParts()
	},
}

func init() {
	rootCmd.AddCommand(getAllPartsCmd)
}

func executeGetAllParts() error {
	client := http.Client{
		Timeout: time.Second * 5,
	}

	usersAPI := api.NewUsersAPI(&client, credentials.Username, credentials.Password, credentials.APIKey)
	allParts, err := usersAPI.GetAllParts()
	if err != nil {
		return err
	}

	err = allParts.ExportToHTML(htmlFile)
	if err != nil {
		return err
	}

	if jsonFile != "" {
		err := model.ExportToJSON(jsonFile, allParts)
		if err != nil {
			return err
		}
	}

	return nil
}
