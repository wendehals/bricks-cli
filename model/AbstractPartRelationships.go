package model

import (
	"strings"

	"github.com/wendehals/bricks-cli/utils"
)

type AbstractPartRelationships struct {
	Map map[string]utils.StringSet
}

func NewAbstractRelationships() *AbstractPartRelationships {
	return &AbstractPartRelationships{
		Map: map[string]utils.StringSet{},
	}
}

func (a *AbstractPartRelationships) Add(part1, part2 string) {
	if _, found := a.Map[part1]; !found {
		a.Map[part1] = utils.NewStringSet()
	}
	a.Map[part1].Add(part2)
}

func (a *AbstractPartRelationships) IsCompatible(part1, part2 string) bool {
	return a.Map[part1].Contains(part2)
}

func (a *AbstractPartRelationships) RepresentativeFor(part string) string {
	if set, found := a.Map[part]; found {
		representative := set.Representative()
		if strings.Compare(part, representative) < 0 {
			return part
		}
		return representative
	}
	return part
}

func (a *AbstractPartRelationships) TransitiveClosure() {
	for part, relatedValues := range a.Map {
		visited := utils.NewStringSet()
		visited.Add(part)

		for related := range relatedValues.Values {
			a.addTransitiveParts(part, related, visited)
		}
	}
}

func (a *AbstractPartRelationships) addTransitiveParts(part, related string, visited utils.StringSet) {
	if !visited.Contains(related) {
		visited.Add(related)

		if _, found := a.Map[related]; found {
			for transitivelyRelated := range a.Map[related].Values {
				a.addTransitiveParts(part, transitivelyRelated, visited)
				if part != transitivelyRelated {
					a.Map[part].Add(transitivelyRelated)
				}
			}
		} else {
			a.Map[related] = utils.NewStringSet()
		}
		a.Map[related].Add(part)
	}
}
