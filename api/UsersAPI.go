package api

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/wendehals/bricks/model"
)

const usersURL string = rebrickableBaseURL + "users/%s/%s"

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

	reqest, err := createRequest(http.MethodPost, rebrickableBaseURL+"users/_token/", u.apiKey, data)
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
func (u *UsersAPI) GetSets() ([]model.UsersSet, error) {
	usersSets := []model.UsersSet{}
	setsPage := setsPageResult{}

	err := u.requestPage(fmt.Sprintf(usersURL, u.token, "sets"), &setsPage)
	if err != nil {
		return nil, fmt.Errorf("user's sets could not be retrieved: %s", err.Error())
	}

	usersSets = append(usersSets, setsPage.Results...)
	for len(setsPage.Next) > 0 {
		url := setsPage.Next
		setsPage := setsPageResult{}
		err = u.requestPage(url, &setsPage)
		if err != nil {
			return nil, fmt.Errorf("user's sets could not be retrieved: %s", err.Error())
		}

		usersSets = append(usersSets, setsPage.Results...)
	}

	return usersSets, nil
}

// GetSetLists returns all owned sets of a uses provided by /api/v3/users/{user_token}/setlists/
func (u *UsersAPI) GetSetLists() ([]model.SetListEntry, error) {
	var setList []model.SetListEntry
	setListsPage := setListsPageResult{}

	err := u.requestPage(fmt.Sprintf(usersURL, u.token, "setlists"), &setListsPage)
	if err != nil {
		return nil, fmt.Errorf("the user's set lists could not be retrieved: %s", err.Error())
	}

	setList = append(setList, setListsPage.Results...)
	for len(setListsPage.Next) > 0 {
		url := setListsPage.Next
		setListsPage := setListsPageResult{}
		err = u.requestPage(url, &setListsPage)
		if err != nil {
			return nil, fmt.Errorf("the user's set lists could not be retrieved: %s", err.Error())
		}

		setList = append(setList, setListsPage.Results...)
	}

	return setList, nil
}

// PartLists returns the result of /api/v3/users/{user_token}/partlists/
func (u *UsersAPI) GetPartLists() ([]model.PartListEntry, error) {
	var partListEntry []model.PartListEntry
	partListsPage := partListsPageResult{}

	err := u.requestPage(fmt.Sprintf(usersURL, u.token, "partlists"), &partListsPage)
	if err != nil {
		return nil, fmt.Errorf("the user's part lists could not be retrieved: %s", err.Error())
	}

	partListEntry = append(partListEntry, partListsPage.Results...)
	for len(partListsPage.Next) > 0 {
		url := partListsPage.Next
		setListsPage := setListsPageResult{}
		err = u.requestPage(url, &setListsPage)
		if err != nil {
			return nil, fmt.Errorf("the user's part lists could not be retrieved: %s", err.Error())
		}

		partListEntry = append(partListEntry, partListsPage.Results...)
	}

	return partListEntry, nil
}

// GetAllParts returns all parts owned by a user provided by /api/v3/users/{user_token}/allparts
func (u *UsersAPI) GetAllParts() (*model.Collection, error) {
	collection := model.Collection{}
	allPartsPage := partsPageResult{}

	err := u.requestPage(fmt.Sprintf(usersURL, u.token, "allparts"), &allPartsPage)
	if err != nil {
		return nil, fmt.Errorf("the list of all user's parts could not be retrieved: %s", err.Error())
	}

	collection.Parts = append(collection.Parts, allPartsPage.Results...)
	for len(allPartsPage.Next) > 0 {
		url := allPartsPage.Next
		allPartsPage = partsPageResult{}
		err = u.requestPage(url, &allPartsPage)
		if err != nil {
			return nil, fmt.Errorf("the list of all user's parts could not be retrieved: %s", err.Error())
		}

		collection.Parts = append(collection.Parts, allPartsPage.Results...)
	}

	return &collection, nil
}

// GetPartListParts returns all parts of a user defined parts list provided by /api/v3/users/{user_token}/partlists/{list_id}/parts
func (u *UsersAPI) GetPartListParts(listId string) (*model.Collection, error) {
	collection := model.Collection{}
	partsPage := partsPageResult{}

	err := u.requestPage(fmt.Sprintf(usersURL, u.token, fmt.Sprintf("partlists/%s/parts", listId)), &partsPage)
	if err != nil {
		return nil,
			fmt.Errorf("the parts of the user defined parts list %s could not be retrieved: %s", listId, err.Error())
	}

	collection.Parts = append(collection.Parts, partsPage.Results...)
	for len(partsPage.Next) > 0 {
		url := partsPage.Next
		partsPage = partsPageResult{}
		err = u.requestPage(url, &partsPage)
		if err != nil {
			return nil,
				fmt.Errorf("the parts of the user defined parts list %s could not be retrieved: %s", listId, err.Error())
		}

		collection.Parts = append(collection.Parts, partsPage.Results...)
	}

	return &collection, nil
}
