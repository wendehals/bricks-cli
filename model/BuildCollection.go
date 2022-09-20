package model

type BuildCollection struct {
	ID    string             `json:"id"`
	Name  string             `json:"name"`
	Parts []PartEntryMapping `json:"parts"`
}

type PartEntryMapping struct {
	Quantity    int         `json:"quantity"`
	Original    PartEntry   `json:"original"`
	Substitutes []PartEntry `json:"substitutes"`
}

func (b *BuildCollection) HasMissingParts() bool {
	for i := range b.Parts {
		if b.Parts[i].Quantity > 0 {
			return true
		}
	}

	return false
}

func (b *BuildCollection) CountMissingParts() int {
	var missing int
	for _, part := range b.Parts {
		missing += part.Quantity
	}

	return missing
}

func (b *BuildCollection) CountProvidedParts() int {
	var provided int
	for _, part := range b.Parts {
		for _, substitute := range part.Substitutes {
			provided += substitute.Quantity
		}
	}

	return provided
}
