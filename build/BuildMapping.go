package build

import "github.com/wendehals/bricks/model"

type BuildMapping struct {
	Parts map[model.PartEntry][]model.PartEntry
}

func NewBuildMapping() *BuildMapping {
	buildMapping := BuildMapping{}
	buildMapping.Parts = make(map[model.PartEntry][]model.PartEntry)

	return &buildMapping
}
