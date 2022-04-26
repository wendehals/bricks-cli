package model

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

// PartType represents the type of a Lego part.
type PartType struct {
	Number   string `json:"part_num"`
	Name     string `json:"name"`
	PartURL  string `json:"part_url"`
	ImageURL string `json:"part_img_url"`
}

// ColorType represents the color of a Lego part.
type ColorType struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

// PartEntry represents an entry of a single part in a collection with its quantity and color.
type PartEntry struct {
	Quantity int       `json:"quantity"`
	Part     PartType  `json:"part"`
	Color    ColorType `json:"color"`
	IsSpare  bool      `json:"is_spare"`
}

// SetList represents a user's set list.
type SetList struct {
	ID          uint   `json:"id"`
	IsBuildable bool   `json:"is_buildable"`
	Name        string `json:"name"`
	NumSets     uint   `json:"num_sets"`
}

// SetLists represents all the user's set lists
type SetLists struct {
	User     string    `json:"user"`
	SetLists []SetList `json:"set_lists"`
}

// PartList represents a user's part list.
type PartList struct {
	ID          uint   `json:"id"`
	IsBuildable bool   `json:"is_buildable"`
	Name        string `json:"name"`
	NumParts    uint   `json:"num_parts"`
}

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

// UserSet represents a set owned by the user.
type UserSet struct {
	Quantity       uint    `json:"quantity"`
	IncludesSpares bool    `json:"include_spares"`
	Set            SetType `json:"set"`
}

// UserSets represents all the user's sets
type UserSets struct {
	User string    `json:"user"`
	ID   uint      `json:"set_list_id"`
	Name string    `json:"set_list_name"`
	Sets []UserSet `json:"sets"`
}

func ImportUserSets(userSetsFile string) *UserSets {
	jsonFile, err := os.Open(userSetsFile)
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer jsonFile.Close()

	data, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatalf(err.Error())
	}

	userSets := UserSets{}
	err = json.Unmarshal(data, &userSets)
	if err != nil {
		log.Fatalf(err.Error())
	}

	return &userSets
}

// SetType represents a Lego set.
type SetType struct {
	SetNum      string `json:"set_num"`
	Name        string `json:"name"`
	Year        uint   `json:"year"`
	NumParts    uint   `json:"num_parts"`
	SetURL      string `json:"set_url"`
	SetImageURL string `json:"set_img_url"`
}

type PartColor struct {
	ColorId  uint   `json:"color_id"`
	ImageURL string `json:"part_img_url"`
}
