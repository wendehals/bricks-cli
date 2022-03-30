package cmd

import (
	"net/http"
	"time"

	"github.com/spf13/cobra"
	"github.com/wendehals/bricks/api"
	"github.com/wendehals/bricks/model"
)

var setsCmd = &cobra.Command{
	Use:   "sets [-c FILE] [-o FILE] [-j FILE]",
	Short: "Get a list of all the user's sets",
	Long: `
The command returns a list of all sets of the user.
`,

	DisableFlagsInUseLine: true,

	RunE: func(cmd *cobra.Command, args []string) error {
		return executeSets()
	},
}

func init() {
	rootCmd.AddCommand(setsCmd)
}

func executeSets() error {
	client := http.Client{
		Timeout: time.Second * 5,
	}

	usersAPI := api.NewUsersAPI(&client, credentials.Username, credentials.Password, credentials.APIKey)
	sets, err := usersAPI.GetSets()
	if err != nil {
		return err
	}

	if jsonFile != "" {
		err := model.ExportToJSON(jsonFile, sets)
		if err != nil {
			return err
		}
	}

	return nil
}
