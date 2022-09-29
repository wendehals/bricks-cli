package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/wendehals/bricks/model"
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
	bricks := BricksAPI{}
	bricks.client = client
	bricks.apiKey = apiKey
	bricks.verbose = verbose

	return &bricks
}

// GetSet returns the result of /api/v3/lego/colors/
func (b *BricksAPI) GetColors() []model.Color {
	log.Print("Retrieving a list of all colors...")

	url := fmt.Sprintf(BRICKS_URL, "colors/?inc_color_details=0")
	colors := []model.Color{}
	colorsPage := colorsPageResult{}

	err := b.requestPage(url, &colorsPage)
	if err != nil {
		log.Fatalf(COLORS_ERR_MSG, err.Error())
	}

	colors = append(colors, colorsPage.Results...)
	for len(colorsPage.Next) > 0 {
		url = colorsPage.Next
		colorsPage = colorsPageResult{}
		err = b.requestPage(url, &colorsPage)
		if err != nil {
			log.Fatalf(COLORS_ERR_MSG, err.Error())
		}

		colors = append(colors, colorsPage.Results...)
	}

	return colors
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

	collection := model.Collection{}

	subPath := fmt.Sprintf("sets/%s/parts/", setNum)
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

	return &collection
}

// GetPartColors returns the result of /api/v3/lego/parts/{part_num}/colors/
func (b *BricksAPI) GetPartColors(partNum string) []model.PartColor {
	partColors := []model.PartColor{}
	partColorsPage := partColorsPageResult{}

	subPath := fmt.Sprintf("parts/%s/colors/", partNum)
	url := fmt.Sprintf(BRICKS_URL, subPath)

	err := b.requestPage(url, &partColorsPage)
	if err != nil {
		log.Fatalf(PART_COLORS_ERR_MSG, err.Error())
	}

	partColors = append(partColors, partColorsPage.Results...)
	for len(partColorsPage.Next) > 0 {
		url = partColorsPage.Next
		partColorsPage = partColorsPageResult{}
		err = b.requestPage(url, &partColorsPage)
		if err != nil {
			log.Fatalf(PART_COLORS_ERR_MSG, err.Error())
		}

		partColors = append(partColors, partColorsPage.Results...)
	}

	return partColors
}
