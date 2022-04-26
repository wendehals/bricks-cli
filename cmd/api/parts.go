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
	includeNonBuildable bool
	mergeParts          bool

	partsCmd = &cobra.Command{
		Use: fmt.Sprintf("parts %s %s {%s | %s | %s | %s | %s} %s %s",
			options.CREDENTIALS_ARG, options.JSON_OUTPUT_ARG,
			ALL_ARG, SET_NUM_ARG, SET_LIST_ID_ARG, PART_LIST_ID_ARG, PART_LISTS_ARG,
			INC_NON_BUILDABLE_ARG, MERGE_PARTS_ARG),
		Short: "Get a list of parts",
		Long: `The parts command returns a list of parts.
	
This list contains either all parts
- owned by a user,
- of a set,
- of a user set list,
- of a part list, or
- of a list of part lists.`,

		DisableFlagsInUseLine: true,

		PreRunE: func(cmd *cobra.Command, args []string) error {
			return checkOptionsParts()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return executeParts()
		},
	}
)

func init() {
	partsCmd.Flags().BoolVarP(&all, ALL_OPT, ALL_SOPT, false, "Get all parts")

	partsCmd.Flags().StringVarP(&setNum, SET_NUM_OPT, SET_NUM_SOPT, "", SET_NUM_USAGE)
	partsCmd.Flags().UintVarP(&setListId, SET_LIST_ID_OPT, SET_LIST_ID_SOPT, 0, SET_LIST_ID_USAGE)

	partsCmd.Flags().UintVarP(&partListId, PART_LIST_ID_OPT, PART_LIST_ID_SOPT, 0, PART_LIST_ID_USAGE)
	partsCmd.Flags().StringVarP(&partListsFile, PART_LISTS_OPT, "", "", PART_LISTS_USAGE)

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

	if optionsProvided < 1 || optionsProvided > 1 {
		return fmt.Errorf("please provide exactly one option of -%s, -%s, -%s, -%s, or --%s",
			ALL_SOPT, SET_NUM_SOPT, SET_LIST_ID_SOPT, PART_LIST_ID_SOPT, PART_LISTS_OPT)
	}

	return nil
}

func executeParts() error {
	if all {
		return executeAllParts()
	}

	if setNum != "" {
		return executeSetParts()
	}

	if setListId != 0 {
		return executeSetListParts()
	}

	if partListId != 0 {
		return executePartListParts()
	}

	if partListsFile != "" {
		return executePartListsParts()
	}

	return nil
}

func executeAllParts() error {
	log.Printf("Retrieving all parts owned by user %s\n", credentials.UserName)
	collection, err := createUsersAPI().GetAllParts()
	if err != nil {
		return err
	}

	if jsonFile == "" {
		jsonFile = options.ReplaceIllegalCharsFromFileName(credentials.UserName) + PARTS_FILE_SUFFIX
	}

	return model.ExportToJSON(jsonFile, collection)
}

func executeSetParts() error {
	collection, err := retrieveSetParts(createBricksAPI(), setNum)
	if err != nil {
		return err
	}

	if jsonFile == "" {
		jsonFile = options.ReplaceIllegalCharsFromFileName(setNum) + PARTS_FILE_SUFFIX
	}

	return model.ExportToJSON(jsonFile, collection)
}

func executeSetListParts() error {
	log.Printf("Retrieving parts of all sets from the set list %d\n", setListId)

	userSets, err := createUsersAPI().GetSetListSets(setListId)
	if err != nil {
		return err
	}

	bricksAPI := createBricksAPI()
	collections := make([]model.Collection, len(userSets.Sets))
	for i := range userSets.Sets {
		collection, err := retrieveSetParts(bricksAPI, userSets.Sets[i].Set.SetNum)
		if err != nil {
			return err
		}
		collections[i] = *collection
	}

	if mergeParts {
		mergeAndExport(collections)
	} else {
		exportAll(collections)
	}

	return nil
}

func executePartListParts() error {
	log.Printf("Retrieving parts of part list %d\n", partListId)
	usersAPI := createUsersAPI()

	partListParts, err := usersAPI.GetPartListParts(partListId)
	if err != nil {
		return err
	}

	if jsonFile == "" {
		jsonFile = fmt.Sprint(partListId) + PARTS_FILE_SUFFIX
	}

	return model.ExportToJSON(jsonFile, partListParts)
}

func executePartListsParts() error {
	log.Printf("Retrieving parts of all part lists from the part lists file %s\n", partListsFile)
	usersAPI := createUsersAPI()

	partLists, err := model.ImportPartLists(partListsFile)
	if err != nil {
		return err
	}

	var collections []model.Collection
	for i := range partLists.PartLists {
		if partLists.PartLists[i].IsBuildable || (includeNonBuildable && !partLists.PartLists[i].IsBuildable) {
			partListParts, err := usersAPI.GetPartListParts(partLists.PartLists[i].ID)
			if err != nil {
				return err
			}

			collections = append(collections, *partListParts)
		}
	}

	if mergeParts {
		mergeAndExport(collections)
	} else {
		exportAll(collections)
	}

	return nil
}

func retrieveSetParts(bricksAPI *api.BricksAPI, setNum string) (*model.Collection, error) {
	log.Printf("Retrieving parts of set %s\n", setNum)

	set, err := bricksAPI.GetSet(setNum)
	if err != nil {
		return nil, err
	}

	collection, err := bricksAPI.GetSetParts(setNum, true)
	if err != nil {
		return nil, err
	}
	collection.IDs = append(collection.IDs, set.SetNum)
	collection.Names = append(collection.Names, set.Name)

	return collection, nil
}

func mergeAndExport(collections []model.Collection) error {
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

	return model.ExportToJSON(jsonFile, collection)
}

func exportAll(collections []model.Collection) error {
	for i := range collections {
		var fileName string
		if jsonFile == "" {
			fileName = collections[i].IDs[0] + PARTS_FILE_SUFFIX
		} else {
			fileName = fmt.Sprintf("%02d_%s", i+1, jsonFile)
		}

		err := model.ExportToJSON(fileName, collections[i])
		if err != nil {
			return err
		}
	}

	return nil
}
