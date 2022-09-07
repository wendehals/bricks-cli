package model

import (
	"encoding/json"
	"io"
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
		log.Fatal(err)
	}
	defer jsonFile.Close()

	data, err := io.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}

	partLists := PartLists{}
	err = json.Unmarshal(data, &partLists)
	if err != nil {
		log.Fatal(err)
	}

	return &partLists
}
