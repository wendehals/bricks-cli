package model

import (
	"testing"

	"github.com/wendehals/bricks/test"
)

func TestSort(t *testing.T) {
	collection := Load(&Collection{}, "test_resources/testCollection1.parts")

	collection.Sort()

	test.AssertSameString(t, "2", collection.Parts[0].Part.Number)
	test.AssertSameString(t, "2a", collection.Parts[1].Part.Number)
	test.AssertSameString(t, "2c", collection.Parts[2].Part.Number)
	test.AssertSameString(t, "25", collection.Parts[3].Part.Number)
	test.AssertSameString(t, "25", collection.Parts[4].Part.Number)
}

func TestAdd(t *testing.T) {
	collection := Load(&Collection{}, "test_resources/testCollection1.parts").Add(Load(&Collection{}, "test_resources/testCollection2.parts")).Sort()

	assertSize(t, collection, 9)

	test.AssertSameString(t, "2", collection.Parts[0].Part.Number)
	test.AssertSameInt(t, 101, collection.Parts[0].Quantity)

	test.AssertSameString(t, "2a", collection.Parts[1].Part.Number)
	test.AssertSameInt(t, 1, collection.Parts[1].Quantity)

	test.AssertSameString(t, "2c", collection.Parts[2].Part.Number)
	test.AssertSameInt(t, 5, collection.Parts[2].Quantity)

	test.AssertSameString(t, "25", collection.Parts[3].Part.Number)
	test.AssertSameInt(t, 10, collection.Parts[3].Quantity)
	test.AssertSameInt(t, 1, collection.Parts[3].Color.ID)

	test.AssertSameString(t, "25", collection.Parts[4].Part.Number)
	test.AssertSameInt(t, 3, collection.Parts[4].Quantity)
	test.AssertSameInt(t, 2, collection.Parts[4].Color.ID)

	test.AssertSameString(t, "25", collection.Parts[5].Part.Number)
	test.AssertSameInt(t, 20, collection.Parts[5].Quantity)
	test.AssertSameInt(t, 3, collection.Parts[5].Color.ID)

	test.AssertSameString(t, "2412a", collection.Parts[6].Part.Number)
	test.AssertSameInt(t, 6, collection.Parts[6].Quantity)
	test.AssertSameInt(t, 1, collection.Parts[6].Color.ID)

	test.AssertSameString(t, "2412a", collection.Parts[7].Part.Number)
	test.AssertSameInt(t, 1, collection.Parts[7].Quantity)
	test.AssertSameInt(t, 3, collection.Parts[7].Color.ID)

	test.AssertSameString(t, "2412b", collection.Parts[8].Part.Number)
	test.AssertSameInt(t, 1, collection.Parts[8].Quantity)
	test.AssertSameInt(t, 1, collection.Parts[8].Color.ID)
}

func TestSubstract(t *testing.T) {
	collection := Load(&Collection{}, "test_resources/testCollection1.parts").Subtract(Load(&Collection{}, "test_resources/testCollection2.parts")).Sort()

	assertSize(t, collection, 9)

	test.AssertSameString(t, "2", collection.Parts[0].Part.Number)
	test.AssertSameInt(t, -99, collection.Parts[0].Quantity)

	test.AssertSameString(t, "2a", collection.Parts[1].Part.Number)
	test.AssertSameInt(t, 1, collection.Parts[1].Quantity)

	test.AssertSameString(t, "2c", collection.Parts[2].Part.Number)
	test.AssertSameInt(t, 5, collection.Parts[2].Quantity)

	test.AssertSameString(t, "25", collection.Parts[3].Part.Number)
	test.AssertSameInt(t, 10, collection.Parts[3].Quantity)

	test.AssertSameString(t, "25", collection.Parts[4].Part.Number)
	test.AssertSameInt(t, 3, collection.Parts[4].Quantity)

	test.AssertSameString(t, "25", collection.Parts[5].Part.Number)
	test.AssertSameInt(t, -20, collection.Parts[5].Quantity)

	test.AssertSameString(t, "2412a", collection.Parts[6].Part.Number)
	test.AssertSameInt(t, 4, collection.Parts[6].Quantity)

	test.AssertSameString(t, "2412a", collection.Parts[7].Part.Number)
	test.AssertSameInt(t, -1, collection.Parts[7].Quantity)

	test.AssertSameString(t, "2412b", collection.Parts[8].Part.Number)
	test.AssertSameInt(t, -1, collection.Parts[8].Quantity)
}

func TestMax(t *testing.T) {
	collection := Load(&Collection{}, "test_resources/testCollection1.parts").Max(Load(&Collection{}, "test_resources/testCollection2.parts")).Sort()

	assertSize(t, collection, 9)

	test.AssertSameString(t, "2", collection.Parts[0].Part.Number)
	test.AssertSameInt(t, 100, collection.Parts[0].Quantity)
	test.AssertSameInt(t, 5, collection.Parts[0].Color.ID)

	test.AssertSameString(t, "2a", collection.Parts[1].Part.Number)
	test.AssertSameInt(t, 1, collection.Parts[1].Quantity)
	test.AssertSameInt(t, 4, collection.Parts[1].Color.ID)

	test.AssertSameString(t, "2c", collection.Parts[2].Part.Number)
	test.AssertSameInt(t, 5, collection.Parts[2].Quantity)
	test.AssertSameInt(t, 3, collection.Parts[2].Color.ID)

	test.AssertSameString(t, "25", collection.Parts[3].Part.Number)
	test.AssertSameInt(t, 10, collection.Parts[3].Quantity)
	test.AssertSameInt(t, 1, collection.Parts[3].Color.ID)

	test.AssertSameString(t, "25", collection.Parts[4].Part.Number)
	test.AssertSameInt(t, 3, collection.Parts[4].Quantity)
	test.AssertSameInt(t, 2, collection.Parts[4].Color.ID)

	test.AssertSameString(t, "25", collection.Parts[5].Part.Number)
	test.AssertSameInt(t, 20, collection.Parts[5].Quantity)
	test.AssertSameInt(t, 3, collection.Parts[5].Color.ID)

	test.AssertSameString(t, "2412a", collection.Parts[6].Part.Number)
	test.AssertSameInt(t, 5, collection.Parts[6].Quantity)
	test.AssertSameInt(t, 1, collection.Parts[6].Color.ID)

	test.AssertSameString(t, "2412a", collection.Parts[7].Part.Number)
	test.AssertSameInt(t, 1, collection.Parts[7].Quantity)
	test.AssertSameInt(t, 3, collection.Parts[7].Color.ID)

	test.AssertSameString(t, "2412b", collection.Parts[8].Part.Number)
	test.AssertSameInt(t, 1, collection.Parts[8].Quantity)
	test.AssertSameInt(t, 1, collection.Parts[8].Color.ID)
}

func assertSize(t *testing.T, collection *Collection, expected int) {
	if len(collection.Parts) != expected {
		t.Errorf("Too many/less distinct part types in collection. Expected: %d, actual: %d", expected, len(collection.Parts))
	}
}
