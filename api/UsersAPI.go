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

	GET_SETS_ERR_MSG            string = "the user's sets could not be retrieved: %s"
	GET_SET_LISTS_ERR_MSG       string = "the user's set lists could not be retrieved: %s"
	GET_PARTS_LISTS_ERR_MSG     string = "the user's part lists could not be retrieved: %s"
	GET_ALL_PARTS_ERR_MSG       string = "the list of all user's parts could not be retrieved: %s"
	GET_PART_LIST_PARTS_ERR_MSG string = "the parts of the user defined part list %d could not be retrieved: %s"
)

// UsersAPI provides API for accessing a Rebrickable user's data
type UsersAPI struct {
	AbstractAPI
	token string
}

// NewUsersAPI creates a new object of UsersAPI and initializes it with a token by issueing a request to the Rebrickable API
func NewUsersAPI(client *http.Client, credentials *Credentials) *UsersAPI {
	users := UsersAPI{}
	users.client = client
	users.apiKey = credentials.APIKey

	err := users.postToken(credentials.Username, credentials.Password)
	if err != nil {
		log.Fatalf("Could not log into Rebrickable API: %s\n", err.Error())
	}

	return &users
}

// PostToken posts /api/v3/users/_token/
func (u *UsersAPI) postToken(userName string, password string) error {
	token := struct {
		Value string `json:"user_token"`
	}{}

	data := url.Values{
		"username": {userName},
		"password": {password},
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

// GetSets returns all sets of a user provided by /api/v3/users/{user_token}/sets
func (u *UsersAPI) GetSets() (*model.UsersSets, error) {
	usersSets := model.UsersSets{}
	var setsPage *setsPageResult

	err := u.requestPage(fmt.Sprintf(USERS_URL, u.token, "sets"), &setsPage)
	if err != nil {
		return nil, fmt.Errorf(GET_SETS_ERR_MSG, err.Error())
	}

	usersSets.Sets = setsPage.Results
	for len(setsPage.Next) > 0 {
		url := setsPage.Next
		err = u.requestPage(url, &setsPage)
		if err != nil {
			return nil, fmt.Errorf(GET_SETS_ERR_MSG, err.Error())
		}

		usersSets.Sets = append(usersSets.Sets, setsPage.Results...)
	}

	return &usersSets, nil
}

// GetSetLists returns all owned sets of a user provided by /api/v3/users/{user_token}/setlists/
func (u *UsersAPI) GetSetLists() (*model.SetLists, error) {
	setLists := model.SetLists{}
	var setListsPage *setListsPageResult

	err := u.requestPage(fmt.Sprintf(USERS_URL, u.token, "setlists"), &setListsPage)
	if err != nil {
		return nil, fmt.Errorf(GET_SET_LISTS_ERR_MSG, err.Error())
	}

	setLists.SetLists = setListsPage.Results
	for len(setListsPage.Next) > 0 {
		url := setListsPage.Next
		err = u.requestPage(url, &setListsPage)
		if err != nil {
			return nil, fmt.Errorf(GET_SET_LISTS_ERR_MSG, err.Error())
		}

		setLists.SetLists = append(setLists.SetLists, setListsPage.Results...)
	}

	return &setLists, nil
}

// PartLists returns the result of /api/v3/users/{user_token}/partlists/
func (u *UsersAPI) GetPartLists() (*model.PartLists, error) {
	partLists := model.PartLists{}
	var partListsPage *partListsPageResult

	err := u.requestPage(fmt.Sprintf(USERS_URL, u.token, "partlists"), &partListsPage)
	if err != nil {
		return nil, fmt.Errorf(GET_PARTS_LISTS_ERR_MSG, err.Error())
	}

	partLists.PartLists = partListsPage.Results
	for len(partListsPage.Next) > 0 {
		url := partListsPage.Next
		err = u.requestPage(url, &partListsPage)
		if err != nil {
			return nil, fmt.Errorf(GET_PARTS_LISTS_ERR_MSG, err.Error())
		}

		partLists.PartLists = append(partLists.PartLists, partListsPage.Results...)
	}

	return &partLists, nil
}

// GetAllParts returns all parts owned by a user provided by /api/v3/users/{user_token}/allparts
func (u *UsersAPI) GetAllParts() (*model.Collection, error) {
	collection := model.Collection{}
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

// GetPartListParts returns all parts of a user defined parts list provided by /api/v3/users/{user_token}/partlists/{list_id}/parts
func (u *UsersAPI) GetPartListParts(listId uint) (*model.Collection, error) {
	collection := model.Collection{}
	collection.IDs = append(collection.IDs, fmt.Sprint(listId))

	partsPage := partsPageResult{}

	err := u.requestPage(fmt.Sprintf(USERS_URL, u.token, fmt.Sprintf("partlists/%d/parts", listId)), &partsPage)
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
