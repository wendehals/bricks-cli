package model

import (
	"testing"

	"github.com/wendehals/bricks/test"
)

func Test_BuildCollection_Sort(t *testing.T) {
	buildCollection := Load[BuildCollection]("test_resources/result.build")

	buildCollection.Sort()

	test.AssertSameString(t, "3005", buildCollection.Mapping[0].Original.Shape.Number)
	test.AssertSameString(t, "3003", buildCollection.Mapping[1].Original.Shape.Number)
	test.AssertSameString(t, "3001", buildCollection.Mapping[2].Original.Shape.Number)
}

func Test_BuildCollection_HasMissingParts(t *testing.T) {
	buildCollection := Load[BuildCollection]("test_resources/result.build")

	test.AssertTrue(t, buildCollection.HasMissingParts())
}

func Test_BuildCollection_CountMissingParts(t *testing.T) {
	buildCollection := Load[BuildCollection]("test_resources/result.build")

	test.AssertSameInt(t, 9, buildCollection.CountMissingParts())
}

func Test_BuildCollection_CountProvidedParts(t *testing.T) {
	buildCollection := Load[BuildCollection]("test_resources/result.build")

	test.AssertSameInt(t, 8, buildCollection.CountProvidedParts())
}
