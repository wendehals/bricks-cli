package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/wendehals/bricks/model"
)

const bricksURL string = RebrickableBaseURL + "lego/%s"

// BricksAPI provides API for accessing Lego's data
type BricksAPI struct {
	client *http.Client
	apiKey string
}

// NewBricksAPI creates a new object of BricksAPI
func NewBricksAPI(client *http.Client, apiKey string) *BricksAPI {
	bricks := BricksAPI{}
	bricks.client = client
	bricks.apiKey = apiKey

	return &bricks
}

// setPartsPageResult contains the result of /api/v3/lego/sets/{set_num}/parts/
type setPartsPageResult struct {
	Next    string            `json:"next"`
	Results []model.PartEntry `json:"results"`
}

// GetSetParts returns the result of /api/v3/lego/sets/{set_num}/parts/
func (b *BricksAPI) GetSetParts(setNum string, includeMiniFigs bool) *model.Collection {
	collection := model.Collection{}

	subPath := fmt.Sprintf("sets/%s/parts", setNum)
	if includeMiniFigs {
		subPath += "/?inc_minifig_parts=1"
	}
	url := fmt.Sprintf(bricksURL, subPath)

	setParts := setPartsPageResult{}
	err := b.requestPage(url, &setParts)
	if err != nil {
		log.Printf("Set number %s parts list could not be retrieved: %s\n", setNum, err.Error())
		return &model.Collection{}
	}

	collection.Parts = append(collection.Parts, setParts.Results...)
	for len(setParts.Next) > 0 {
		url = setParts.Next
		setParts = setPartsPageResult{}
		err = b.requestPage(url, &setParts)
		if err != nil {
			log.Printf("Set number %s parts list could not be retrieved: %s\n", setNum, err.Error())
			return &model.Collection{}
		}

		collection.Parts = append(collection.Parts, setParts.Results...)
	}

	return &collection
}

func (b *BricksAPI) requestPage(url string, v interface{}) error {
	reqest, err := CreateGetRequest(url, b.apiKey)
	if err != nil {
		return err
	}

	err = DoRequest(b.client, reqest, v)
	if err != nil {
		return err
	}

	return nil
}
