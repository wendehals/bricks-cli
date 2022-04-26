package api

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/wendehals/bricks/cmd/options"
	"github.com/wendehals/bricks/model"
)

const (
	SETS_OPT   = "sets"
	SETS_ARG   = "--" + SETS_OPT
	SETS_USAGE = "Get all set lists of the user"

	PARTS_OPT   = "parts"
	PARTS_ARG   = "--" + PARTS_OPT
	PARTS_USAGE = "Get all part lists of the user"
)

var (
	sets  bool
	parts bool

	listsCmd = &cobra.Command{
		Use:   fmt.Sprintf("lists %s %s [%s | %s]", options.CREDENTIALS_ARG, options.JSON_OUTPUT_ARG, SETS_ARG, PARTS_ARG),
		Short: "Get the user's set or part lists",
		Long:  "The lists command returns all set or part lists of the user.",

		DisableFlagsInUseLine: true,

		PreRunE: func(cmd *cobra.Command, args []string) error {
			return checkOptionsLists()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return executeLists()
		},
	}
)

func init() {
	listsCmd.Flags().BoolVarP(&sets, SETS_OPT, "", false, SETS_USAGE)
	listsCmd.Flags().BoolVarP(&parts, PARTS_OPT, "", false, PARTS_USAGE)
}

func checkOptionsLists() error {
	if (!sets && !parts) || (sets && parts) {
		return fmt.Errorf("please provide either %s or %s", SETS_ARG, PARTS_ARG)
	}

	return nil
}

func executeLists() error {
	if sets {
		return executeSetLists()
	}
	if parts {
		return executePartLists()
	}

	return nil
}

func executeSetLists() error {
	log.Printf("Retrieving set lists of user %s\n", credentials.UserName)
	setLists, err := createUsersAPI().GetSetLists()
	if err != nil {
		return err
	}

	if jsonFile == "" {
		jsonFile = fmt.Sprintf("%s_setLists.json",
			options.ReplaceIllegalCharsFromFileName(credentials.UserName))
	}

	return model.ExportToJSON(jsonFile, setLists)
}

func executePartLists() error {
	log.Printf("Retrieving part lists of user %s\n", credentials.UserName)
	partLists, err := createUsersAPI().GetPartLists()
	if err != nil {
		return err
	}

	if jsonFile == "" {
		jsonFile = fmt.Sprintf("%s_partLists.json",
			options.ReplaceIllegalCharsFromFileName(credentials.UserName))
	}

	return model.ExportToJSON(jsonFile, partLists)
}
