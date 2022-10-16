package services

import (
	"log"

	"github.com/wendehals/bricks/model"
	"github.com/wendehals/bricks/utils"
)

// MergeAllCollections merges all given collections to a new single collection.
func MergeAllCollections(collections []model.Collection) *model.Collection {
	log.Println("Merging parts of collections")

	mergedCollection := model.NewCollection()
	for _, collection := range collections {
		mergedCollection.Add(&collection)
	}

	return mergedCollection
}

// MergeByColor merges all parts of the same shape ignoring the color.
// The Color field of Part will be invalid afterwards and set to unknown color.
// The IsSpare flag of Part will be invalid afterwards and set to false for all parts.
func MergeByColor(collection *model.Collection) {
	collection.Sets = []model.Set{}
	collection.Comment = "Merged by color"

	partsMap := collection.MapPartsByPartNumber(utils.Identity)

	newParts := []model.Part{}
	for _, value := range partsMap {
		var currentPart = value[0]
		currentPart.Color = model.UNKNOWN_COLOR

		for i := 1; i < len(value); i++ {
			currentPart.Quantity += value[i].Quantity
		}

		newParts = append(newParts, currentPart)
	}

	collection.Parts = newParts
}

// MergeByPrint merges all parts by using their representative in terms of prints.
// The Color field of Part will be invalid afterwards and set to unknown color.
// The IsSpare flag of Part will be invalid afterwards and set to false for all parts.
func MergeByPrint(collection *model.Collection) {
	collection.Sets = []model.Set{}
	collection.Comment = "Merged by prints"

	representative := func(part string) string {
		return GetPrints(false).RepresentativeFor(part)
	}
	mergeByRepresentative(collection, representative)
}

// MergeByMold merges all parts by using their representative in terms of molds.
// The Color field of Part will be invalid afterwards and set to unknown color.
// The IsSpare flag of Part will be invalid afterwards and set to false for all parts.
func MergeByMold(collection *model.Collection) {
	collection.Sets = []model.Set{}
	collection.Comment = "Merged by molds"

	representative := func(part string) string {
		return GetMolds(false).RepresentativeFor(part)
	}
	mergeByRepresentative(collection, representative)
}

// MergeByAlternate merges all parts by using their representative in terms of alternates.
// The Color field of Part will be invalid afterwards and set to unknown color.
// The IsSpare flag of Part will be invalid afterwards and set to false for all parts.
func MergeByAlternate(collection *model.Collection) {
	collection.Sets = []model.Set{}
	collection.Comment = "Merged by alternates"

	representative := func(part string) string {
		return GetAlternates(false).RepresentativeFor(part)
	}
	mergeByRepresentative(collection, representative)
}

func mergeByRepresentative(collection *model.Collection, representative func(string) string) {
	collection.Sets = []model.Set{}

	partsMap := collection.MapPartsByPartNumber(representative)

	shapes := GetShapes(false)

	newParts := []model.Part{}
	for key, value := range partsMap {
		representative := model.NewPart()
		representative.Shape = shapes.Shapes[key]
		representative.Color = model.UNKNOWN_COLOR

		for i := 0; i < len(value); i++ {
			representative.Quantity += value[i].Quantity
		}

		newParts = append(newParts, *representative)
	}

	collection.Parts = newParts
}
