package api

import (
	"fmt"
	"log"
	"strings"

	"github.com/wendehals/bricks/model"
)

func RetrieveSetParts(bricksAPI *BricksAPI, setNum string, includeMiniFigs bool) *model.Collection {
	set := bricksAPI.GetSet(setNum)

	collection := bricksAPI.GetSetParts(setNum, includeMiniFigs)
	collection.Sets = append(collection.Sets, *set)

	return collection
}

func RetrieveUserSetParts(bricksAPI *BricksAPI, usersAPI *UsersAPI, setNum string, includeMiniFigs bool) *model.Collection {
	setParts := RetrieveSetParts(bricksAPI, setNum, includeMiniFigs)
	sets := setParts.Sets

	lostParts := usersAPI.GetLostParts()
	var setLostParts model.Collection
	for _, part := range lostParts.Parts {
		if part.SetNum == setNum {
			setLostParts.Parts = append(setLostParts.Parts, part)
		}
	}

	setParts.Subtract(&setLostParts)
	setParts.User = usersAPI.userName
	setParts.Comment = "User set parts"
	setParts.Sets = sets

	return setParts
}

func RetrieveSetListParts(bricksAPI *BricksAPI, usersAPI *UsersAPI, setListId uint, includeMiniFigs bool) []model.Collection {
	log.Printf("Retrieving parts of all sets from the set list %d", setListId)

	userSets := usersAPI.GetSetListSets(setListId)
	collections := []model.Collection{}
	for _, userSet := range userSets.Sets {
		collection := RetrieveSetParts(bricksAPI, userSet.Set.Number, includeMiniFigs)
		for i := 0; i < int(userSet.Quantity); i++ {
			collections = append(collections, *collection)
		}
	}

	return collections
}

func RetrievePartListParts(usersAPI *UsersAPI, listId uint) *model.Collection {
	partList := usersAPI.GetPartList(listId)
	collection := usersAPI.GetPartListParts(listId)
	collection.Comment = fmt.Sprintf("Part list %d - %s", listId, partList.Name)

	return collection
}

func RetrievePartListsParts(usersAPI *UsersAPI, partListsFile string, includeNonBuildable bool) []model.Collection {
	log.Printf("Retrieving parts from all part lists from the part lists file %s", partListsFile)

	partLists := model.Load(&model.PartLists{}, partListsFile)
	var collections []model.Collection
	for _, partList := range partLists.PartLists {
		if partList.IsBuildable || (includeNonBuildable && !partList.IsBuildable) {
			partListParts := usersAPI.GetPartListParts(partList.ID)
			partList := usersAPI.GetPartList(partList.ID)
			partListParts.Comment = strings.Join([]string{partListParts.Comment, partList.Name}, " ")
			collections = append(collections, *partListParts)
		}
	}

	return collections
}
