package cmd

import (
	"net/http"
	"time"

	"github.com/spf13/cobra"
	"github.com/wendehals/bricks/api"
	"github.com/wendehals/bricks/model"
)

var partListsCmd = &cobra.Command{
	Use:   "partLists [-c FILE] [-o FILE] [-j FILE]",
	Short: "Get a list of all the user's part lists",
	Long: `
The command returns a list of all part lists of the user.
`,

	DisableFlagsInUseLine: true,

	RunE: func(cmd *cobra.Command, args []string) error {
		return executePartLists()
	},
}

func init() {
	rootCmd.AddCommand(partListsCmd)
}

func executePartLists() error {
	client := http.Client{
		Timeout: time.Second * 5,
	}

	usersAPI := api.NewUsersAPI(&client, credentials.Username, credentials.Password, credentials.APIKey)
	partLists, err := usersAPI.GetPartLists()
	if err != nil {
		return err
	}

	if jsonFile != "" {
		err := model.ExportToJSON(jsonFile, partLists)
		if err != nil {
			return err
		}
	}

	return nil
}
