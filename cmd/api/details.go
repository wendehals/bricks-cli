package api

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wendehals/bricks/cmd/options"
	"github.com/wendehals/bricks/model"
)

var detailsCmd = &cobra.Command{
	Use: fmt.Sprintf("details %s %s {%s | %s | %s}", options.CREDENTIALS_ARG,
		options.OUTPUT_FILE_ARG, SET_NUM_ARG, SET_LIST_ID_ARG, PART_LIST_ID_ARG),
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

	if outputFile == "" {
		outputFile = fmt.Sprintf("%s.set", setNum)
	}

	model.Save(set, outputFile)
}

func executeSetListDetails() {
	setList := createUsersAPI().GetSetList(setListId)

	if outputFile == "" {
		outputFile = fmt.Sprintf("%d.setList", setListId)
	}

	model.Save(setList, outputFile)
}

func executePartListDetails() {
	partList := createUsersAPI().GetPartList(partListId)

	if outputFile == "" {
		outputFile = fmt.Sprintf("%d.partList", partListId)
	}

	model.Save(partList, outputFile)
}
