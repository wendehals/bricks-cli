package model

import (
	"testing"

	"github.com/wendehals/bricks/test"
)

func TestSort(t *testing.T) {
	collection := ImportCollection("test_resources/testCollection1.json")

	collection.Sort()

	test.AssertSameString(t, "2", collection.Parts[0].Part.Number)
	test.AssertSameString(t, "2a", collection.Parts[1].Part.Number)
	test.AssertSameString(t, "2c", collection.Parts[2].Part.Number)
	test.AssertSameString(t, "25", collection.Parts[3].Part.Number)
	test.AssertSameString(t, "25", collection.Parts[4].Part.Number)
}

func TestAdd(t *testing.T) {
	collection := ImportCollection("test_resources/testCollection1.json").Add(ImportCollection("test_resources/testCollection2.json")).Sort()

	assertSize(t, collection, 9)

	test.AssertSameString(t, "2", collection.Parts[0].Part.Number)
	test.AssertSameInt(t, 101, collection.Parts[0].Quantity)

	test.AssertSameString(t, "2a", collection.Parts[1].Part.Number)
	test.AssertSameInt(t, 1, collection.Parts[1].Quantity)

	test.AssertSameString(t, "2c", collection.Parts[2].Part.Number)
	test.AssertSameInt(t, 5, collection.Parts[2].Quantity)

	test.AssertSameString(t, "25", collection.Parts[3].Part.Number)
	test.AssertSameInt(t, 10, collection.Parts[3].Quantity)

	test.AssertSameString(t, "25", collection.Parts[4].Part.Number)
	test.AssertSameInt(t, 3, collection.Parts[4].Quantity)

	test.AssertSameString(t, "25", collection.Parts[5].Part.Number)
	test.AssertSameInt(t, 20, collection.Parts[5].Quantity)

	test.AssertSameString(t, "2412a", collection.Parts[6].Part.Number)
	test.AssertSameInt(t, 1, collection.Parts[6].Quantity)
}

func TestSubstract(t *testing.T) {
	collection := ImportCollection("test_resources/testCollection1.json").Subtract(ImportCollection("test_resources/testCollection2.json")).Sort()

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
	test.AssertSameInt(t, -1, collection.Parts[6].Quantity)

	test.AssertSameString(t, "2412a", collection.Parts[7].Part.Number)
	test.AssertSameInt(t, -1, collection.Parts[7].Quantity)

	test.AssertSameString(t, "2412b", collection.Parts[8].Part.Number)
	test.AssertSameInt(t, -1, collection.Parts[8].Quantity)
}

func TestMax(t *testing.T) {
	collection := ImportCollection("test_resources/testCollection1.json").MergeByColor().Max(ImportCollection("test_resources/testCollection2.json").MergeByColor()).Sort()

	assertSize(t, collection, 6)

	test.AssertSameString(t, "2", collection.Parts[0].Part.Number)
	test.AssertSameInt(t, 100, collection.Parts[0].Quantity)

	test.AssertSameString(t, "2a", collection.Parts[1].Part.Number)
	test.AssertSameInt(t, 1, collection.Parts[1].Quantity)

	test.AssertSameString(t, "2c", collection.Parts[2].Part.Number)
	test.AssertSameInt(t, 5, collection.Parts[2].Quantity)

	test.AssertSameString(t, "25", collection.Parts[3].Part.Number)
	test.AssertSameInt(t, 20, collection.Parts[3].Quantity)

	test.AssertSameString(t, "2412a", collection.Parts[4].Part.Number)
	test.AssertSameInt(t, 2, collection.Parts[4].Quantity)

	test.AssertSameString(t, "2412b", collection.Parts[5].Part.Number)
	test.AssertSameInt(t, 1, collection.Parts[5].Quantity)
}

func TestMergeByColor(t *testing.T) {
	collection := ImportCollection("test_resources/testCollection1.json").MergeByColor().Sort()

	assertSize(t, collection, 4)

	test.AssertSameString(t, "2", collection.Parts[0].Part.Number)
	test.AssertSameInt(t, 1, collection.Parts[0].Quantity)

	part1 := collection.Parts[1]
	test.AssertSameString(t, "2a", part1.Part.Number)
	test.AssertSameInt(t, 1, part1.Quantity)
	test.AssertSameInt(t, 0, int(part1.Color.ID))
	test.AssertSameString(t, "", part1.Color.Name)
	test.AssertFalse(t, part1.IsSpare)

	test.AssertSameString(t, "2c", collection.Parts[2].Part.Number)
	test.AssertSameInt(t, 5, collection.Parts[2].Quantity)

	test.AssertSameString(t, "25", collection.Parts[3].Part.Number)
	test.AssertSameInt(t, 13, collection.Parts[3].Quantity)
}

func TestMergeByVariant(t *testing.T) {
	collection := ImportCollection("test_resources/testCollection2.json").MergeByVariant().Sort()

	assertSize(t, collection, 4)

	test.AssertSameString(t, "2", collection.Parts[0].Part.Number)
	test.AssertSameInt(t, 100, collection.Parts[0].Quantity)

	test.AssertSameString(t, "25", collection.Parts[1].Part.Number)
	test.AssertSameInt(t, 20, collection.Parts[1].Quantity)

	test.AssertSameString(t, "2412b", collection.Parts[2].Part.Number)
	test.AssertSameInt(t, 1, collection.Parts[2].Quantity)

	test.AssertSameString(t, "2412b", collection.Parts[3].Part.Number)
	test.AssertSameInt(t, 2, collection.Parts[3].Quantity)
}

func assertSize(t *testing.T, collection *Collection, expected int) {
	if len(collection.Parts) != expected {
		t.Errorf("Too many/less distinct part types in collection. Expected: %d, actual: %d", expected, len(collection.Parts))
	}
}
