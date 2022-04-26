package api

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wendehals/bricks/cmd/options"
	"github.com/wendehals/bricks/model"
)

var userSetsCmd = &cobra.Command{
	Use:   fmt.Sprintf("userSets %s %s {%s | %s}", options.CREDENTIALS_ARG, options.JSON_OUTPUT_ARG, ALL_ARG, SET_LIST_ID_ARG),
	Short: "Get details about sets owned by the user",
	Long:  "The userSets command returns details about all sets owned by the user or about sets of a certain set list.",

	DisableFlagsInUseLine: true,

	PreRunE: func(cmd *cobra.Command, args []string) error {
		return checkOptionsUserSets()
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return executeUserSets()
	},
}

func init() {
	userSetsCmd.Flags().BoolVarP(&all, ALL_OPT, ALL_SOPT, false, "Get all sets")
	userSetsCmd.Flags().UintVarP(&setListId, SET_LIST_ID_OPT, SET_LIST_ID_SOPT, 0, SET_LIST_ID_USAGE)
}

func checkOptionsUserSets() error {
	if setListId == 0 && !all {
		return fmt.Errorf("please provide either -%s or -%s option", ALL_SOPT, SET_LIST_ID_OPT)
	}
	if setListId != 0 && all {
		return fmt.Errorf("please provide only one option of -%s or -%s", ALL_SOPT, SET_LIST_ID_OPT)
	}

	return nil
}
func executeUserSets() error {
	if all {
		return executeAllSets()
	}
	if setListId != 0 {
		return executeSetListSets()
	}

	return nil
}

func executeAllSets() error {
	sets, err := createUsersAPI().GetSets()
	if err != nil {
		return err
	}

	if jsonFile == "" {
		jsonFile = fmt.Sprintf("%s_userSets.json",
			options.ReplaceIllegalCharsFromFileName(credentials.UserName))
	}

	return model.ExportToJSON(jsonFile, sets)
}

func executeSetListSets() error {
	sets, err := createUsersAPI().GetSetListSets(setListId)
	if err != nil {
		return err
	}

	setList, err := createUsersAPI().GetSetList(setListId)
	if err != nil {
		return err
	}
	sets.Name = setList.Name

	if jsonFile == "" {
		jsonFile = fmt.Sprintf("%d_userSets.json", setListId)
	}

	return model.ExportToJSON(jsonFile, sets)
}
