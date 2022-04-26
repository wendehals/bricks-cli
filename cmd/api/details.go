package api

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wendehals/bricks/cmd/options"
	"github.com/wendehals/bricks/model"
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
	RunE: func(cmd *cobra.Command, args []string) error {
		return executeDetails()
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

func executeDetails() error {
	if setNum != "" {
		return executeSetDetails()
	}
	if setListId != 0 {
		return executeSetListDetails()
	}
	if partListId != 0 {
		return executePartListDetails()
	}

	return nil
}

func executeSetDetails() error {
	set, err := createBricksAPI().GetSet(setNum)
	if err != nil {
		return err
	}

	if jsonFile == "" {
		jsonFile = fmt.Sprintf("%s_set.json", setNum)
	}

	return model.ExportToJSON(jsonFile, set)
}

func executeSetListDetails() error {
	setList, err := createUsersAPI().GetSetList(setListId)
	if err != nil {
		return err
	}

	if jsonFile == "" {
		jsonFile = fmt.Sprintf("%d_setList.json", setListId)
	}

	return model.ExportToJSON(jsonFile, setList)
}

func executePartListDetails() error {
	partList, err := createUsersAPI().GetPartList(partListId)
	if err != nil {
		return err
	}

	if jsonFile == "" {
		jsonFile = fmt.Sprintf("%d_partList.json", partListId)
	}

	return model.ExportToJSON(jsonFile, partList)
}
