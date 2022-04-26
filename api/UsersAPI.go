package api

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/wendehals/bricks/model"
)

const (
	USERS_URL string = REBRICKABLE_BASE_URL + "users/%s/%s"

	GET_ALL_PARTS_LISTS_ERR_MSG string = "the user's part lists could not be retrieved: %s"
	GET_ALL_PARTS_ERR_MSG       string = "the list of all user's parts could not be retrieved: %s"
	GET_ALL_SET_LISTS_ERR_MSG   string = "the user's set lists could not be retrieved: %s"
	GET_ALL_SETS_ERR_MSG        string = "the user's sets could not be retrieved: %s"
	GET_PART_LIST_ERR_MSG       string = "the details about part list %d could not be retrieved: %s"
	GET_PART_LIST_PARTS_ERR_MSG string = "the parts of the user defined part list %d could not be retrieved: %s"
	GET_SET_LIST_ERR_MSG        string = "the details about set list %d could not be retrieved: %s"
	GET_SET_LIST_SETS_ERR_MSG   string = "the sets of list %d could not be retrieved: %s"
)

// UsersAPI provides API for accessing the user's data at Rebrickable
type UsersAPI struct {
	AbstractAPI
	userName string
	password string
	token    string
}

// NewUsersAPI creates a new object of UsersAPI and initializes it with a token by
// issueing a request to the Rebrickable API
func NewUsersAPI(client *http.Client, credentials *Credentials) *UsersAPI {
	usersAPI := UsersAPI{}
	usersAPI.client = client
	usersAPI.apiKey = credentials.APIKey
	usersAPI.userName = credentials.UserName
	usersAPI.password = credentials.Password

	err := usersAPI.postToken()
	if err != nil {
		log.Fatalf("Could not log into Rebrickable API: %s\n", err.Error())
	}

	return &usersAPI
}

// GetPartLists returns all part lists of the user provided by /api/v3/users/{user_token}/partlists/
func (u *UsersAPI) GetPartLists() (*model.PartLists, error) {
	partLists := model.PartLists{}
	partLists.User = u.userName

	var partListsPage *partListsPageResult

	err := u.requestPage(fmt.Sprintf(USERS_URL, u.token, "partlists"), &partListsPage)
	if err != nil {
		return nil, fmt.Errorf(GET_ALL_PARTS_LISTS_ERR_MSG, err.Error())
	}

	partLists.PartLists = partListsPage.Results
	for len(partListsPage.Next) > 0 {
		url := partListsPage.Next
		err = u.requestPage(url, &partListsPage)
		if err != nil {
			return nil, fmt.Errorf(GET_ALL_PARTS_LISTS_ERR_MSG, err.Error())
		}

		partLists.PartLists = append(partLists.PartLists, partListsPage.Results...)
	}

	return &partLists, nil
}

// GetAllParts returns all parts owned by the user provided by /api/v3/users/{user_token}/allparts
func (u *UsersAPI) GetAllParts() (*model.Collection, error) {
	collection := model.Collection{}
	collection.User = u.userName
	collection.IDs = []string{}
	collection.Names = []string{"All Parts"}

	allPartsPage := partsPageResult{}

	err := u.requestPage(fmt.Sprintf(USERS_URL, u.token, "allparts"), &allPartsPage)
	if err != nil {
		return nil, fmt.Errorf(GET_ALL_PARTS_ERR_MSG, err.Error())
	}

	collection.Parts = append(collection.Parts, allPartsPage.Results...)
	for len(allPartsPage.Next) > 0 {
		url := allPartsPage.Next
		allPartsPage = partsPageResult{}
		err = u.requestPage(url, &allPartsPage)
		if err != nil {
			return nil, fmt.Errorf(GET_ALL_PARTS_ERR_MSG, err.Error())
		}

		collection.Parts = append(collection.Parts, allPartsPage.Results...)
	}

	return &collection, nil
}

// GetSetLists returns all set lists of the user provided by /api/v3/users/{user_token}/setlists/
func (u *UsersAPI) GetSetLists() (*model.SetLists, error) {
	setLists := model.SetLists{}
	setLists.User = u.userName

	var setListsPage *setListsPageResult

	err := u.requestPage(fmt.Sprintf(USERS_URL, u.token, "setlists"), &setListsPage)
	if err != nil {
		return nil, fmt.Errorf(GET_ALL_SET_LISTS_ERR_MSG, err.Error())
	}

	setLists.SetLists = setListsPage.Results
	for len(setListsPage.Next) > 0 {
		url := setListsPage.Next
		err = u.requestPage(url, &setListsPage)
		if err != nil {
			return nil, fmt.Errorf(GET_ALL_SET_LISTS_ERR_MSG, err.Error())
		}

		setLists.SetLists = append(setLists.SetLists, setListsPage.Results...)
	}

	return &setLists, nil
}

