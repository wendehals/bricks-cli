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

	// search for same part type, same color
	for _, neededPartEntry := range neededCollection.Parts {
		mapping := model.PartMapping{}
		mapping.Quantity = neededPartEntry.Quantity
		mapping.Original = *model.DeepClone(&neededPartEntry, &model.Part{})
		mapping.Substitutes = []model.Part{}

		mapPartEntry(&mapping, providedCollection, neededPartEntry.Color.ID, utils.Equals)

		buildCollection.Mapping = append(buildCollection.Mapping, mapping)
	}

	// search for equivalent part type, same color
	if mode&MOLDS != 0 || mode&PRINTS != 0 || mode&ALTERNATES != 0 {
		for i := range buildCollection.Mapping {
			mapping := &buildCollection.Mapping[i]

			if mode&MOLDS != 0 {
				mapPartEntry(mapping, providedCollection, mapping.Original.Color.ID, isMold)
			}
			if mode&PRINTS != 0 {
				mapPartEntry(mapping, providedCollection, mapping.Original.Color.ID, isPrint)
			}
			if mode&ALTERNATES != 0 {
				mapPartEntry(mapping, providedCollection, mapping.Original.Color.ID, isAlternative)
			}
		}
	}

	// search for same or equivalent part type but different color
	if mode&COLOR != 0 {
		for i := range buildCollection.Mapping {
			mapping := &buildCollection.Mapping[i]
			for j := 0; j < len(colors.Colors) && mapping.Quantity > 0; j++ {
				colorId := colors.Colors[j].ID
				if mapping.Original.Color.ID != colorId {
					mapPartEntry(mapping, providedCollection, colorId, utils.Equals)
					if mode&MOLDS != 0 {
						mapPartEntry(mapping, providedCollection, colorId, isMold)
					}
					if mode&PRINTS != 0 {
						mapPartEntry(mapping, providedCollection, colorId, isPrint)
					}
					if mode&ALTERNATES != 0 {
						mapPartEntry(mapping, providedCollection, colorId, isAlternative)
					}
				}
			}
		}
	}

	return &buildCollection
}

func mapPartEntry(mapping *model.PartMapping, providedCollection *model.Collection,
	colorId int, eqFunc func(string, string) bool) {

	if mapping.Quantity > 0 {
		providedPartEntry := providedCollection.FindByEquivalence(
			mapping.Original.Shape.Number, colorId, eqFunc)

		if providedPartEntry != nil && providedPartEntry.Quantity > 0 {
			mappedPartEntry := model.DeepClone(providedPartEntry, &model.Part{})
			if providedPartEntry.Quantity >= mapping.Quantity {
				providedPartEntry.Quantity -= mapping.Quantity
				mappedPartEntry.Quantity = mapping.Quantity
			} else {
				providedPartEntry.Quantity = 0
			}

			mapping.Quantity -= mappedPartEntry.Quantity
			mapping.Substitutes = append(mapping.Substitutes, *mappedPartEntry)

			providedCollection.Remove(providedPartEntry.Shape.Number,
				providedPartEntry.Color.ID, providedPartEntry.Quantity)
		}
	}
}
