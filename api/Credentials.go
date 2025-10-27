package api

import (
	"fmt"

	"github.com/wendehals/bricks-cli/model"
)

const CREDENTIALS_ERR_MSG = "please provide a valid JSON file containing the Rebrickable credentials:\n   %s"

type Credentials struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	APIKey   string `json:"api_key"`
}

func ImportCredentials(fileName string) (*Credentials, error) {
	credentials, err := model.LoadE[Credentials](fileName)
	if err != nil {
		return nil, fmt.Errorf(CREDENTIALS_ERR_MSG, err)
	}

	return &credentials, nil
}
