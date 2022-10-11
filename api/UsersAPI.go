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

	GET_ALL_PARTS_ERR_MSG       string = "the list of all user's parts could not be retrieved: %s"
	GET_SETS_ERR_MSG            string = "the user's sets could not be retrieved: %s"
	GET_SET_LIST_ERR_MSG        string = "the details about set list %d could not be retrieved: %s"
	GET_SET_LISTS_ERR_MSG       string = "the user's set lists could not be retrieved: %s"
	GET_SET_LIST_SETS_ERR_MSG   string = "the sets of list %d could not be retrieved: %s"
	GET_PART_LIST_ERR_MSG       string = "the details about part list %d could not be retrieved: %s"
	GET_PART_LISTS_ERR_MSG      string = "the user's part lists could not be retrieved: %s"
	GET_PART_LIST_PARTS_ERR_MSG string = "the parts of the user defined part list %d could not be retrieved: %s"
	GET_LOST_PARTS_ERR_MSG      string = "the list of the user's lost parts could not be retrieved: %s"
	TOKEN_ERR_MSG               string = "could not log into Rebrickable API: %s"
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
func NewUsersAPI(client *http.Client, credentials *Credentials, verbose bool) *UsersAPI {
	usersAPI := UsersAPI{}
	usersAPI.client = client
	usersAPI.apiKey = credentials.APIKey
	usersAPI.userName = credentials.UserName
	usersAPI.password = credentials.Password
	usersAPI.verbose = verbose

	usersAPI.postToken()

	return &usersAPI
}

// GetAllParts returns all parts owned by the user provided by /api/v3/users/{user_token}/allparts/
func (u *UsersAPI) GetAllParts() *model.Collection {
	log.Printf("Retrieving all parts owned by user %s", u.userName)

	collection := model.NewCollection()
	collection.User = u.userName
	collection.Names = []string{"All Parts"}

	allPartsPage := partsPageResult{}

	err := u.requestPage(fmt.Sprintf(USERS_URL, u.token, "allparts?page_size=200"), &allPartsPage)
	if err != nil {
		log.Fatalf(GET_ALL_PARTS_ERR_MSG, err.Error())
	}

	collection.Parts = append(collection.Parts, allPartsPage.Results...)
	for len(allPartsPage.Next) > 0 {
		url := allPartsPage.Next
		allPartsPage = partsPageResult{}
		err = u.requestPage(url, &allPartsPage)
		if err != nil {
			log.Fatalf(GET_ALL_PARTS_ERR_MSG, err.Error())
		}

		collection.Parts = append(collection.Parts, allPartsPage.Results...)
	}

	return collection
}

// GetSets returns all sets of the user provided by /api/v3/users/{user_token}/sets/
func (u *UsersAPI) GetSets() *model.UserSets {
	log.Printf("Retrieving all sets owned by user %s", u.userName)

	usersSets := model.UserSets{}
	usersSets.User = u.userName
	usersSets.ID = 0
	usersSets.Name = "All sets"

	var setsPage *setsPageResult

	err := u.requestPage(fmt.Sprintf(USERS_URL, u.token, "sets"), &setsPage)
	if err != nil {
		log.Fatalf(GET_SETS_ERR_MSG, err.Error())
	}

	usersSets.Sets = setsPage.Results
	for len(setsPage.Next) > 0 {
		url := setsPage.Next
		err = u.requestPage(url, &setsPage)
		if err != nil {
			log.Fatalf(GET_SETS_ERR_MSG, err.Error())
		}

		usersSets.Sets = append(usersSets.Sets, setsPage.Results...)
	}

	return &usersSets
}

// GetSetList returns details about a certain set list of the user provided by
// /api/v3/users/{user_token}/setlists/{list_id}/
func (u *UsersAPI) GetSetList(listId uint) *model.SetList {
	log.Printf("Retrieving details about set list %d", listId)

	var setList *model.SetList

	subPath := fmt.Sprintf("setlists/%d/", listId)
	err := u.requestPage(fmt.Sprintf(USERS_URL, u.token, subPath), &setList)
	if err != nil {
		log.Fatalf(GET_SET_LIST_ERR_MSG, listId, err.Error())
	}

	return setList
}

// GetSetLists returns all set lists of the user provided by /api/v3/users/{user_token}/setlists/
func (u *UsersAPI) GetSetLists() *model.SetLists {
	log.Printf("Retrieving set lists of user %s", u.userName)

	setLists := model.SetLists{}
	setLists.User = u.userName

	var setListsPage *setListsPageResult

	err := u.requestPage(fmt.Sprintf(USERS_URL, u.token, "setlists"), &setListsPage)
	if err != nil {
		log.Fatalf(GET_SET_LISTS_ERR_MSG, err.Error())
	}

	setLists.SetLists = setListsPage.Results
	for len(setListsPage.Next) > 0 {
		url := setListsPage.Next
		err = u.requestPage(url, &setListsPage)
		if err != nil {
			log.Fatalf(GET_SET_LISTS_ERR_MSG, err.Error())
		}

		setLists.SetLists = append(setLists.SetLists, setListsPage.Results...)
	}

	return &setLists
}

