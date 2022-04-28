package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/wendehals/bricks/api"
	"github.com/wendehals/bricks/cmd/options"
	"github.com/wendehals/bricks/scripting"
)

var scriptCmd = &cobra.Command{
	Use:   fmt.Sprintf("script %s", options.CREDENTIALS_ARG),
	Short: "Executes a bricks script",
	Long:  "The script command exxecutes a bricks script.",

	DisableFlagsInUseLine: true,

	Args: cobra.ExactArgs(1),

	PreRunE: func(cmd *cobra.Command, args []string) error {
		var err error
		credentials, err = api.ImportCredentials(credentialsFile)
		return err
	},
	Run: func(cmd *cobra.Command, args []string) {
		executeScript(args)
	},
}

func init() {
	scriptCmd.Flags().StringVarP(&credentialsFile, options.CREDENTIALS_OPT, options.CREDENTIALS_SOPT,
		"credentials.json", options.CREDENTIALS_USAGE)
}

func executeScript(args []string) {
	log.Printf("Executing script '%s'\n", args[0])

	bricksScript := scripting.NewBricksScript(credentials, args[0], options.Verbose)
	bricksScript.Execute()
}
