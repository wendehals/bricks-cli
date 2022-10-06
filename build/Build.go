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
	buildCollection := &model.BuildCollection{}
	if len(neededCollection.Sets) > 0 {
		buildCollection.Set = neededCollection.Sets[0]
	}
	buildCollection.Mapping = []model.PartMapping{}

	// filter spare parts
	neededCollection.Filter(func(p model.Part) bool {
		return !p.IsSpare
	})

	mapSameShapeSameColor(neededCollection, providedCollection, buildCollection)
	mapEquivalentShapeSameColor(providedCollection, buildCollection, mode)
	mapEquivalentShapeDifferentColor(providedCollection, buildCollection, mode)

	buildCollection.Sort()

	model.Save(buildCollection, fmt.Sprintf("%s/result.build", outputDir))
	export.ExportBuildToHTML(buildCollection, outputDir, "build")

	providedCollection.RemoveQuantityZero()
	model.Save(providedCollection, fmt.Sprintf("%s/remaining.parts", outputDir))
	export.ExportCollectionToHTML(providedCollection, outputDir, "remaining")
}

func mapSameShapeSameColor(neededCollection *model.Collection, providedCollection *model.Collection, buildCollection *model.BuildCollection) {
	for _, neededPart := range neededCollection.Parts {
		mapping := model.PartMapping{}
		mapping.Quantity = neededPart.Quantity
		mapping.Original = *model.DeepClone(&neededPart, &model.Part{})
		mapping.Substitutes = []model.Part{}

		mapPart(providedCollection, &mapping, neededPart.Color.ID, utils.Equals)

		buildCollection.Mapping = append(buildCollection.Mapping, mapping)
	}
}

func mapEquivalentShapeSameColor(providedCollection *model.Collection, buildCollection *model.BuildCollection, mode uint8) {
	if mode&MOLDS != 0 || mode&PRINTS != 0 || mode&ALTERNATES != 0 {
		for i := range buildCollection.Mapping {
			mapping := &buildCollection.Mapping[i]

			if mode&MOLDS != 0 {
				mapPart(providedCollection, mapping, mapping.Original.Color.ID, isMold)
			}
			if mode&PRINTS != 0 {
				mapPart(providedCollection, mapping, mapping.Original.Color.ID, isPrint)
			}
			if mode&ALTERNATES != 0 {
				mapPart(providedCollection, mapping, mapping.Original.Color.ID, isAlternative)
			}
		}
	}
}

func mapEquivalentShapeDifferentColor(providedCollection *model.Collection, buildCollection *model.BuildCollection, mode uint8) {
	if mode&COLOR != 0 {
		colors := model.GetColors(false)

		for i := range buildCollection.Mapping {
			mapping := &buildCollection.Mapping[i]
			for j := 0; j < len(colors.Colors) && mapping.Quantity > 0; j++ {
				colorId := colors.Colors[j].ID
				if mapping.Original.Color.ID != colorId {
					mapPart(providedCollection, mapping, colorId, utils.Equals)
					if mode&MOLDS != 0 {
						mapPart(providedCollection, mapping, colorId, isMold)
					}
					if mode&PRINTS != 0 {
						mapPart(providedCollection, mapping, colorId, isPrint)
					}
					if mode&ALTERNATES != 0 {
						mapPart(providedCollection, mapping, colorId, isAlternative)
					}
				}
			}
		}
	}
}

func mapPart(providedCollection *model.Collection, mapping *model.PartMapping, colorId int, eqFunc func(string, string) bool) {
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

func isMold(s1, s2 string) bool {
	return model.GetPartRelationships(false).IsMold(s1, s2)
}

func isPrint(s1, s2 string) bool {
	return model.GetPartRelationships(false).IsPrint(s1, s2)
}

func isAlternative(s1, s2 string) bool {
	return model.GetPartRelationships(false).IsAlternative(s1, s2)
}
