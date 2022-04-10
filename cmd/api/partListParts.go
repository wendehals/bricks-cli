package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/wendehals/bricks/api"
	"github.com/wendehals/bricks/cmd/options"
	"github.com/wendehals/bricks/model"
)

var (
	partListsFile       string
	includeNonBuildable bool
	mergeParts          bool

	partListPartsCmd = &cobra.Command{
		Use:   fmt.Sprintf("partListParts %s %s {-l LIST_ID|-p PART_LISTS_FILE} [-i] [-m]", options.CREDENTIALS_ARG, options.JSON_OUTPUT_ARG),
		Short: "Get all parts of (a) user defined part list(s).",
		Long:  "The partListParts command returns all parts of a single or multiple user defined part lists.",

		DisableFlagsInUseLine: true,

		PreRunE: func(cmd *cobra.Command, args []string) error {
			return checkOptionsPartListParts()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return executePartListParts()
		},
	}
)

func init() {
	partListPartsCmd.Flags().UintVarP(&listId, "listId", "l", 0,
		"The list id of a user defined part list")
	partListPartsCmd.Flags().StringVarP(&partListsFile, "partLists", "p", "",
		"A JSON file containing the user's part lists.")
	partListPartsCmd.Flags().BoolVarP(&includeNonBuildable, "includeNonBuildable", "i", false,
		"Include non buildable part lists from part lists file.")
	partListPartsCmd.Flags().BoolVarP(&mergeParts, "mergeParts", "m", false,
		"Merge the parts of the given part lists to a single JSON file.")
}

func checkOptionsPartListParts() error {
	if listId == 0 && partListsFile == "" {
		return fmt.Errorf("please provide either a list id of a user defined part list or a part lists JSON file")
	}
	return nil
}

func executePartListParts() error {
	client := http.Client{
		Timeout: time.Second * 5,
	}

	usersAPI := api.NewUsersAPI(&client, credentials)

	if listId != 0 {
		return processListId(usersAPI, listId)
	}
	if partListsFile != "" {
		return processPartLists(usersAPI)
	}
	return nil
}

func processListId(u *api.UsersAPI, id uint) error {
	partListParts, err := u.GetPartListParts(id)
	if err != nil {
		return err
	}

	fileName := jsonFile
	if fileName == "" {
		fileName = fmt.Sprint(id) + "_parts.json"
	}

	return model.ExportToJSON(fileName, partListParts)
}

func processPartLists(u *api.UsersAPI) error {
	partLists, err := readPartLists()
	if err != nil {
		return err
	}

	collection := model.Collection{}
	var fileName strings.Builder
	for i := range partLists.PartLists {
		if partLists.PartLists[i].IsBuildable || (includeNonBuildable && !partLists.PartLists[i].IsBuildable) {
			if mergeParts {
				partListParts, err := u.GetPartListParts(partLists.PartLists[i].ID)
				if err != nil {
					return err
				}
				collection.Add(partListParts)
				if i < 10 {
					fileName.WriteString(fmt.Sprint(partLists.PartLists[i].ID))
					fileName.WriteString(`_`)
				}
			} else {
				err = processListId(u, partLists.PartLists[i].ID)
				if err != nil {
					return err
				}
			}
		}
	}

	if mergeParts {
		if jsonFile == "" {
			jsonFile = fmt.Sprint(fileName.String()) + "parts.json"
		}

		return model.ExportToJSON(jsonFile, collection)
	}

	return nil
}

func readPartLists() (*model.PartLists, error) {
	jsonFile, err := os.Open(partListsFile)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	data, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	partLists := model.PartLists{}
	err = json.Unmarshal(data, &partLists)
	if err != nil {
		return nil, err
	}

	return &partLists, nil
}
