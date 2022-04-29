package api

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/wendehals/bricks/api"
	"github.com/wendehals/bricks/cmd/options"
	"github.com/wendehals/bricks/model"
)

const (
	PART_LISTS_OPT   = "partLists"
	PART_LISTS_ARG   = "--" + PART_LISTS_OPT + " PART_LISTS_FILE"
	PART_LISTS_USAGE = "A JSON file containing the user's part lists."

	LOST_OPT   = "lost"
	LOST_ARG   = "--" + LOST_OPT
	LOST_USAGE = "Get a list of the user's lost parts"

	INC_NON_BUILDABLE_OPT   = "includeNonBuildable"
	INC_NON_BUILDABLE_SOPT  = "i"
	INC_NON_BUILDABLE_ARG   = "[-" + INC_NON_BUILDABLE_SOPT + "]"
	INC_NON_BUILDABLE_USAGE = "Include non buildable lists from lists file."

	MERGE_PARTS_OPT   = "mergeParts"
	MERGE_PARTS_SOPT  = "m"
	MERGE_PARTS_ARG   = "[-" + MERGE_PARTS_SOPT + "]"
	MERGE_PARTS_USAGE = "Merge the parts of the given lists to a single JSON file."

	PARTS_FILE_SUFFIX = "_parts.json"
)

var (
	partListsFile       string
	lost                bool
	includeNonBuildable bool
	mergeParts          bool

	partsCmd = &cobra.Command{
		Use: fmt.Sprintf("parts %s %s {%s | %s | %s | %s | %s | %s} %s %s",
			options.CREDENTIALS_ARG, options.JSON_OUTPUT_ARG,
			ALL_ARG, SET_NUM_ARG, SET_LIST_ID_ARG, PART_LIST_ID_ARG, PART_LISTS_ARG, LOST_ARG,
			INC_NON_BUILDABLE_ARG, MERGE_PARTS_ARG),
		Short: "Get a list of parts",
		Long: `The parts command returns a list of parts.
	
This list contains either all parts
- owned by a user,
- of a set,
- of a user set list,
- of a part list,
- of a list of part lists, or
- that are lost by the user.`,

		DisableFlagsInUseLine: true,

		PreRunE: func(cmd *cobra.Command, args []string) error {
			return checkOptionsParts()
		},
		Run: func(cmd *cobra.Command, args []string) {
			executeParts()
		},
	}
)

func init() {
	partsCmd.Flags().BoolVarP(&all, ALL_OPT, ALL_SOPT, false, "Get all parts")

	partsCmd.Flags().StringVarP(&setNum, SET_NUM_OPT, SET_NUM_SOPT, "", SET_NUM_USAGE)
	partsCmd.Flags().UintVarP(&setListId, SET_LIST_ID_OPT, SET_LIST_ID_SOPT, 0, SET_LIST_ID_USAGE)

	partsCmd.Flags().UintVarP(&partListId, PART_LIST_ID_OPT, PART_LIST_ID_SOPT, 0, PART_LIST_ID_USAGE)
	partsCmd.Flags().StringVarP(&partListsFile, PART_LISTS_OPT, "", "", PART_LISTS_USAGE)

	partsCmd.Flags().BoolVarP(&lost, LOST_OPT, "", false, LOST_USAGE)

	partsCmd.Flags().BoolVarP(&includeNonBuildable, INC_NON_BUILDABLE_OPT, INC_NON_BUILDABLE_SOPT, false, INC_NON_BUILDABLE_USAGE)
	partsCmd.Flags().BoolVarP(&mergeParts, MERGE_PARTS_OPT, MERGE_PARTS_SOPT, false, MERGE_PARTS_USAGE)
}

func checkOptionsParts() error {
	optionsProvided := 0
	if all {
		optionsProvided++
	}
	if setNum != "" {
		optionsProvided++
	}
	if setListId != 0 {
		optionsProvided++
	}
	if partListId != 0 {
		optionsProvided++
	}
	if partListsFile != "" {
		optionsProvided++
	}
	if lost {
		optionsProvided++
	}

	if optionsProvided < 1 || optionsProvided > 1 {
		return fmt.Errorf("please provide exactly one option of %s, %s, %s, %s, %s, or %s",
			ALL_ARG, SET_NUM_ARG, SET_LIST_ID_ARG, PART_LIST_ID_ARG, PART_LISTS_ARG, LOST_ARG)
	}

	return nil
}

func executeParts() {
	if all {
		executeAllParts()
	} else if setNum != "" {
		executeSetParts()
	} else if setListId != 0 {
		executeSetListParts()
	} else if partListId != 0 {
		executePartListParts()
	} else if partListsFile != "" {
		executePartListsParts()
	} else if lost {
		executeLostParts()
	}
}

func executeAllParts() {
	allParts := createUsersAPI().GetAllParts()

	if jsonFile == "" {
		jsonFile = options.ReplaceIllegalCharsFromFileName(credentials.UserName) + "_all" + PARTS_FILE_SUFFIX
	}

	model.Save(allParts, jsonFile)
}

func executeSetParts() {
	setParts := api.RetrieveSetParts(createBricksAPI(), setNum)

	if jsonFile == "" {
		jsonFile = options.ReplaceIllegalCharsFromFileName(setNum) + PARTS_FILE_SUFFIX
	}

	model.Save(setParts, jsonFile)
}

func executeSetListParts() {
	setListParts := api.RetrieveSetListParts(createBricksAPI(), createUsersAPI(), setListId)

	if mergeParts {
		mergeAndExport(setListParts)
	} else {
		exportAll(setListParts)
	}
}

func executePartListParts() {
	partListParts := createUsersAPI().GetPartListParts(partListId)

	if jsonFile == "" {
		jsonFile = fmt.Sprint(partListId) + PARTS_FILE_SUFFIX
	}

	model.Save(partListParts, jsonFile)
}

func executePartListsParts() {
	partListsParts := api.RetrievePartListParts(createUsersAPI(), partListsFile, includeNonBuildable)

	if mergeParts {
		mergeAndExport(partListsParts)
	} else {
		exportAll(partListsParts)
	}
}

func executeLostParts() {
	lostParts := createUsersAPI().GetLostParts()

	if jsonFile == "" {
		jsonFile = options.ReplaceIllegalCharsFromFileName(credentials.UserName) + "_lost" + PARTS_FILE_SUFFIX
	}

	model.Save(lostParts, jsonFile)
}

func mergeAndExport(collections []*model.Collection) {
	collection := model.MergeAllCollections(collections)

	if jsonFile == "" {
		var b strings.Builder
		for i := 0; i < len(collection.IDs) && i < 5; i++ {
			b.WriteString(collection.IDs[i])
			b.WriteString("_")
		}
		b.WriteString("parts.json")
		jsonFile = b.String()
	}

	model.Save(collection, jsonFile)
}

func exportAll(collections []*model.Collection) {
	for i, collection := range collections {
		var fileName string
		if jsonFile == "" {
			fileName = collection.IDs[0] + PARTS_FILE_SUFFIX
		} else {
			fileName = fmt.Sprintf("%02d_%s", i+1, jsonFile)
		}

		model.Save(collection, fileName)
	}
}
