package model

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

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
