package model

import (
	"testing"

	"github.com/wendehals/bricks/test"
)

func TestSort(t *testing.T) {
	collection := Load(&Collection{}, "test_resources/testCollection1.parts")

	collection.SortByColorAndName(false)

	test.AssertSameString(t, "black", collection.Parts[0].Color.Name)
	test.AssertSameString(t, "name2412a", collection.Parts[0].Shape.Name)

	test.AssertSameString(t, "black", collection.Parts[1].Color.Name)
	test.AssertSameString(t, "name25", collection.Parts[1].Shape.Name)

	test.AssertSameString(t, "blue", collection.Parts[2].Color.Name)
	test.AssertSameString(t, "name2", collection.Parts[2].Shape.Name)

	test.AssertSameString(t, "gray", collection.Parts[3].Color.Name)
	test.AssertSameString(t, "name2a", collection.Parts[3].Shape.Name)

	test.AssertSameString(t, "red", collection.Parts[4].Color.Name)
	test.AssertSameString(t, "name25", collection.Parts[4].Shape.Name)

	test.AssertSameString(t, "yellow", collection.Parts[5].Color.Name)
	test.AssertSameString(t, "name2c", collection.Parts[5].Shape.Name)
}

func TestAdd(t *testing.T) {
	collection := Load(&Collection{}, "test_resources/testCollection1.parts").Add(Load(&Collection{}, "test_resources/testCollection2.parts")).SortByColorAndName(false)

	assertSize(t, collection, 9)

	test.AssertSameString(t, "2412a", collection.Parts[0].Shape.Number)
	test.AssertSameInt(t, 6, collection.Parts[0].Quantity)
	test.AssertSameInt(t, 1, collection.Parts[0].Color.ID)

	test.AssertSameString(t, "2412b", collection.Parts[1].Shape.Number)
	test.AssertSameInt(t, 1, collection.Parts[1].Quantity)
	test.AssertSameInt(t, 1, collection.Parts[1].Color.ID)

	test.AssertSameString(t, "25", collection.Parts[2].Shape.Number)
	test.AssertSameInt(t, 10, collection.Parts[2].Quantity)
	test.AssertSameInt(t, 1, collection.Parts[2].Color.ID)

	test.AssertSameString(t, "2", collection.Parts[3].Shape.Number)
	test.AssertSameInt(t, 101, collection.Parts[3].Quantity)

	test.AssertSameString(t, "2a", collection.Parts[4].Shape.Number)
	test.AssertSameInt(t, 1, collection.Parts[4].Quantity)

	test.AssertSameString(t, "25", collection.Parts[5].Shape.Number)
	test.AssertSameInt(t, 3, collection.Parts[5].Quantity)
	test.AssertSameInt(t, 2, collection.Parts[5].Color.ID)

	test.AssertSameString(t, "2412a", collection.Parts[6].Shape.Number)
	test.AssertSameInt(t, 1, collection.Parts[6].Quantity)
	test.AssertSameInt(t, 3, collection.Parts[6].Color.ID)

	test.AssertSameString(t, "25", collection.Parts[7].Shape.Number)
	test.AssertSameInt(t, 20, collection.Parts[7].Quantity)
	test.AssertSameInt(t, 3, collection.Parts[7].Color.ID)

	test.AssertSameString(t, "2c", collection.Parts[8].Shape.Number)
	test.AssertSameInt(t, 5, collection.Parts[8].Quantity)
}

func TestAddSpareParts(t *testing.T) {
	collection := Load(&Collection{}, "test_resources/addSpareParts1.parts").Add(Load(&Collection{}, "test_resources/addSpareParts2.parts")).SortByColorAndName(false)

	assertSize(t, collection, 2)

	test.AssertSameString(t, "42", collection.Parts[0].Shape.Number)
	test.AssertSameInt(t, 6, collection.Parts[0].Quantity)

	test.AssertSameString(t, "99", collection.Parts[1].Shape.Number)
	test.AssertSameInt(t, 7, collection.Parts[1].Quantity)
}

