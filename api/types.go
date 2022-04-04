package api

import "github.com/wendehals/bricks/model"

// setListsPageResult contains the result of /api/v3/users/{user_token}/setlists/?page={page}
type setListsPageResult struct {
	Next    string               `json:"next"`
	Results []model.SetListEntry `json:"results"`
}

// setsPageResult contains the result of /api/v3/users/{user_token}/sets/?page={page}
type setsPageResult struct {
	Next    string           `json:"next"`
	Results []model.UsersSet `json:"results"`
}

// partListsPageResult contains the result of /api/v3/users/{user_token}/allparts/?page={page}
type partListsPageResult struct {
	Next    string                `json:"next"`
	Results []model.PartListEntry `json:"results"`
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

// partColorsPageResult contains the result of /api/v3/lego/parts/{part_num}/colors/
type partColorsPageResult struct {
	Next    string            `json:"next"`
	Results []model.PartColor `json:"results"`
}
