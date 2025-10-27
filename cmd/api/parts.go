package api

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/wendehals/bricks-cli/api"
	"github.com/wendehals/bricks-cli/cmd/options"
	"github.com/wendehals/bricks-cli/model"
	"github.com/wendehals/bricks-cli/services"
)

const (
	USER_SET_NUM_OPT   = "userSet"
	USER_SET_NUM_SOPT  = "u"
	USER_SET_NUM_ARG   = "-" + USER_SET_NUM_SOPT + " SET_NUM"
	USER_SET_NUM_USAGE = "all parts of a user's set without lost parts"

	PART_LISTS_OPT   = "partLists"
	PART_LISTS_ARG   = "--" + PART_LISTS_OPT + " PART_LISTS_FILE"
	PART_LISTS_USAGE = "a JSON file containing the user's part lists"

	LOST_OPT   = "lost"
	LOST_ARG   = "--" + LOST_OPT
	LOST_USAGE = "get a list of the user's lost parts"

	INC_NON_BUILDABLE_OPT   = "includeNonBuildable"
	INC_NON_BUILDABLE_SOPT  = "b"
	INC_NON_BUILDABLE_ARG   = "[-" + INC_NON_BUILDABLE_SOPT + "]"
	INC_NON_BUILDABLE_USAGE = "include non buildable lists from lists file"

	INC_MINI_FIGS_OPT   = "includeMiniFigs"
	INC_MINI_FIGS_SOPT  = "f"
	INC_MINI_FIGS_ARG   = "[-" + INC_MINI_FIGS_SOPT + "]"
	INC_MINI_FIGS_USAGE = "include mini figures from sets"

	MERGE_PARTS_OPT   = "mergeParts"
	MERGE_PARTS_SOPT  = "m"
	MERGE_PARTS_ARG   = "[-" + MERGE_PARTS_SOPT + "]"
	MERGE_PARTS_USAGE = "merge the parts of the given lists to a single parts file"

	PARTS_FILE_SUFFIX = ".parts"
)

