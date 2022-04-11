package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/wendehals/bricks/api"
	"github.com/wendehals/bricks/cmd/options"
	"github.com/wendehals/bricks/model"
)

var (
	setsFile string

	setPartsCmd = &cobra.Command{
		Use:   fmt.Sprintf("setParts %s %s {-n SET_NUMBER|-f SETS_FILE}", options.CREDENTIALS_ARG, options.JSON_OUTPUT_ARG),
		Short: "Get all parts used in the given set(s)",
		Long: `
The command returns a list of parts of the given set.

If a list of sets is given instead it merges the parts of all sets to a single
collection of parts.`,

		DisableFlagsInUseLine: true,

		PreRunE: func(cmd *cobra.Command, args []string) error {
			return checkOptionsSetParts()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return executeSetParts()
		},
	}
)

func init() {
	setPartsCmd.Flags().StringVarP(&setNum, "set", "n", "", "A set number")
	setPartsCmd.Flags().StringVarP(&setsFile, "setsFile", "f", "", "A JSON file containing a list of sets")
}

func checkOptionsSetParts() error {
	if (setNum == "" && setsFile == "") || (setNum != "" && setsFile != "") {
		return fmt.Errorf("please provide either a single set number with --set or a JSON file containing a list of sets with --sets")
	}
	return nil
}

func executeSetParts() error {
	bricksAPI := api.NewBricksAPI(createClient(), credentials.APIKey)

	if setsFile != "" {
		sets, err := readUsersSets()
		if err != nil {
			return err
		}

		return processUsersSets(bricksAPI, sets)
	}

	if setNum != "" {
		return processSet(bricksAPI)
	}

	return nil
}

func readUsersSets() (*model.UsersSets, error) {
	jsonFile, err := os.Open(setsFile)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	data, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	usersSets := model.UsersSets{}
	err = json.Unmarshal(data, &usersSets)
	if err != nil {
		return nil, err
	}

	return &usersSets, nil
}

func processSet(bricksAPI *api.BricksAPI) error {
	log.Printf("Retrieving details about set %s\n", setNum)
	set, err := bricksAPI.GetSet(setNum)
	if err != nil {
		return err
	}

	log.Printf("Retrieving parts of set %s\n", setNum)
	collection, err := bricksAPI.GetSetParts(setNum, false)
	if err != nil {
		return err
	}
	collection.IDs = append(collection.IDs, set.SetNum)
	collection.Names = append(collection.Names, set.Name)

	if jsonFile == "" {
		jsonFile = setNum + "_setParts.json"
	}

	return model.ExportToJSON(jsonFile, collection)
}

func processUsersSets(bricksAPI *api.BricksAPI, usersSets *model.UsersSets) error {
	collection := model.Collection{}
	for i := range usersSets.Sets {
		log.Printf("Retrieving set number %s\n", setNum)
		setParts, err := bricksAPI.GetSetParts(usersSets.Sets[i].Set.SetNum, false)
		if err != nil {
			return err
		}
		setParts.IDs = append(setParts.IDs, usersSets.Sets[i].Set.SetNum)
		setParts.Names = append(setParts.Names, usersSets.Sets[i].Set.Name)

		collection.Add(setParts)
	}

	if jsonFile == "" {
		var b strings.Builder
		for i := 0; i < len(usersSets.Sets) && i < 5; i++ {
			b.WriteString(usersSets.Sets[i].Set.SetNum)
		}
		b.WriteString("_parts.json")
		jsonFile = b.String()
	}

	return model.ExportToJSON(jsonFile, collection)
}
