package model

import (
	"encoding/json"
	"log"
	"sort"
)

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

// HasMissingParts returns true if there are parts missing in the BuildCollection.
func (b *BuildCollection) HasMissingParts() bool {
	for _, entry := range b.Mapping {
		if entry.Quantity > 0 {
			return true
		}
	}

	return false
}

// CountMissingParts counts all parts that are still missing in the BuildCollection.
func (b *BuildCollection) CountMissingParts() int {
	var missing int
	for _, entry := range b.Mapping {
		missing += entry.Quantity
	}

	return missing
}

// CountProvidedParts counts all parts that are provided in the BuildCollection, including substitutes.
func (b *BuildCollection) CountProvidedParts() int {
	var provided int
	for _, entry := range b.Mapping {
		for _, substitute := range entry.Substitutes {
			provided += substitute.Quantity
		}
	}

	return provided
}

// Save writes the BuildCollection into a JSON encoded file. In case of an error it terminates with a fatal log entry.
func (b BuildCollection) Save(filePath string) {
	Save(b, filePath)
}

// Print outputs a human-readable representation of the BuildCollection to the standard output.
func (b BuildCollection) Print() {
	Print(b)
}

// Clone creates a deep copy of the BuildCollection and returns it as a CollectionType.
func (b BuildCollection) Clone() CollectionType {
	origJSON, err := json.Marshal(b)
	if err != nil {
		log.Fatalf(CLONING_FAILED_MSG, err.Error())
	}

	clone := new(BuildCollection)
	err = json.Unmarshal(origJSON, clone)
	if err != nil {
		log.Fatalf(CLONING_FAILED_MSG, err.Error())
	}

	return *clone
}
