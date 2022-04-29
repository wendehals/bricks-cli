package model

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

// PartLists represents all the user's part lists
type PartLists struct {
	User      string     `json:"user"`
	PartLists []PartList `json:"part_lists"`
}

func ImportPartLists(partListsFile string) *PartLists {
	jsonFile, err := os.Open(partListsFile)
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer jsonFile.Close()

	data, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatalf(err.Error())
	}

	partLists := PartLists{}
	err = json.Unmarshal(data, &partLists)
	if err != nil {
		log.Fatalf(err.Error())
	}

	return &partLists
}
