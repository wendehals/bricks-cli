package api

import "github.com/wendehals/bricks/model"

// setListsPageResult contains the result of /api/v3/users/{user_token}/setlists/?page={page}
type setListsPageResult struct {
	Next    string          `json:"next"`
	Results []model.SetList `json:"results"`
}

// setsPageResult contains the result of /api/v3/users/{user_token}/sets/?page={page}
type setsPageResult struct {
	Next    string          `json:"next"`
	Results []model.UserSet `json:"results"`
}

// partListsPageResult contains the result of /api/v3/users/{user_token}/allparts/?page={page}
type partListsPageResult struct {
	Next    string           `json:"next"`
	Results []model.PartList `json:"results"`
}

/*
	partsPageResult contains the result of
	/api/v3/users/{user_token}/allparts/
	/api/v3/users/{user_token}/partlists/{list_id}/parts/
	/api/v3/lego/sets/{set_num}/parts/
*/
type partsPageResult struct {
	Next    string            `json:"next"`
	Results []model.PartEntry `json:"results"`
}

// colorsPageResult contains the result of /api/v3/lego/colors/?page={page}
type colorsPageResult struct {
	Next    string        `json:"next"`
	Results []model.Color `json:"results"`
}

// partColorsPageResult contains the result of /api/v3/lego/parts/{part_num}/colors/
type partColorsPageResult struct {
	Next    string            `json:"next"`
	Results []model.PartColor `json:"results"`
}

type invPart struct {
	Part    model.Part  `json:"part"`
	Color   model.Color `json:"color"`
	SetNum  string      `json:"set_num"`
	IsSpare bool        `json:"is_spare"`
}

type lostPart struct {
	ID       uint    `json:"lost_part_id"`
	Quantity int     `json:"lost_quantity"`
	InvPart  invPart `json:"inv_part"`
}

// lostPartsPageResult contains the result of /api/v3/users/{user_token}/lost_parts/
type lostPartsPageResult struct {
	Next    string     `json:"next"`
	Results []lostPart `json:"results"`
}

func (l *lostPartsPageResult) convertToPartEntries() []model.PartEntry {
	partEntries := make([]model.PartEntry, len(l.Results))
	for i, lostPart := range l.Results {
		partEntry := model.PartEntry{}
		partEntry.Quantity = lostPart.Quantity
		partEntry.Part = lostPart.InvPart.Part
		partEntry.Color = lostPart.InvPart.Color
		partEntry.IsSpare = lostPart.InvPart.IsSpare
		partEntry.SetNum = lostPart.InvPart.SetNum
		partEntries[i] = partEntry
	}

	return partEntries
}