func TestSubstract(t *testing.T) {
	collection := Load(&Collection{}, "test_resources/testCollection1.parts").Subtract(Load(&Collection{}, "test_resources/testCollection2.parts")).SortByColorAndName(false)

	assertSize(t, collection, 9)

	test.AssertSameString(t, "2412a", collection.Parts[0].Shape.Number)
	test.AssertSameInt(t, 4, collection.Parts[0].Quantity)

	test.AssertSameString(t, "2412b", collection.Parts[1].Shape.Number)
	test.AssertSameInt(t, -1, collection.Parts[1].Quantity)

	test.AssertSameString(t, "25", collection.Parts[2].Shape.Number)
	test.AssertSameInt(t, 10, collection.Parts[2].Quantity)

	test.AssertSameString(t, "2", collection.Parts[3].Shape.Number)
	test.AssertSameInt(t, -99, collection.Parts[3].Quantity)

	test.AssertSameString(t, "2a", collection.Parts[4].Shape.Number)
	test.AssertSameInt(t, 1, collection.Parts[4].Quantity)

	test.AssertSameString(t, "25", collection.Parts[5].Shape.Number)
	test.AssertSameInt(t, 3, collection.Parts[5].Quantity)

	test.AssertSameString(t, "2412a", collection.Parts[6].Shape.Number)
	test.AssertSameInt(t, -1, collection.Parts[6].Quantity)

	test.AssertSameString(t, "25", collection.Parts[7].Shape.Number)
	test.AssertSameInt(t, -20, collection.Parts[7].Quantity)

	test.AssertSameString(t, "2c", collection.Parts[8].Shape.Number)
	test.AssertSameInt(t, 5, collection.Parts[8].Quantity)
}

func TestMax(t *testing.T) {
	collection := Load(&Collection{}, "test_resources/testCollection1.parts").Max(Load(&Collection{}, "test_resources/testCollection2.parts")).SortByColorAndName(false)

	assertSize(t, collection, 9)

	test.AssertSameString(t, "2412a", collection.Parts[0].Shape.Number)
	test.AssertSameInt(t, 5, collection.Parts[0].Quantity)
	test.AssertSameInt(t, 1, collection.Parts[0].Color.ID)

	test.AssertSameString(t, "2412b", collection.Parts[1].Shape.Number)
	test.AssertSameInt(t, 1, collection.Parts[1].Quantity)
	test.AssertSameInt(t, 1, collection.Parts[1].Color.ID)

	test.AssertSameString(t, "25", collection.Parts[2].Shape.Number)
	test.AssertSameInt(t, 10, collection.Parts[2].Quantity)
	test.AssertSameInt(t, 1, collection.Parts[2].Color.ID)

	test.AssertSameString(t, "2", collection.Parts[3].Shape.Number)
	test.AssertSameInt(t, 100, collection.Parts[3].Quantity)
	test.AssertSameInt(t, 5, collection.Parts[3].Color.ID)

	test.AssertSameString(t, "2a", collection.Parts[4].Shape.Number)
	test.AssertSameInt(t, 1, collection.Parts[4].Quantity)
	test.AssertSameInt(t, 4, collection.Parts[4].Color.ID)

	test.AssertSameString(t, "25", collection.Parts[5].Shape.Number)
	test.AssertSameInt(t, 3, collection.Parts[5].Quantity)
	test.AssertSameInt(t, 2, collection.Parts[5].Color.ID)

	test.AssertSameString(t, "2412a", collection.Parts[6].Shape.Number)
	test.AssertSameInt(t, 1, collection.Parts[6].Quantity)
	test.AssertSameInt(t, 3, collection.Parts[6].Color.ID)

	test.AssertSameString(t, "25", collection.Parts[7].Shape.Number)
	test.AssertSameInt(t, 20, collection.Parts[7].Quantity)
	test.AssertSameInt(t, 3, collection.Parts[7].Color.ID)

	test.AssertSameString(t, "2c", collection.Parts[8].Shape.Number)
	test.AssertSameInt(t, 5, collection.Parts[8].Quantity)
	test.AssertSameInt(t, 3, collection.Parts[8].Color.ID)
}

func TestMergeByColor(t *testing.T) {
	collection := Load(&Collection{}, "test_resources/testCollection1.parts").MergeByColor().SortByColorAndName(false)

	assertSize(t, collection, 5)

	test.AssertSameString(t, "25", collection.Parts[2].Shape.Number)
	test.AssertSameInt(t, 13, collection.Parts[2].Quantity)
	test.AssertSameInt(t, -1, collection.Parts[2].Color.ID)
}

func assertSize(t *testing.T, collection *Collection, expected int) {
	if len(collection.Parts) != expected {
		t.Errorf("Too many/less distinct part types in collection. Expected: %d, actual: %d", expected, len(collection.Parts))
	}
}
