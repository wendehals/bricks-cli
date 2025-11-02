package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/wendehals/bricks-cli/api"
	"github.com/wendehals/bricks-cli/cmd/options"
	"github.com/wendehals/bricks-cli/scripting"
	"github.com/wendehals/bricks-cli/utils"
)

var scriptCmd = &cobra.Command{
	Use:   fmt.Sprintf("script %s SCRIPT_FILE", options.CREDENTIALS_ARG),
	Short: "Executes a bricks-cli script",
	Long:  "The script command executes a bricks-cli script.",

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
		utils.CredentialsDefaultPath(), options.CREDENTIALS_USAGE)
}

func executeScript(args []string) {
	log.Printf("Executing script '%s'", args[0])

	scriptContent, err := os.ReadFile(args[0])
	if err != nil {
		log.Fatalf("reading script from file '%s' failed: %s", args[0], err.Error())
	}

	bricksScript := scripting.NewBricksScript(credentials, string(scriptContent), options.Verbose)
	bricksScript.Execute()
}
