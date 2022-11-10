package model

import "sort"

type BuildCollection struct {
	Set     Set
	Mapping []PartMapping `json:"mapping"`
}

func NewBuildCollection() *BuildCollection {
	return &BuildCollection{
		Set:     Set{},
		Mapping: []PartMapping{},
	}
}

// Sort the Parts of a collection by their Part Number
func (b *BuildCollection) Sort() *BuildCollection {
	sort.Slice(b.Mapping, func(i, j int) bool {
		return b.Mapping[i].Compare(&b.Mapping[j]) < 0
	})

	return b
}

func (b *BuildCollection) HasMissingParts() bool {
	for _, entry := range b.Mapping {
		if entry.Quantity > 0 {
			return true
		}
	}

	return false
}

func (b *BuildCollection) CountMissingParts() int {
	var missing int
	for _, entry := range b.Mapping {
		missing += entry.Quantity
	}

	return missing
}

func (b *BuildCollection) CountProvidedParts() int {
	var provided int
	for _, entry := range b.Mapping {
		for _, substitute := range entry.Substitutes {
			provided += substitute.Quantity
		}
	}

	return provided
}

func (b *BuildCollection) Save(filePath string) {
	Save(b, filePath)
}
