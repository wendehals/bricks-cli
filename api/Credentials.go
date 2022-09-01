package api

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

const CREDENTIALS_ERR_MSG = "please provide a valid JSON file containing the Rebrickable credentials:\n   %s"

type Credentials struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	APIKey   string `json:"api_key"`
}

func ImportCredentials(fileName string) (*Credentials, error) {
	credentials := &Credentials{}

	jsonFile, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf(CREDENTIALS_ERR_MSG, err)
	}
	defer jsonFile.Close()

	data, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, fmt.Errorf(CREDENTIALS_ERR_MSG, err)
	}

	err = json.Unmarshal(data, credentials)
	if err != nil {
		return nil, fmt.Errorf(CREDENTIALS_ERR_MSG, err)
	}

	return credentials, nil
}
