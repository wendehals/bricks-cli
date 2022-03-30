package cmd

import (
	"net/http"
	"time"

	"github.com/spf13/cobra"
	"github.com/wendehals/bricks/api"
	"github.com/wendehals/bricks/model"
)

var setListsCmd = &cobra.Command{
	Use:   "setLists [-c FILE] [-o FILE] [-j FILE]",
	Short: "Get a list of all the user's set lists",
	Long: `
The command returns a list of all set lists of the user.
`,

	DisableFlagsInUseLine: true,

	RunE: func(cmd *cobra.Command, args []string) error {
		return excuteSetLists()
	},
}

func init() {
	rootCmd.AddCommand(setListsCmd)
}

func excuteSetLists() error {
	client := http.Client{
		Timeout: time.Second * 5,
	}

	usersAPI := api.NewUsersAPI(&client, credentials.Username, credentials.Password, credentials.APIKey)
	setLists, err := usersAPI.GetSetLists()
	if err != nil {
		return err
	}

	if jsonFile != "" {
		model.ExportToJSON(jsonFile, setLists)
	}

	return nil
}
