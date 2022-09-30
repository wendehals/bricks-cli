package api

import (
	"log"

	"github.com/wendehals/bricks/model"
)

func RetrieveSetParts(bricksAPI *BricksAPI, setNum string, includeMiniFigs bool) *model.Collection {
	set := bricksAPI.GetSet(setNum)

	collection := bricksAPI.GetSetParts(setNum, includeMiniFigs)
	collection.Sets = append(collection.Sets, *set)

	return collection
}

func RetrieveSetPartsWithoutLost(bricksAPI *BricksAPI, usersAPI *UsersAPI, setNum string, includeMiniFigs bool) *model.Collection {
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
	setParts.Sets = sets

	return setParts
}

func RetrieveSetListParts(bricksAPI *BricksAPI, usersAPI *UsersAPI, setListId uint, includeMiniFigs bool) []*model.Collection {
	log.Printf("Retrieving parts of all sets from the set list %d", setListId)

	userSets := usersAPI.GetSetListSets(setListId)
	collections := make([]*model.Collection, len(userSets.Sets))
	for i, userSet := range userSets.Sets {
		collection := RetrieveSetParts(bricksAPI, userSet.Set.Number, includeMiniFigs)
		collections[i] = collection
	}

	return collections
}

func RetrievePartListParts(usersAPI *UsersAPI, partListsFile string, includeNonBuildable bool) []*model.Collection {
	log.Printf("Retrieving parts of all part lists from the part lists file %s", partListsFile)

	partLists := model.Load(&model.PartLists{}, partListsFile)
	var collections []*model.Collection
	for _, partList := range partLists.PartLists {
		if partList.IsBuildable || (includeNonBuildable && !partList.IsBuildable) {
			partListParts := usersAPI.GetPartListParts(partList.ID)
			partList := usersAPI.GetPartList(partList.ID)
			partListParts.Names = append(partListParts.Names, partList.Name)
			collections = append(collections, partListParts)
		}
	}

	return collections
}
