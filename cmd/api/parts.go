package api

import (
	"fmt"
	"log"
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
	log.Printf("Retrieving all parts owned by user %s\n", credentials.UserName)
	collection := createUsersAPI().GetAllParts()

	if jsonFile == "" {
		jsonFile = options.ReplaceIllegalCharsFromFileName(credentials.UserName) + "_all" + PARTS_FILE_SUFFIX
	}

	model.ExportToJSON(jsonFile, collection)
}

func executeSetParts() {
	collection := retrieveSetParts(createBricksAPI(), setNum)

	if jsonFile == "" {
		jsonFile = options.ReplaceIllegalCharsFromFileName(setNum) + PARTS_FILE_SUFFIX
	}

	model.ExportToJSON(jsonFile, collection)
}

func executeSetListParts() {
	log.Printf("Retrieving parts of all sets from the set list %d\n", setListId)

	userSets := createUsersAPI().GetSetListSets(setListId)

	bricksAPI := createBricksAPI()
	collections := make([]model.Collection, len(userSets.Sets))
	for i := range userSets.Sets {
		collection := retrieveSetParts(bricksAPI, userSets.Sets[i].Set.SetNum)
		collections[i] = *collection
	}

	if mergeParts {
		mergeAndExport(collections)
	} else {
		exportAll(collections)
	}
}

func executePartListParts() {
	log.Printf("Retrieving parts of part list %d\n", partListId)
	usersAPI := createUsersAPI()

	partListParts := usersAPI.GetPartListParts(partListId)

	if jsonFile == "" {
		jsonFile = fmt.Sprint(partListId) + PARTS_FILE_SUFFIX
	}

	model.ExportToJSON(jsonFile, partListParts)
}

func executePartListsParts() {
	log.Printf("Retrieving parts of all part lists from the part lists file %s\n", partListsFile)
	usersAPI := createUsersAPI()

	partLists := model.ImportPartLists(partListsFile)

	var collections []model.Collection
	for i := range partLists.PartLists {
		if partLists.PartLists[i].IsBuildable || (includeNonBuildable && !partLists.PartLists[i].IsBuildable) {
			partListParts := usersAPI.GetPartListParts(partLists.PartLists[i].ID)

			collections = append(collections, *partListParts)
		}
	}

	if mergeParts {
		mergeAndExport(collections)
	} else {
		exportAll(collections)
	}
}

func executeLostParts() {
	lostParts := createUsersAPI().GetLostParts()

	if jsonFile == "" {
		jsonFile = options.ReplaceIllegalCharsFromFileName(credentials.UserName) + "_lost" + PARTS_FILE_SUFFIX
	}

	model.ExportToJSON(jsonFile, lostParts)
}

func retrieveSetParts(bricksAPI *api.BricksAPI, setNum string) *model.Collection {
	log.Printf("Retrieving parts of set %s\n", setNum)

	set := bricksAPI.GetSet(setNum)

	collection := bricksAPI.GetSetParts(setNum, true)
	collection.IDs = append(collection.IDs, set.SetNum)
	collection.Names = append(collection.Names, set.Name)

	return collection
}

func mergeAndExport(collections []model.Collection) {
	log.Println("Merging parts")

	collection := model.Collection{}
	for i := range collections {
		collection.Add(&collections[i])
		collection.IDs = append(collection.IDs, collection.IDs...)
		collection.Names = append(collection.Names, collection.Names...)
	}

	if jsonFile == "" {
		var b strings.Builder
		for i := 0; i < len(collection.IDs) && i < 5; i++ {
			b.WriteString(collection.IDs[i])
			b.WriteString("_")
		}
		b.WriteString("parts.json")
		jsonFile = b.String()
	}

	model.ExportToJSON(jsonFile, collection)
}

func exportAll(collections []model.Collection) {
	for i := range collections {
		var fileName string
		if jsonFile == "" {
			fileName = collections[i].IDs[0] + PARTS_FILE_SUFFIX
		} else {
			fileName = fmt.Sprintf("%02d_%s", i+1, jsonFile)
		}

		model.ExportToJSON(fileName, collections[i])
	}
}