// GetSetListSets returns all sets of the user's set list provided by
// /api/v3/users/{user_token}/setlists/{list_id}/sets/
func (u *UsersAPI) GetSetListSets(listId uint) *model.UserSets {
	log.Printf("Retrieving all sets of set list %d", listId)

	usersSets := model.UserSets{}
	usersSets.User = u.userName
	usersSets.ID = listId

	var setsPage *setsPageResult

	subPath := fmt.Sprintf("setlists/%d/sets/", listId)
	err := u.requestPage(fmt.Sprintf(USERS_URL, u.token, subPath), &setsPage)
	if err != nil {
		log.Fatalf(GET_SET_LIST_SETS_ERR_MSG, listId, err.Error())
	}

	usersSets.Sets = setsPage.Results
	for len(setsPage.Next) > 0 {
		url := setsPage.Next
		err = u.requestPage(url, &setsPage)
		if err != nil {
			log.Fatalf(GET_SET_LIST_SETS_ERR_MSG, listId, err.Error())
		}

		usersSets.Sets = append(usersSets.Sets, setsPage.Results...)
	}

	return &usersSets
}

// GetPartList returns details about a certain set list of the user provided by
// /api/v3/users/{user_token}/partlists/{list_id}/
func (u *UsersAPI) GetPartList(listId uint) *model.PartList {
	log.Printf("Retrieving details about part list %d", listId)

	var partList *model.PartList

	subPath := fmt.Sprintf("partlists/%d/", listId)
	err := u.requestPage(fmt.Sprintf(USERS_URL, u.token, subPath), &partList)
	if err != nil {
		log.Fatalf(GET_PART_LIST_ERR_MSG, listId, err.Error())
	}

	return partList
}

// GetPartLists returns all part lists of the user provided by /api/v3/users/{user_token}/partlists/
func (u *UsersAPI) GetPartLists() *model.PartLists {
	log.Printf("Retrieving part lists of user %s", u.userName)

	partLists := model.PartLists{}
	partLists.User = u.userName

	var partListsPage *partListsPageResult

	err := u.requestPage(fmt.Sprintf(USERS_URL, u.token, "partlists"), &partListsPage)
	if err != nil {
		log.Fatalf(GET_PART_LISTS_ERR_MSG, err.Error())
	}

	partLists.PartLists = partListsPage.Results
	for len(partListsPage.Next) > 0 {
		url := partListsPage.Next
		err = u.requestPage(url, &partListsPage)
		if err != nil {
			log.Fatalf(GET_PART_LISTS_ERR_MSG, err.Error())
		}

		partLists.PartLists = append(partLists.PartLists, partListsPage.Results...)
	}

	return &partLists
}

// GetPartListParts returns all parts of the user defined part list provided by
// /api/v3/users/{user_token}/partlists/{list_id}/parts/
func (u *UsersAPI) GetPartListParts(listId uint) *model.Collection {
	log.Printf("Retrieving parts of part list %d", listId)

	collection := model.NewCollection()
	collection.Names = append(collection.Names, fmt.Sprint(listId))

	partsPage := partsPageResult{}

	subPath := fmt.Sprintf("partlists/%d/parts", listId)
	err := u.requestPage(fmt.Sprintf(USERS_URL, u.token, subPath), &partsPage)
	if err != nil {
		log.Fatalf(GET_PART_LIST_PARTS_ERR_MSG, listId, err.Error())
	}

	collection.Parts = append(collection.Parts, partsPage.Results...)
	for len(partsPage.Next) > 0 {
		url := partsPage.Next
		partsPage = partsPageResult{}
		err = u.requestPage(url, &partsPage)
		if err != nil {
			log.Fatalf(GET_PART_LIST_PARTS_ERR_MSG, listId, err.Error())
		}

		collection.Parts = append(collection.Parts, partsPage.Results...)
	}

	return collection
}

// GetLostParts returns all parts owned by the user provided by /api/v3/users/{user_token}/lost_parts/
func (u *UsersAPI) GetLostParts() *model.Collection {
	log.Printf("Retrieving lost parts of user %s", u.userName)

	lostParts := model.NewCollection()
	lostParts.User = u.userName
	lostParts.Names = append(lostParts.Names, "Lost parts")

	lostPartsPage := lostPartsPageResult{}

	err := u.requestPage(fmt.Sprintf(USERS_URL, u.token, "lost_parts"), &lostPartsPage)
	if err != nil {
		log.Fatalf(GET_LOST_PARTS_ERR_MSG, err.Error())
	}

	lostParts.Parts = append(lostParts.Parts, lostPartsPage.convertToParts()...)
	for len(lostPartsPage.Next) > 0 {
		url := lostPartsPage.Next
		lostPartsPage = lostPartsPageResult{}
		err = u.requestPage(url, &lostPartsPage)
		if err != nil {
			log.Fatalf(GET_LOST_PARTS_ERR_MSG, err.Error())
		}

		lostParts.Parts = append(lostParts.Parts, lostPartsPage.convertToParts()...)
	}

	return lostParts
}

// postToken posts /api/v3/users/_token/
func (u *UsersAPI) postToken() {
	token := struct {
		Value string `json:"user_token"`
	}{}

	data := url.Values{
		"username": {u.userName},
		"password": {u.password},
	}

	reqest, err := u.createRequest(http.MethodPost, REBRICKABLE_BASE_URL+"users/_token/", data)
	if err != nil {
		log.Fatalf(TOKEN_ERR_MSG, err.Error())
	}

	reqest.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	err = u.doRequest(reqest, &token)
	if err != nil {
		log.Fatalf(TOKEN_ERR_MSG, err.Error())
	}

	u.token = token.Value
	if u.verbose {
		log.Printf("User token: %s", u.token)
	}
}
