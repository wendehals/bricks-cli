package api

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wendehals/bricks/cmd/options"
)

const (
	SETS_OPT   = "sets"
	SETS_ARG   = "--" + SETS_OPT
	SETS_USAGE = "get all set lists of the user"

	PARTS_OPT   = "parts"
	PARTS_ARG   = "--" + PARTS_OPT
	PARTS_USAGE = "get all part lists of the user"
)

var (
	sets  bool
	parts bool

	listsCmd = &cobra.Command{
		Use:   fmt.Sprintf("lists %s %s [%s | %s]", options.CREDENTIALS_ARG, options.OUTPUT_FILE_ARG, SETS_ARG, PARTS_ARG),
		Short: "Get the user's set or part lists",
		Long:  "The lists command returns all set or part lists of the user.",

		DisableFlagsInUseLine: true,

		PreRunE: func(cmd *cobra.Command, args []string) error {
			return checkOptionsLists()
		},
		Run: func(cmd *cobra.Command, args []string) {
			executeLists()
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

func executeLists() {
	if sets {
		executeSetLists()
	} else if parts {
		executePartLists()
	}
}

func executeSetLists() {
	setLists := createUsersAPI().GetSetLists()

	if outputFile == "" {
		outputFile = fmt.Sprintf("%s.setLists",
			options.ReplaceIllegalCharsFromFileName(credentials.UserName))
	}

	setLists.Save(outputFile)
}

func executePartLists() {
	partLists := createUsersAPI().GetPartLists()

	if outputFile == "" {
		outputFile = fmt.Sprintf("%s.partLists",
			options.ReplaceIllegalCharsFromFileName(credentials.UserName))
	}

	partLists.Save(outputFile)
}
