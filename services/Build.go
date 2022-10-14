package services

import (
	"github.com/wendehals/bricks/model"
	"github.com/wendehals/bricks/utils"
)

const (
	MODE_COLOR      uint8 = 1
	MODE_ALTERNATES uint8 = 2
	MODE_MOLDS      uint8 = 4
	MODE_PRINTS     uint8 = 8
)

func Build(neededCollection *model.Collection, providedCollection *model.Collection, mode uint8) *model.BuildCollection {
	buildCollection := model.NewBuildCollection()
	if len(neededCollection.Sets) > 0 {
		buildCollection.Set = neededCollection.Sets[0]
	}

	// filter spare parts
	neededCollection.Filter(func(p model.Part) bool {
		return !p.IsSpare
	})

	mapSameShapeSameColor(neededCollection, providedCollection, buildCollection)
	mapEquivalentShapeSameColor(providedCollection, buildCollection, mode)
	mapEquivalentShapeDifferentColor(providedCollection, buildCollection, mode)

	buildCollection.Sort()
	providedCollection.RemoveQuantityZero()

	return buildCollection
}

func mapSameShapeSameColor(neededCollection *model.Collection, providedCollection *model.Collection, buildCollection *model.BuildCollection) {
	for _, neededPart := range neededCollection.Parts {
		mapping := model.NewPartMapping()
		mapping.Quantity = neededPart.Quantity
		mapping.Original = *model.DeepClone(&neededPart, &model.Part{})

		mapPart(providedCollection, mapping, neededPart.Color.ID, utils.Equals)

		buildCollection.Mapping = append(buildCollection.Mapping, *mapping)
	}
}

func mapEquivalentShapeSameColor(providedCollection *model.Collection, buildCollection *model.BuildCollection, mode uint8) {
	if mode&MODE_MOLDS != 0 || mode&MODE_PRINTS != 0 || mode&MODE_ALTERNATES != 0 {
		for i := range buildCollection.Mapping {
			mapping := &buildCollection.Mapping[i]

			if mode&MODE_MOLDS != 0 {
				mapPart(providedCollection, mapping, mapping.Original.Color.ID, isMold)
			}
			if mode&MODE_PRINTS != 0 {
				mapPart(providedCollection, mapping, mapping.Original.Color.ID, isPrint)
			}
			if mode&MODE_ALTERNATES != 0 {
				mapPart(providedCollection, mapping, mapping.Original.Color.ID, isAlternative)
			}
		}
	}
}

func mapEquivalentShapeDifferentColor(providedCollection *model.Collection, buildCollection *model.BuildCollection, mode uint8) {
	if mode&MODE_COLOR != 0 {
		colors := GetColors(false)

		for i := range buildCollection.Mapping {
			mapping := &buildCollection.Mapping[i]
			for j := 0; j < len(colors.Colors) && mapping.Quantity > 0; j++ {
				colorId := colors.Colors[j].ID
				if mapping.Original.Color.ID != colorId {
					mapPart(providedCollection, mapping, colorId, utils.Equals)
					if mode&MODE_MOLDS != 0 {
						mapPart(providedCollection, mapping, colorId, isMold)
					}
					if mode&MODE_PRINTS != 0 {
						mapPart(providedCollection, mapping, colorId, isPrint)
					}
					if mode&MODE_ALTERNATES != 0 {
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
			mappedPart := model.DeepClone(providedPart, &model.Part{})
			if providedPart.Quantity >= mapping.Quantity {
				providedPart.Quantity -= mapping.Quantity
				mappedPart.Quantity = mapping.Quantity
			} else {
				providedPart.Quantity = 0
			}

			mapping.Quantity -= mappedPart.Quantity
			mapping.Substitutes = append(mapping.Substitutes, *mappedPart)
		}
	}
}

func isAlternative(s1, s2 string) bool {
	return GetAlternates(false).IsCompatible(s1, s2)
}

func isMold(s1, s2 string) bool {
	return GetMolds(false).IsCompatible(s1, s2)
}

func isPrint(s1, s2 string) bool {
	return GetPrints(false).IsCompatible(s1, s2)
}
