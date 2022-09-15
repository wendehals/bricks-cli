package model

type BuildMapping struct {
	Parts map[PartEntry][]PartEntry
}

func NewBuildMapping() *BuildMapping {
	buildMapping := BuildMapping{}
	buildMapping.Parts = make(map[PartEntry][]PartEntry)

	return &buildMapping
}
