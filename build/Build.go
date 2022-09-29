package build

import (
	"fmt"

	"github.com/wendehals/bricks/api"
	"github.com/wendehals/bricks/export"
	"github.com/wendehals/bricks/model"
	"github.com/wendehals/bricks/utils"
)

func Build(neededCollection *model.Collection, providedCollection *model.Collection, outputDir string,
	bricksAPI *api.BricksAPI, verbose bool) {
	colors := bricksAPI.GetColors()
	partRelationships := model.Load(&model.PartRelationships{}, utils.GetPartRelationshipsPath())

	buildCollection := build(neededCollection, providedCollection, &colors, partRelationships)
	buildCollection.Sort()

	model.Save(buildCollection, fmt.Sprintf("%s/result.build", outputDir))
	export.ExportBuildToHTML(buildCollection, outputDir, "build")

	providedCollection.RemoveQuantityZero()
	model.Save(providedCollection, fmt.Sprintf("%s/remaining.parts", outputDir))
	export.ExportCollectionToHTML(providedCollection, outputDir, "remaining")
}

func build(neededCollection *model.Collection, providedCollection *model.Collection, colors *[]model.Color,
	partRelationships *model.PartRelationships) *model.BuildCollection {

	buildCollection := model.BuildCollection{}
	if len(neededCollection.Sets) > 0 {
		buildCollection.Set = neededCollection.Sets[0]
	}
	buildCollection.Parts = []model.PartEntryMapping{}

	isMold := func(s1, s2 string) bool {
		return partRelationships.IsMold(s1, s2)
	}

	isPrint := func(s1, s2 string) bool {
		return partRelationships.IsPrint(s1, s2)
	}

	isAlternative := func(s1, s2 string) bool {
		return partRelationships.IsAlternative(s1, s2)
	}

	// filter spare parts
	neededCollection.Filter(func(p model.PartEntry) bool {
		return !p.IsSpare
	})

	// search for same part type, same color
	for _, neededPartEntry := range neededCollection.Parts {
		partEntryMapping := model.PartEntryMapping{}
		partEntryMapping.Quantity = neededPartEntry.Quantity
		partEntryMapping.Original = *model.DeepClone(&neededPartEntry, &model.PartEntry{})
		partEntryMapping.Substitutes = []model.PartEntry{}

		mapPartEntry(&partEntryMapping, providedCollection, neededPartEntry.Color.ID, utils.Equals)

		buildCollection.Parts = append(buildCollection.Parts, partEntryMapping)
	}

	// search for equivalent part type, same color
	for i := range buildCollection.Parts {
		partEntryMapping := &buildCollection.Parts[i]

		mapPartEntry(partEntryMapping, providedCollection, partEntryMapping.Original.Color.ID, isMold)
		mapPartEntry(partEntryMapping, providedCollection, partEntryMapping.Original.Color.ID, isPrint)
		mapPartEntry(partEntryMapping, providedCollection, partEntryMapping.Original.Color.ID, isAlternative)
	}

	// search for same or equivalent part type but different color
	for i := range buildCollection.Parts {
		partEntryMapping := &buildCollection.Parts[i]
		for j := 0; j < len(*colors) && partEntryMapping.Quantity > 0; j++ {
			colorId := (*colors)[j].ID
			if partEntryMapping.Original.Color.ID != colorId {
				mapPartEntry(partEntryMapping, providedCollection, colorId, utils.Equals)
				mapPartEntry(partEntryMapping, providedCollection, colorId, isMold)
				mapPartEntry(partEntryMapping, providedCollection, colorId, isPrint)
				mapPartEntry(partEntryMapping, providedCollection, colorId, isAlternative)
			}
		}
	}

	return &buildCollection
}

func mapPartEntry(partEntryMapping *model.PartEntryMapping, providedCollection *model.Collection,
	colorId int, eqFunc func(string, string) bool) {

	if partEntryMapping.Quantity > 0 {
		providedPartEntry := providedCollection.FindByEquivalence(
			partEntryMapping.Original.Part.Number, colorId, eqFunc)

		if providedPartEntry != nil && providedPartEntry.Quantity > 0 {
			mappedPartEntry := model.DeepClone(providedPartEntry, &model.PartEntry{})
			if providedPartEntry.Quantity >= partEntryMapping.Quantity {
				providedPartEntry.Quantity -= partEntryMapping.Quantity
				mappedPartEntry.Quantity = partEntryMapping.Quantity
			} else {
				providedPartEntry.Quantity = 0
			}

			partEntryMapping.Quantity -= mappedPartEntry.Quantity
			partEntryMapping.Substitutes = append(partEntryMapping.Substitutes, *mappedPartEntry)

			providedCollection.Remove(providedPartEntry.Part.Number,
				providedPartEntry.Color.ID, providedPartEntry.Quantity)
		}
	}
}
