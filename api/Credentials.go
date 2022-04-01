package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

const errorMessage = "please provide a valid JSON file containing the Rebrickable credentials:\n   %s"

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
	APIKey   string `json:"api_key"`
}

func ImportCredentials(fileName string) (*Credentials, error) {
	credentials := &Credentials{}

	jsonFile, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf(errorMessage, err)
	}
	defer jsonFile.Close()

	data, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, fmt.Errorf(errorMessage, err)
	}

	err = json.Unmarshal(data, credentials)
	if err != nil {
		return nil, fmt.Errorf(errorMessage, err)
	}

	return credentials, nil
}
