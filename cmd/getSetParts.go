/*
Copyright © 2022 Lothar Wendehals
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/wendehals/bricks/api"
	"github.com/wendehals/bricks/model"
)

var (
	set      string
	setsFile string

	sets *setsList

	getSetPartsCmd = &cobra.Command{
		Use:   "getSetParts [-c FILE] [-o FILE] [-j FILE] {-s SET_NUMBER | --sets FILE}",
		Short: "Returns a list of parts used in the given set or sets",
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
			return excuteGetSetParts()
		},
	}
)

type setsList struct {
	Sets []string `json:"sets"`
}

func init() {
	rootCmd.AddCommand(getSetPartsCmd)

	getSetPartsCmd.Flags().StringVarP(&set, "set", "s", "", "A set number")
	getSetPartsCmd.Flags().StringVar(&setsFile, "sets", "", "A JSON file containing a list of sets")
}

func excuteGetSetParts() error {
	err := readSets()
	if err != nil {
		return err
	}

	client := http.Client{
		Timeout: time.Second * 5,
	}

	bricksAPI := api.NewBricksAPI(&client, credentials.APIKey)
	if set != "" {
		processSet(bricksAPI)
	} else if sets != nil {
		processSets(bricksAPI)
	}

	return nil
}

func readSets() error {
	if setsFile == "" {
		return nil
	}

	c := &setsList{}

	jsonFile, err := os.Open(setsFile)
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	data, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, c)
	if err != nil {
		return err
	}

	sets = c

	return nil
}

func processSet(bricksAPI *api.BricksAPI) {
	fmt.Println("Retrieving set number ", set)
	collection := bricksAPI.GetSetParts(set, false)
	collection.ExportToHTML(htmlFile)

	if jsonFile != "" {
		collection.Export(jsonFile)
	}
}

func processSets(bricksAPI *api.BricksAPI) {
	var collection *model.Collection
	for _, set := range sets.Sets {
		fmt.Println("Retrieving set number ", set)
		collection.Add(bricksAPI.GetSetParts(set, false))
	}

	collection.ExportToHTML(htmlFile)

	if jsonFile != "" {
		collection.Export(jsonFile)
	}
}
