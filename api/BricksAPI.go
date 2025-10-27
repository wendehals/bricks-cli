package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/wendehals/bricks-cli/model"
	"github.com/wendehals/bricks-cli/services"
)

const (
	BRICKS_URL string = REBRICKABLE_BASE_URL + "lego/%s"

	COLORS_ERR_MSG      string = "list of colors could not be retrieved: %s"
	SET_PARTS_ERR_MSG   string = "parts list of set number %s could not be retrieved: %s"
	PART_COLORS_ERR_MSG string = "part colors could not be retrieved: %s"
)

// BricksAPI provides API for accessing Lego's data at Rebrickable
type BricksAPI struct {
	AbstractAPI
}

// NewBricksAPI creates a new object of BricksAPI
func NewBricksAPI(client *http.Client, apiKey string, verbose bool) *BricksAPI {
	return &BricksAPI{
		AbstractAPI: *NewAbstractAPI(client, apiKey, verbose),
	}
}

// GetSet returns the result of /api/v3/lego/sets/{set_num}/
func (b *BricksAPI) GetSet(setNum string) *model.Set {
	log.Printf("Retrieving details about set %s", setNum)

	set := model.Set{}

	subPath := fmt.Sprintf("sets/%s/", setNum)
	url := fmt.Sprintf(BRICKS_URL, subPath)

	err := b.requestPage(url, &set)
	if err != nil {
		log.Fatalf("details of set %s could not be retrieved: %s", setNum, err.Error())
	}

	return &set
}

// GetSetParts returns the result of /api/v3/lego/sets/{set_num}/parts/
func (b *BricksAPI) GetSetParts(setNum string, includeMiniFigs bool) *model.Collection {
	log.Printf("Retrieving parts of set %s", setNum)

	collection := model.NewCollection()

	subPath := fmt.Sprintf("sets/%s/parts/?inc_color_details=0", setNum)
	if includeMiniFigs {
		subPath += "?inc_minifig_parts=1"
	}
	url := fmt.Sprintf(BRICKS_URL, subPath)

	setParts := partsPageResult{}
	err := b.requestPage(url, &setParts)
	if err != nil {
		log.Fatalf(SET_PARTS_ERR_MSG, setNum, err.Error())
	}

	collection.Parts = append(collection.Parts, setParts.Results...)
	for len(setParts.Next) > 0 {
		url = setParts.Next
		setParts = partsPageResult{}
		err = b.requestPage(url, &setParts)
		if err != nil {
			log.Fatalf(SET_PARTS_ERR_MSG, setNum, err.Error())
		}

		collection.Parts = append(collection.Parts, setParts.Results...)
	}

	services.GetShapes(false).ComplementShapesData(collection.Parts)

	return collection
}
