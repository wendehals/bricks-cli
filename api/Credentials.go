package api

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
	APIKey   string `json:"api_key"`
}

func NewCredentials(credentialsFile string) (*Credentials, error) {
	c := &Credentials{}

	jsonFile, err := os.Open(credentialsFile)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	data, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, c)
	if err != nil {
		return nil, err
	}

	return c, nil
}
