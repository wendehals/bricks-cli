package model

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/wendehals/bricks-cli/test"
)

func Test_BuildCollection_Save(t *testing.T) {
	expectedBuildCollection := createBuildCollection()

	tmpFile := filepath.Join(os.TempDir(), "buildCollection.build")
	defer os.Remove(tmpFile)
	expectedBuildCollection.Save(tmpFile)

	actualBuildCollection := Load[BuildCollection](tmpFile)
	test.AssertTrue(t, cmp.Equal(expectedBuildCollection, &actualBuildCollection))
}

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

	buildCollection = *createBuildCollection()
	test.AssertFalse(t, buildCollection.HasMissingParts())
}

func Test_BuildCollection_CountMissingParts(t *testing.T) {
	buildCollection := Load[BuildCollection]("test_resources/result.build")

	test.AssertSameInt(t, 9, buildCollection.CountMissingParts())
}

func Test_BuildCollection_CountProvidedParts(t *testing.T) {
	buildCollection := Load[BuildCollection]("test_resources/result.build")

	test.AssertSameInt(t, 8, buildCollection.CountProvidedParts())
}

func createBuildCollection() *BuildCollection {
	expectedBuildCollection := NewBuildCollection()
	expectedBuildCollection.Set.Name = "MyBuild"
	expectedBuildCollection.Set.Number = "123-1"
	expectedBuildCollection.Set.NumParts = 1

	mapping := NewPartMapping()

	original := NewPart()
	original.Shape.Number = "3001"
	original.Shape.Name = "Brick 2 x 4"

	substitute := NewPart()
	substitute.Shape.Number = "3001pr0007"
	substitute.Shape.Name = "Brick 2 x 4 with White Diagonal Stripes Print on Both Sides"

	mapping.Quantity = 0
	mapping.Original = *original
	mapping.Substitutes = append(mapping.Substitutes, *substitute)

	expectedBuildCollection.Mapping = append(expectedBuildCollection.Mapping, *mapping)

	return expectedBuildCollection
}
