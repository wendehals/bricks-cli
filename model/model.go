package model

import (
	"encoding/json"
	"io/ioutil"
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
	SetLists []SetList `json:"setLists"`
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
	PartLists []PartList `json:"partLists"`
}

func ImportPartLists(partListsFile string) (*PartLists, error) {
	jsonFile, err := os.Open(partListsFile)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	data, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	partLists := PartLists{}
	err = json.Unmarshal(data, &partLists)
	if err != nil {
		return nil, err
	}

	return &partLists, nil
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
	ID   uint      `json:"setListId"`
	Name string    `json:"setListName"`
	Sets []UserSet `json:"sets"`
}

func ImportUserSets(userSetsFile string) (*UserSets, error) {
	jsonFile, err := os.Open(userSetsFile)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	data, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	userSets := UserSets{}
	err = json.Unmarshal(data, &userSets)
	if err != nil {
		return nil, err
	}

	return &userSets, nil
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
