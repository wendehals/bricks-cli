package build

import (
	"fmt"

	"github.com/wendehals/bricks/export"
	"github.com/wendehals/bricks/model"
	"github.com/wendehals/bricks/utils"
)

const (
	COLOR      uint8 = 1
	ALTERNATES uint8 = 2
	MOLDS      uint8 = 4
	PRINTS     uint8 = 8
)

func Build(neededCollection *model.Collection, providedCollection *model.Collection, mode uint8, outputDir string, verbose bool) {
	colors := model.GetColors(false)
	partRelationships := model.GetPartRelationships(false)

	buildCollection := build(neededCollection, providedCollection, colors, partRelationships, mode)
	buildCollection.Sort()

	model.Save(buildCollection, fmt.Sprintf("%s/result.build", outputDir))
	export.ExportBuildToHTML(buildCollection, outputDir, "build")

	providedCollection.RemoveQuantityZero()
	model.Save(providedCollection, fmt.Sprintf("%s/remaining.parts", outputDir))
	export.ExportCollectionToHTML(providedCollection, outputDir, "remaining")
}

func build(neededCollection *model.Collection, providedCollection *model.Collection, colors *model.Colors,
	partRelationships *model.PartRelationships, mode uint8) *model.BuildCollection {

	buildCollection := model.BuildCollection{}
	if len(neededCollection.Sets) > 0 {
		buildCollection.Set = neededCollection.Sets[0]
	}
	buildCollection.Mapping = []model.PartMapping{}

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
	neededCollection.Filter(func(p model.Part) bool {
		return !p.IsSpare
	})

	mapSameShapeSameColor(neededCollection, providedCollection, buildCollection)
	mapEquivalentShapeSameColor(mode, buildCollection, providedCollection, isMold, isPrint, isAlternative)
	mapEquivalentShapeDifferentColor(mode, buildCollection, colors, providedCollection, isMold, isPrint, isAlternative)

	return &buildCollection
}

func mapSameShapeSameColor(neededCollection *model.Collection, providedCollection *model.Collection, buildCollection model.BuildCollection) {
	for _, neededPart := range neededCollection.Parts {
		mapping := model.PartMapping{}
		mapping.Quantity = neededPart.Quantity
		mapping.Original = *model.DeepClone(&neededPart, &model.Part{})
		mapping.Substitutes = []model.Part{}

		mapPart(&mapping, providedCollection, neededPart.Color.ID, utils.Equals)

		buildCollection.Mapping = append(buildCollection.Mapping, mapping)
	}
}

func mapEquivalentShapeSameColor(mode uint8, buildCollection model.BuildCollection, providedCollection *model.Collection,
	isMold func(s1 string, s2 string) bool, isPrint func(s1 string, s2 string) bool, isAlternative func(s1 string, s2 string) bool) {

	if mode&MOLDS != 0 || mode&PRINTS != 0 || mode&ALTERNATES != 0 {
		for i := range buildCollection.Mapping {
			mapping := &buildCollection.Mapping[i]

			if mode&MOLDS != 0 {
				mapPart(mapping, providedCollection, mapping.Original.Color.ID, isMold)
			}
			if mode&PRINTS != 0 {
				mapPart(mapping, providedCollection, mapping.Original.Color.ID, isPrint)
			}
			if mode&ALTERNATES != 0 {
				mapPart(mapping, providedCollection, mapping.Original.Color.ID, isAlternative)
			}
		}
	}
}

func mapEquivalentShapeDifferentColor(mode uint8, buildCollection model.BuildCollection, colors *model.Colors, providedCollection *model.Collection,
	isMold func(s1 string, s2 string) bool, isPrint func(s1 string, s2 string) bool, isAlternative func(s1 string, s2 string) bool) {

	if mode&COLOR != 0 {
		for i := range buildCollection.Mapping {
			mapping := &buildCollection.Mapping[i]
			for j := 0; j < len(colors.Colors) && mapping.Quantity > 0; j++ {
				colorId := colors.Colors[j].ID
				if mapping.Original.Color.ID != colorId {
					mapPart(mapping, providedCollection, colorId, utils.Equals)
					if mode&MOLDS != 0 {
						mapPart(mapping, providedCollection, colorId, isMold)
					}
					if mode&PRINTS != 0 {
						mapPart(mapping, providedCollection, colorId, isPrint)
					}
					if mode&ALTERNATES != 0 {
						mapPart(mapping, providedCollection, colorId, isAlternative)
					}
				}
			}
		}
	}
}

func mapPart(mapping *model.PartMapping, providedCollection *model.Collection,
	colorId int, eqFunc func(string, string) bool) {

	if mapping.Quantity > 0 {
		providedPart := providedCollection.FindByEquivalence(
			mapping.Original.Shape.Number, colorId, eqFunc)

		if providedPart != nil && providedPart.Quantity > 0 {
			mappedPartEntry := model.DeepClone(providedPart, &model.Part{})
			if providedPart.Quantity >= mapping.Quantity {
				providedPart.Quantity -= mapping.Quantity
				mappedPartEntry.Quantity = mapping.Quantity
			} else {
				providedPart.Quantity = 0
			}

			mapping.Quantity -= mappedPartEntry.Quantity
			mapping.Substitutes = append(mapping.Substitutes, *mappedPartEntry)

			providedCollection.Remove(providedPart.Shape.Number,
				providedPart.Color.ID, providedPart.Quantity)
		}
	}
}
