package build

import (
	"fmt"

	"github.com/wendehals/bricks/api"
	"github.com/wendehals/bricks/model"
	"github.com/wendehals/bricks/utils"
)

func Build(neededCollection *model.Collection, providedCollection *model.Collection, outputDir string, bricksAPI *api.BricksAPI, verbose bool) {
	colors := bricksAPI.GetColors()
	partRelationships := model.Load(&model.PartRelationships{}, utils.GetPartRelationshipsPath())
	buildCollection := build(neededCollection, providedCollection, &colors, partRelationships)

	neededCollection.RemoveQuantityZero()
	providedCollection.RemoveQuantityZero()

	partEntries := []model.PartEntry{}
	for _, partEntry := range buildCollection.Parts {
		partEntries = append(partEntries, partEntry...)
	}

	resultCollection := model.Collection{}
	resultCollection.Parts = partEntries

	neededCollection.ExportToHTML(outputDir)
	providedCollection.ExportToHTML(outputDir)
	resultCollection.ExportToHTML(outputDir)

	model.Save(neededCollection, fmt.Sprintf("%s/missing.parts", outputDir))
	model.Save(providedCollection, fmt.Sprintf("%s/remaining.parts", outputDir))
	model.Save(resultCollection, fmt.Sprintf("%s/result.parts", outputDir))
}

func build(neededParts *model.Collection, providedParts *model.Collection, colors *[]model.Color, partRelationships *model.PartRelationships) *BuildMapping {
	buildMapping := NewBuildMapping()

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
	neededParts.Filter(func(p model.PartEntry) bool {
		return !p.IsSpare
	})

	// search for same part type, same color
	for i := range neededParts.Parts {
		neededPartEntry := &neededParts.Parts[i]
		providedPartEntry := providedParts.Find(neededPartEntry.Part.Number, neededPartEntry.Color.ID)
		mapPartEntry(neededPartEntry, providedPartEntry, buildMapping)
	}

	// search for mold/print/alternative, same color
	for i := range neededParts.Parts {
		neededPartEntry := &neededParts.Parts[i]
		if neededPartEntry.Quantity > 0 {
			number := neededPartEntry.Part.Number
			colorId := neededPartEntry.Color.ID

			mapPartEntry(neededPartEntry, providedParts.FindByEquality(number, colorId, isMold), buildMapping)
			mapPartEntry(neededPartEntry, providedParts.FindByEquality(number, colorId, isPrint), buildMapping)
			mapPartEntry(neededPartEntry, providedParts.FindByEquality(number, colorId, isAlternative), buildMapping)
		}
	}

	// search for same part type but different color
	for i := range neededParts.Parts {
		neededPartEntry := &neededParts.Parts[i]
		for j := 0; j < len(*colors) && neededPartEntry.Quantity > 0; j++ {
			colorId := (*colors)[j].ID
			if neededPartEntry.Color.ID != colorId {
				number := neededPartEntry.Part.Number
				colorId := colorId

				mapPartEntry(neededPartEntry, providedParts.Find(number, colorId), buildMapping)
				mapPartEntry(neededPartEntry, providedParts.FindByEquality(number, colorId, isMold), buildMapping)
				mapPartEntry(neededPartEntry, providedParts.FindByEquality(number, colorId, isPrint), buildMapping)
				mapPartEntry(neededPartEntry, providedParts.FindByEquality(number, colorId, isAlternative), buildMapping)
			}
		}
	}

	// if there are neededParts with quantity > 0, then they are missing in providedParts
	for i := range neededParts.Parts {
		neededPartEntry := &neededParts.Parts[i]
		if neededPartEntry.Quantity > 0 {
			neededPartEntry.Quantity = -1 * neededPartEntry.Quantity
		}
	}

	return buildMapping
}

func mapPartEntry(neededPartEntry *model.PartEntry, providedPartEntry *model.PartEntry, buildMapping *BuildMapping) {
	if providedPartEntry != nil {
		mappedPartEntry := model.DeepClone(providedPartEntry, &model.PartEntry{})
		if providedPartEntry.Quantity >= neededPartEntry.Quantity {
			providedPartEntry.Quantity -= neededPartEntry.Quantity
			mappedPartEntry.Quantity = neededPartEntry.Quantity
			neededPartEntry.Quantity = 0
		} else {
			providedPartEntry.Quantity = 0
			neededPartEntry.Quantity -= mappedPartEntry.Quantity
		}
		buildMapping.Parts[*neededPartEntry] = append(buildMapping.Parts[*neededPartEntry], *mappedPartEntry)
	}
}
