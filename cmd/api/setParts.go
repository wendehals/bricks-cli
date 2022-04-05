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
	set      string
	setsFile string

	sets *setsList

	setPartsCmd = &cobra.Command{
		Use:   fmt.Sprintf("setParts %s %s {-s SET_NUMBER | --sets SETS_FILE}", options.CREDENTIALS_ARG, options.JSON_OUTPUT_ARG),
		Short: "Get all parts used in the given set(s)",
		Long: `
The command returns a list of parts of the given set.

If a list of sets is given instead it merges the parts of all sets to a single
collection of parts.`,

		DisableFlagsInUseLine: true,

		PreRunE: func(cmd *cobra.Command, args []string) error {
			if (set == "" && setsFile == "") || (set != "" && setsFile != "") {
				return fmt.Errorf("please provide either a single set number with --set or a JSON file containing a list of sets with --sets")
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return executeSetParts()
		},
	}
)

type setsList struct {
	Sets []string `json:"sets"`
}

func init() {
	setPartsCmd.Flags().StringVarP(&set, "set", "s", "", "A set number")
	setPartsCmd.Flags().StringVar(&setsFile, "sets", "", "A JSON file containing a list of set numbers")
}

func executeSetParts() error {
	err := readSets()
	if err != nil {
		return err
	}

	client := http.Client{
		Timeout: time.Second * 5,
	}

	bricksAPI := api.NewBricksAPI(&client, credentials.APIKey)
	if set != "" {
		return processSet(bricksAPI)
	}
	if sets != nil {
		return processSets(bricksAPI)
	}

	return nil
}

func readSets() error {
	if setsFile == "" {
		return nil
	}

	jsonFile, err := os.Open(setsFile)
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	data, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, sets)
	if err != nil {
		return err
	}

	return nil
}

func processSet(bricksAPI *api.BricksAPI) error {
	fmt.Println("Retrieving set number ", set)
	collection, err := bricksAPI.GetSetParts(set, false)
	if err != nil {
		return err
	}

	if jsonFile == "" {
		jsonFile = set + "_setParts.json"
	}

	return model.ExportToJSON(jsonFile, collection)
}

func processSets(bricksAPI *api.BricksAPI) error {
	var collection *model.Collection
	for _, set := range sets.Sets {
		fmt.Println("Retrieving set number ", set)
		setParts, err := bricksAPI.GetSetParts(set, false)
		if err != nil {
			return err
		}

		collection.Add(setParts)
	}

	if jsonFile == "" {
		var b strings.Builder
		for i := 0; i < len(sets.Sets) && i < 5; i++ {
			b.WriteString(sets.Sets[i])
		}
		b.WriteString("_parts.json")
		jsonFile = b.String()
	}

	return model.ExportToJSON(jsonFile, collection)
}
