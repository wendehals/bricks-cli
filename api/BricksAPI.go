package api

import (
	"fmt"
	"net/http"

	"github.com/wendehals/bricks/model"
)

const (
	BRICKS_URL string = REBRICKABLE_BASE_URL + "lego/%s"

	SET_PARTS_ERROR_MSG   string = "set number %s parts list could not be retrieved: %s"
	PART_COLORS_ERROR_MSG string = "part colors could not be retrieved: %s"
)

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
func (b *BricksAPI) GetSetParts(setNum string, includeMiniFigs bool) (*model.Collection, error) {
	collection := model.Collection{}

	subPath := fmt.Sprintf("sets/%s/parts", setNum)
	if includeMiniFigs {
		subPath += "/?inc_minifig_parts=1"
	}
	url := fmt.Sprintf(BRICKS_URL, subPath)

	setParts := partsPageResult{}
	err := b.requestPage(url, &setParts)
	if err != nil {
		return nil, fmt.Errorf(SET_PARTS_ERROR_MSG, setNum, err.Error())
	}

	collection.Parts = append(collection.Parts, setParts.Results...)
	for len(setParts.Next) > 0 {
		url = setParts.Next
		setParts = partsPageResult{}
		err = b.requestPage(url, &setParts)
		if err != nil {
			return nil, fmt.Errorf(SET_PARTS_ERROR_MSG, setNum, err.Error())
		}

		collection.Parts = append(collection.Parts, setParts.Results...)
	}

	return &collection, nil
}

// GetPartColors returns the result of /api/v3/lego/parts/{part_num}/colors/
func (b *BricksAPI) GetPartColors(partNum string) ([]model.PartColor, error) {
	partColors := []model.PartColor{}
	partColorsPage := partColorsPageResult{}

	subPath := fmt.Sprintf("parts/%s/colors/", partNum)
	url := fmt.Sprintf(BRICKS_URL, subPath)

	err := b.requestPage(url, &partColorsPage)
	if err != nil {
		return nil, fmt.Errorf(PART_COLORS_ERROR_MSG, err.Error())
	}

	partColors = append(partColors, partColorsPage.Results...)
	for len(partColorsPage.Next) > 0 {
		url = partColorsPage.Next
		partColorsPage = partColorsPageResult{}
		err = b.requestPage(url, &partColorsPage)
		if err != nil {
			return nil, fmt.Errorf(PART_COLORS_ERROR_MSG, err.Error())
		}

		partColors = append(partColors, partColorsPage.Results...)
	}

	return partColors, nil
}

func (b *BricksAPI) ReplaceImagesByMatchingColor(collection *model.Collection) error {
	for i := range collection.Parts {
		partEntry := collection.Parts[i]
		partColors, err := b.GetPartColors(partEntry.Part.Number)
		if err != nil {
			return err
		}

		for j := range partColors {
			if partColors[j].ColorId == partEntry.Color.ID {
				collection.Parts[i].Part.ImageURL = partColors[j].ImageURL
				break
			}
		}
	}

	return nil
}
