package api

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wendehals/bricks/cmd/options"
)

var detailsCmd = &cobra.Command{
	Use: fmt.Sprintf("details %s %s {%s | %s | %s}", options.CREDENTIALS_ARG,
		options.JSON_OUTPUT_ARG, SET_NUM_ARG, SET_LIST_ID_ARG, PART_LIST_ID_ARG),
	Short: "Get details about a certain set, set list, or part list",
	Long:  "The details command returns details about a certain set, set list, or part list.",

	DisableFlagsInUseLine: true,

	PreRunE: func(cmd *cobra.Command, args []string) error {
		return checkOptionsDetails()
	},
	Run: func(cmd *cobra.Command, args []string) {
		executeDetails()
	},
}

func init() {
	detailsCmd.Flags().StringVarP(&setNum, SET_NUM_OPT, SET_NUM_SOPT, "", SET_NUM_USAGE)
	detailsCmd.Flags().UintVarP(&setListId, SET_LIST_ID_OPT, SET_LIST_ID_SOPT, 0, SET_LIST_ID_USAGE)
	detailsCmd.Flags().UintVarP(&partListId, PART_LIST_ID_OPT, PART_LIST_ID_SOPT, 0, PART_LIST_ID_USAGE)
}

func checkOptionsDetails() error {
	optionsProvided := 0
	if setNum != "" {
		optionsProvided++
	}
	if setListId != 0 {
		optionsProvided++
	}
	if partListId != 0 {
		optionsProvided++
	}

	if optionsProvided < 1 || optionsProvided > 1 {
		return fmt.Errorf("please provide exactly one option of %s, %s, or %s",
			SET_NUM_ARG, SET_LIST_ID_ARG, PART_LIST_ID_ARG)
	}

	return nil
}

func executeDetails() {
	if setNum != "" {
		executeSetDetails()
	} else if setListId != 0 {
		executeSetListDetails()
	} else if partListId != 0 {
		executePartListDetails()
	}
}

func executeSetDetails() {
	set := createBricksAPI().GetSet(setNum)

	if jsonFile == "" {
		jsonFile = fmt.Sprintf("%s_set.json", setNum)
	}

	set.Save(jsonFile)
}

func executeSetListDetails() {
	setList := createUsersAPI().GetSetList(setListId)

	if jsonFile == "" {
		jsonFile = fmt.Sprintf("%d_setList.json", setListId)
	}

	setList.Save(jsonFile)
}

func executePartListDetails() {
	partList := createUsersAPI().GetPartList(partListId)

	if jsonFile == "" {
		jsonFile = fmt.Sprintf("%d_partList.json", partListId)
	}

	partList.Save(jsonFile)
}