// GetSets returns all sets of the user provided by /api/v3/users/{user_token}/sets
func (u *UsersAPI) GetSets() (*model.UserSets, error) {
	usersSets := model.UserSets{}
	usersSets.User = u.userName
	usersSets.ID = 0
	usersSets.Name = "All sets"

	var setsPage *setsPageResult

	err := u.requestPage(fmt.Sprintf(USERS_URL, u.token, "sets"), &setsPage)
	if err != nil {
		return nil, fmt.Errorf(GET_ALL_SETS_ERR_MSG, err.Error())
	}

	usersSets.Sets = setsPage.Results
	for len(setsPage.Next) > 0 {
		url := setsPage.Next
		err = u.requestPage(url, &setsPage)
		if err != nil {
			return nil, fmt.Errorf(GET_ALL_SETS_ERR_MSG, err.Error())
		}

		usersSets.Sets = append(usersSets.Sets, setsPage.Results...)
	}

	return &usersSets, nil
}

// PostToken posts /api/v3/users/_token/
func (u *UsersAPI) postToken() error {
	token := struct {
		Value string `json:"user_token"`
	}{}

	data := url.Values{
		"username": {u.userName},
		"password": {u.password},
	}

	reqest, err := createRequest(http.MethodPost, REBRICKABLE_BASE_URL+"users/_token/", u.apiKey, data)
	if err != nil {
		return err
	}

	reqest.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	err = doRequest(u.client, reqest, &token)
	if err != nil {
		return err
	}

	u.token = token.Value
	log.Printf("User token: %s\n", u.token)

	return nil
}

// GetSetList returns details about a certain set list of the user provided by
// /api/v3/users/{user_token}/setlists/{list_id}/
func (u *UsersAPI) GetSetList(listId uint) (*model.SetList, error) {
	var setList *model.SetList

	subPath := fmt.Sprintf("setlists/%d/", listId)
	err := u.requestPage(fmt.Sprintf(USERS_URL, u.token, subPath), &setList)
	if err != nil {
		return nil, fmt.Errorf(GET_SET_LIST_ERR_MSG, listId, err.Error())
	}

	return setList, nil
}

// GetSetListSets returns all sets of the user's set list provided by
// /api/v3/users/{user_token}/setlists/{list_id}/sets/
func (u *UsersAPI) GetSetListSets(listId uint) (*model.UserSets, error) {
	usersSets := model.UserSets{}
	usersSets.User = u.userName
	usersSets.ID = listId

	var setsPage *setsPageResult

	subPath := fmt.Sprintf("setlists/%d/sets/", listId)
	err := u.requestPage(fmt.Sprintf(USERS_URL, u.token, subPath), &setsPage)
	if err != nil {
		return nil, fmt.Errorf(GET_SET_LIST_SETS_ERR_MSG, listId, err.Error())
	}

	usersSets.Sets = setsPage.Results
	for len(setsPage.Next) > 0 {
		url := setsPage.Next
		err = u.requestPage(url, &setsPage)
		if err != nil {
			return nil, fmt.Errorf(GET_SET_LIST_SETS_ERR_MSG, listId, err.Error())
		}

		usersSets.Sets = append(usersSets.Sets, setsPage.Results...)
	}

	return &usersSets, nil
}

// GetPartList returns details about a certain set list of the user provided by
// /api/v3/users/{user_token}/partlists/{list_id}/
func (u *UsersAPI) GetPartList(listId uint) (*model.PartList, error) {
	var partList *model.PartList

	subPath := fmt.Sprintf("partlists/%d/", listId)
	err := u.requestPage(fmt.Sprintf(USERS_URL, u.token, subPath), &partList)
	if err != nil {
		return nil, fmt.Errorf(GET_PART_LIST_ERR_MSG, listId, err.Error())
	}

	return partList, nil
}

// GetPartListParts returns all parts of the user defined part list provided by
// /api/v3/users/{user_token}/partlists/{list_id}/parts
func (u *UsersAPI) GetPartListParts(listId uint) (*model.Collection, error) {
	collection := model.Collection{}
	collection.IDs = append(collection.IDs, fmt.Sprint(listId))

	partsPage := partsPageResult{}

	subPath := fmt.Sprintf("partlists/%d/parts", listId)
	err := u.requestPage(fmt.Sprintf(USERS_URL, u.token, subPath), &partsPage)
	if err != nil {
		return nil,
			fmt.Errorf(GET_PART_LIST_PARTS_ERR_MSG, listId, err.Error())
	}

	collection.Parts = append(collection.Parts, partsPage.Results...)
	for len(partsPage.Next) > 0 {
		url := partsPage.Next
		partsPage = partsPageResult{}
		err = u.requestPage(url, &partsPage)
		if err != nil {
			return nil,
				fmt.Errorf(GET_PART_LIST_PARTS_ERR_MSG, listId, err.Error())
		}

		collection.Parts = append(collection.Parts, partsPage.Results...)
	}

	return &collection, nil
}
