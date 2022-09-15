package api

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wendehals/bricks/cmd/options"
	"github.com/wendehals/bricks/model"
)

var setsCmd = &cobra.Command{
	Use:   fmt.Sprintf("sets %s %s {%s | %s}", options.CREDENTIALS_ARG, options.OUTPUT_FILE_ARG, ALL_ARG, SET_LIST_ID_ARG),
	Short: "Get details about sets owned by the user",
	Long:  "The sets command returns details about all sets owned by the user or about sets of a certain set list.",

	DisableFlagsInUseLine: true,

	PreRunE: func(cmd *cobra.Command, args []string) error {
		return checkOptionsSets()
	},
	Run: func(cmd *cobra.Command, args []string) {
		executeSets()
	},
}

func init() {
	setsCmd.Flags().BoolVarP(&all, ALL_OPT, ALL_SOPT, false, "get all sets")
	setsCmd.Flags().UintVarP(&setListId, SET_LIST_ID_OPT, SET_LIST_ID_SOPT, 0, SET_LIST_ID_USAGE)
}

func checkOptionsSets() error {
	if (setListId == 0 && !all) || (setListId != 0 && all) {
		return fmt.Errorf("please provide either %s or %s option", ALL_ARG, SET_LIST_ID_ARG)
	}

	return nil
}

func executeSets() {
	if all {
		executeAllSets()
	} else if setListId != 0 {
		executeSetListSets()
	}
}

func executeAllSets() {
	sets := CreateUsersAPI().GetSets()

	if outputFile == "" {
		outputFile = fmt.Sprintf("%s.sets",
			options.ReplaceIllegalCharsFromFileName(credentials.UserName))
	}

	model.Save(sets, outputFile)
}

func executeSetListSets() {
	usersAPI := CreateUsersAPI()
	sets := usersAPI.GetSetListSets(setListId)
	setList := usersAPI.GetSetList(setListId)
	sets.Name = setList.Name

	if outputFile == "" {
		outputFile = fmt.Sprintf("%d.sets", setListId)
	}

	model.Save(sets, outputFile)
}
