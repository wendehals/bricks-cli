package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/wendehals/bricks/model"
)

const bricksURL string = rebrickableBaseURL + "lego/%s"

// BricksAPI provides API for accessing Lego's data
type BricksAPI struct {
	AbstractAPI
}

// NewBricksAPI creates a new object of BricksAPI
func NewBricksAPI(client *http.Client, apiKey string) *BricksAPI {
	bricks := BricksAPI{}
	bricks.client = client
	bricks.apiKey = apiKey

	return &bricks
}

// GetSetParts returns the result of /api/v3/lego/sets/{set_num}/parts/
func (b *BricksAPI) GetSetParts(setNum string, includeMiniFigs bool) *model.Collection {
	collection := model.Collection{}

	subPath := fmt.Sprintf("sets/%s/parts", setNum)
	if includeMiniFigs {
		subPath += "/?inc_minifig_parts=1"
	}
	url := fmt.Sprintf(bricksURL, subPath)

	setParts := partsPageResult{}
	err := b.requestPage(url, &setParts)
	if err != nil {
		log.Printf("Set number %s parts list could not be retrieved: %s\n", setNum, err.Error())
		return &model.Collection{}
	}

	collection.Parts = append(collection.Parts, setParts.Results...)
	for len(setParts.Next) > 0 {
		url = setParts.Next
		setParts = partsPageResult{}
		err = b.requestPage(url, &setParts)
		if err != nil {
			log.Printf("Set number %s parts list could not be retrieved: %s\n", setNum, err.Error())
			return &model.Collection{}
		}

		collection.Parts = append(collection.Parts, setParts.Results...)
	}

	return &collection
}
