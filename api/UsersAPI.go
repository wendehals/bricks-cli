package api

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/wendehals/bricks/model"
)

const usersURL string = RebrickableBaseURL + "users/%s/%s"

// UsersAPI provides API for accessing a Rebrickable user's data
type UsersAPI struct {
	AbstractAPI
	token string
}

// NewUsersAPI creates a new object of UsersAPI and initializes it with a token by issueing a request to the Rebrickable API
func NewUsersAPI(client *http.Client, userName string, password string, apiKey string) *UsersAPI {
	users := UsersAPI{}
	users.client = client
	users.apiKey = apiKey

	err := users.postToken(userName, password)
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

	reqest, err := CreateRequest(http.MethodPost, RebrickableBaseURL+"users/_token/", u.apiKey, data)
	if err != nil {
		return err
	}

	reqest.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	err = DoRequest(u.client, reqest, &token)
	if err != nil {
		return err
	}

	u.token = token.Value
	log.Printf("User token: %s\n", u.token)

	return nil
}

// setListsPageResult contains the result of /api/v3/users/{user_token}/setlists/?page={page}
type setListsPageResult struct {
	Next    string               `json:"next"`
	Results []model.SetListEntry `json:"results"`
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

// allPartsPageResult contains the result of /api/v3/users/{user_token}/allparts/?page={page}
type partListsResult struct {
	Next    string                `json:"next"`
	Results []model.PartListEntry `json:"results"`
}

// PartLists returns the result of /api/v3/users/{user_token}/partlists/
func (u *UsersAPI) GetPartLists() ([]model.PartListEntry, error) {
	var partListEntry []model.PartListEntry
	partListsPage := partListsResult{}

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

// allPartsPageResult contains the result of /api/v3/users/{user_token}/allparts/?page={page}
type allPartsPageResult struct {
	Next    string            `json:"next"`
	Results []model.PartEntry `json:"results"`
}

// GetAllParts returns all owned parts of a user provided by /api/v3/users/{user_token}/allparts
func (u *UsersAPI) GetAllParts() (*model.Collection, error) {
	collection := model.Collection{}
	allPartsPage := allPartsPageResult{}

	err := u.requestPage(fmt.Sprintf(usersURL, u.token, "allparts"), &allPartsPage)
	if err != nil {
		return nil, fmt.Errorf("the list of all user's parts could not be retrieved: %s", err.Error())
	}

	collection.Parts = append(collection.Parts, allPartsPage.Results...)
	for len(allPartsPage.Next) > 0 {
		url := allPartsPage.Next
		allPartsPage = allPartsPageResult{}
		err = u.requestPage(url, &allPartsPage)
		if err != nil {
			return nil, fmt.Errorf("the list of all user's parts could not be retrieved: %s", err.Error())
		}

		collection.Parts = append(collection.Parts, allPartsPage.Results...)
	}

	return &collection, nil
}

// setsPageResult contains the result of /api/v3/users/{user_token}/sets/?page={page}
type setsPageResult struct {
	Next    string           `json:"next"`
	Results []model.UsersSet `json:"results"`
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
