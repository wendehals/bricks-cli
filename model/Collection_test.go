package model

import (
	"testing"
)

func TestSort(t *testing.T) {
	collection := readCollection(t, "test_resources/testCollection1.json")

	collection.Sort()

	assertSameString(t, "2", collection.Parts[0].Part.Number)
	assertSameString(t, "2a", collection.Parts[1].Part.Number)
	assertSameString(t, "2c", collection.Parts[2].Part.Number)
	assertSameString(t, "25", collection.Parts[3].Part.Number)
	assertSameString(t, "25", collection.Parts[4].Part.Number)
}

func TestAdd(t *testing.T) {
	collection := readCollection(t, "test_resources/testCollection1.json").Add(readCollection(t, "test_resources/testCollection2.json")).Sort()

	assertSize(t, collection, 9)

	assertSameString(t, "2", collection.Parts[0].Part.Number)
	assertSameInt(t, 101, collection.Parts[0].Quantity)

	assertSameString(t, "2a", collection.Parts[1].Part.Number)
	assertSameInt(t, 1, collection.Parts[1].Quantity)

	assertSameString(t, "2c", collection.Parts[2].Part.Number)
	assertSameInt(t, 5, collection.Parts[2].Quantity)

	assertSameString(t, "25", collection.Parts[3].Part.Number)
	assertSameInt(t, 10, collection.Parts[3].Quantity)

	assertSameString(t, "25", collection.Parts[4].Part.Number)
	assertSameInt(t, 3, collection.Parts[4].Quantity)

	assertSameString(t, "25", collection.Parts[5].Part.Number)
	assertSameInt(t, 20, collection.Parts[5].Quantity)

	assertSameString(t, "2412a", collection.Parts[6].Part.Number)
	assertSameInt(t, 1, collection.Parts[6].Quantity)
}

func TestSubstract(t *testing.T) {
	collection := readCollection(t, "test_resources/testCollection1.json").Subtract(readCollection(t, "test_resources/testCollection2.json")).Sort()

	assertSize(t, collection, 9)

	assertSameString(t, "2", collection.Parts[0].Part.Number)
	assertSameInt(t, -99, collection.Parts[0].Quantity)

	assertSameString(t, "2a", collection.Parts[1].Part.Number)
	assertSameInt(t, 1, collection.Parts[1].Quantity)

	assertSameString(t, "2c", collection.Parts[2].Part.Number)
	assertSameInt(t, 5, collection.Parts[2].Quantity)

	assertSameString(t, "25", collection.Parts[3].Part.Number)
	assertSameInt(t, 10, collection.Parts[3].Quantity)

	assertSameString(t, "25", collection.Parts[4].Part.Number)
	assertSameInt(t, 3, collection.Parts[4].Quantity)

	assertSameString(t, "25", collection.Parts[5].Part.Number)
	assertSameInt(t, -20, collection.Parts[5].Quantity)

	assertSameString(t, "2412a", collection.Parts[6].Part.Number)
	assertSameInt(t, -1, collection.Parts[6].Quantity)

	assertSameString(t, "2412a", collection.Parts[7].Part.Number)
	assertSameInt(t, -1, collection.Parts[7].Quantity)

	assertSameString(t, "2412b", collection.Parts[8].Part.Number)
	assertSameInt(t, -1, collection.Parts[8].Quantity)
}

func TestMax(t *testing.T) {
	collection := readCollection(t, "test_resources/testCollection1.json").MergeByColor().Max(readCollection(t, "test_resources/testCollection2.json").MergeByColor()).Sort()

	assertSize(t, collection, 6)

	assertSameString(t, "2", collection.Parts[0].Part.Number)
	assertSameInt(t, 100, collection.Parts[0].Quantity)

	assertSameString(t, "2a", collection.Parts[1].Part.Number)
	assertSameInt(t, 1, collection.Parts[1].Quantity)

	assertSameString(t, "2c", collection.Parts[2].Part.Number)
	assertSameInt(t, 5, collection.Parts[2].Quantity)

	assertSameString(t, "25", collection.Parts[3].Part.Number)
	assertSameInt(t, 20, collection.Parts[3].Quantity)

	assertSameString(t, "2412a", collection.Parts[4].Part.Number)
	assertSameInt(t, 2, collection.Parts[4].Quantity)

	assertSameString(t, "2412b", collection.Parts[5].Part.Number)
	assertSameInt(t, 1, collection.Parts[5].Quantity)
}

func TestMergeByColor(t *testing.T) {
	collection := readCollection(t, "test_resources/testCollection1.json").MergeByColor().Sort()

	assertSize(t, collection, 4)

	assertSameString(t, "2", collection.Parts[0].Part.Number)
	assertSameInt(t, 1, collection.Parts[0].Quantity)

	part1 := collection.Parts[1]
	assertSameString(t, "2a", part1.Part.Number)
	assertSameInt(t, 1, part1.Quantity)
	assertSameInt(t, 0, int(part1.Color.ID))
	assertSameString(t, "", part1.Color.Name)
	assertFalse(t, part1.IsSpare)

	assertSameString(t, "2c", collection.Parts[2].Part.Number)
	assertSameInt(t, 5, collection.Parts[2].Quantity)

	assertSameString(t, "25", collection.Parts[3].Part.Number)
	assertSameInt(t, 13, collection.Parts[3].Quantity)
}

func TestMergeByVariant(t *testing.T) {
	collection := readCollection(t, "test_resources/testCollection2.json").MergeByVariant().Sort()

	assertSize(t, collection, 4)

	assertSameString(t, "2", collection.Parts[0].Part.Number)
	assertSameInt(t, 100, collection.Parts[0].Quantity)

	assertSameString(t, "25", collection.Parts[1].Part.Number)
	assertSameInt(t, 20, collection.Parts[1].Quantity)

	assertSameString(t, "2412b", collection.Parts[2].Part.Number)
	assertSameInt(t, 1, collection.Parts[2].Quantity)

	assertSameString(t, "2412b", collection.Parts[3].Part.Number)
	assertSameInt(t, 2, collection.Parts[3].Quantity)
}

func assertSize(t *testing.T, collection *Collection, expected int) {
	if len(collection.Parts) != expected {
		t.Errorf("Too many/less distinct part types in collection. Expected: %d, actual: %d", expected, len(collection.Parts))
	}
}

func assertFalse(t *testing.T, b bool) {
	if b {
		t.Error("Expected false but was true")
	}
}

func assertSameInt(t *testing.T, expected int, actual int) {
	if actual != expected {
		t.Errorf("Actual value '%d' does not equal expected value '%d", actual, expected)
	}
}

func assertSameString(t *testing.T, expected string, actual string) {
	if actual != expected {
		t.Errorf("Actual value '%s' does not equal expected value '%s", actual, expected)
	}
}

func readCollection(t *testing.T, fileName string) *Collection {
	collection, err := ImportCollection(fileName)
	if err != nil {
		t.FailNow()
	}

	return collection
}