var (
	userSetNum          string
	partListsFile       string
	lost                bool
	includeNonBuildable bool
	includeMiniFigs     bool
	mergeParts          bool

	partsCmd = &cobra.Command{
		Use: fmt.Sprintf("parts %s %s {%s | %s | %s | %s | %s | %s | %s} %s %s %s",
			options.CREDENTIALS_ARG, options.OUTPUT_FILE_ARG,
			ALL_ARG, options.SET_NUM_ARG, USER_SET_NUM_ARG, SET_LIST_ID_ARG, PART_LIST_ID_ARG, PART_LISTS_ARG, LOST_ARG,
			INC_NON_BUILDABLE_ARG, INC_MINI_FIGS_ARG, MERGE_PARTS_ARG),
		Short: "Get a list of parts",
		Long: `The parts command returns a list of parts.
	
This list contains either all parts
- owned by a user,
- of a set,
- of a user's set without lost parts,
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

	partsCmd.Flags().StringVarP(&setNum, options.SET_NUM_OPT, options.SET_NUM_SOPT, "", options.SET_NUM_USAGE)
	partsCmd.Flags().StringVarP(&userSetNum, USER_SET_NUM_OPT, "", "", USER_SET_NUM_USAGE)
	partsCmd.Flags().UintVarP(&setListId, SET_LIST_ID_OPT, SET_LIST_ID_SOPT, 0, SET_LIST_ID_USAGE)

	partsCmd.Flags().UintVarP(&partListId, PART_LIST_ID_OPT, PART_LIST_ID_SOPT, 0, PART_LIST_ID_USAGE)
	partsCmd.Flags().StringVarP(&partListsFile, PART_LISTS_OPT, "", "", PART_LISTS_USAGE)

	partsCmd.Flags().BoolVarP(&lost, LOST_OPT, "", false, LOST_USAGE)

	partsCmd.Flags().BoolVarP(&includeNonBuildable, INC_NON_BUILDABLE_OPT, INC_NON_BUILDABLE_SOPT, false, INC_NON_BUILDABLE_USAGE)
	partsCmd.Flags().BoolVarP(&includeMiniFigs, INC_MINI_FIGS_OPT, INC_MINI_FIGS_SOPT, false, INC_MINI_FIGS_USAGE)
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
	if userSetNum != "" {
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
		return fmt.Errorf("please provide exactly one option of %s, %s, %s, %s, %s, %s, or %s",
			ALL_ARG, options.SET_NUM_ARG, USER_SET_NUM_ARG, SET_LIST_ID_ARG, PART_LIST_ID_ARG,
			PART_LISTS_ARG, LOST_ARG)
	}

	return nil
}

func executeParts() {
	if all {
		executeAllParts()
	} else if setNum != "" {
		executeSetParts()
	} else if userSetNum != "" {
		executeUserSetParts()
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

	if outputFile == "" {
		outputFile = options.ReplaceIllegalCharsFromFileName(credentials.UserName) + "_all" + PARTS_FILE_SUFFIX
	}

	allParts.Save(outputFile)
}

func executeSetParts() {
	setParts := api.RetrieveSetParts(createBricksAPI(), setNum, includeMiniFigs)

	if outputFile == "" {
		outputFile = options.ReplaceIllegalCharsFromFileName(setNum) + PARTS_FILE_SUFFIX
	}

	setParts.Save(outputFile)
}

func executeUserSetParts() {
	setPartsWithoutLost := api.RetrieveUserSetParts(createBricksAPI(), createUsersAPI(), userSetNum, includeMiniFigs)

	if outputFile == "" {
		outputFile = options.ReplaceIllegalCharsFromFileName(userSetNum) + "_wo_lost" + PARTS_FILE_SUFFIX
	}

	setPartsWithoutLost.Save(outputFile)
}

func executeSetListParts() {
	setListParts := api.RetrieveSetListParts(createBricksAPI(), createUsersAPI(), setListId, includeMiniFigs)

	if mergeParts {
		mergeAndSave(setListParts)
	} else {
		saveAll(setListParts)
	}
}

func executePartListParts() {
	partListParts := api.RetrievePartListParts(createUsersAPI(), partListId)

	if outputFile == "" {
		outputFile = fmt.Sprint(partListId) + PARTS_FILE_SUFFIX
	}

	partListParts.Save(outputFile)
}

func executePartListsParts() {
	partListsParts := api.RetrievePartListsParts(createUsersAPI(), partListsFile, includeNonBuildable)

	if mergeParts {
		mergeAndSave(partListsParts)
	} else {
		saveAll(partListsParts)
	}
}

func executeLostParts() {
	lostParts := createUsersAPI().GetLostParts()

	if outputFile == "" {
		outputFile = options.ReplaceIllegalCharsFromFileName(credentials.UserName) + "_lost" + PARTS_FILE_SUFFIX
	}

	lostParts.Save(outputFile)
}

func mergeAndSave(collections []model.Collection) {
	collection := services.MergeAllCollections(collections)
	collection.SortByColorAndName(false)

	if outputFile == "" {
		var b strings.Builder
		for i := 0; i < len(collection.Sets) && i < 5; i++ {
			b.WriteString(collection.Sets[i].Number)
			if i < len(collection.Sets)-1 || i < 4 {
				b.WriteString("_")
			}
		}
		b.WriteString(".parts")
		outputFile = b.String()
	}

	collection.Save(outputFile)
}

func saveAll(collections []model.Collection) {
	for i, collection := range collections {
		var fileName string
		if outputFile == "" {
			fileName = collection.Sets[0].Number + PARTS_FILE_SUFFIX
		} else {
			fileName = fmt.Sprintf("%02d_%s", i+1, outputFile)
		}

		collection.Save(fileName)
	}
}
