package model

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/wendehals/bricks/test"
)

func Test_Collection_Save(t *testing.T) {
	collection := NewCollection()
	collection.User = "user"
	collection.Comment = "comment"

	set := Set{Name: "setName", Number: "123-1", Year: 1992, NumParts: 5}
	collection.Sets = append(collection.Sets, set)

	part := NewPart()
	part.Quantity = 1
	part.Shape.Name = "Brick 2 x 4"
	part.Shape.Number = "3001"
	collection.Parts = append(collection.Parts, *part)

	tmpFile := filepath.Join(os.TempDir(), "collection.parts")
	defer os.Remove(tmpFile)
	collection.Save(tmpFile)

	actualCollection := Load[Collection](tmpFile)
	test.AssertTrue(t, cmp.Equal(collection, &actualCollection))
}

func Test_Collection_SortByColorAndName(t *testing.T) {
	collection := Load[Collection]("test_resources/testCollection1.parts")

	collection.SortByColorAndName(false)

	test.AssertSameString(t, "42005-1", collection.Sets[0].Number)
	test.AssertSameString(t, "42093-1", collection.Sets[1].Number)

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

func Test_Collection_SortByColorAndName_descending(t *testing.T) {
	collection := Load[Collection]("test_resources/testCollection1.parts")

	collection.SortByColorAndName(true)

	test.AssertSameString(t, "42005-1", collection.Sets[0].Number)
	test.AssertSameString(t, "42093-1", collection.Sets[1].Number)

	test.AssertSameString(t, "yellow", collection.Parts[0].Color.Name)
	test.AssertSameString(t, "name2c", collection.Parts[0].Shape.Name)

	test.AssertSameString(t, "red", collection.Parts[1].Color.Name)
	test.AssertSameString(t, "name25", collection.Parts[1].Shape.Name)

	test.AssertSameString(t, "gray", collection.Parts[2].Color.Name)
	test.AssertSameString(t, "name2a", collection.Parts[2].Shape.Name)

	test.AssertSameString(t, "blue", collection.Parts[3].Color.Name)
	test.AssertSameString(t, "name2", collection.Parts[3].Shape.Name)

	test.AssertSameString(t, "black", collection.Parts[4].Color.Name)
	test.AssertSameString(t, "name25", collection.Parts[4].Shape.Name)

	test.AssertSameString(t, "black", collection.Parts[5].Color.Name)
	test.AssertSameString(t, "name2412a", collection.Parts[5].Shape.Name)
}

func Test_Collection_SortByQuantityAndName(t *testing.T) {
	collection := Load[Collection]("test_resources/testCollection1.parts")

	collection.SortByQuantityAndName(false)

	test.AssertSameString(t, "42005-1", collection.Sets[0].Number)
	test.AssertSameString(t, "42093-1", collection.Sets[1].Number)

	test.AssertSameString(t, "2", collection.Parts[0].Shape.Number)
	test.AssertSameInt(t, 1, collection.Parts[0].Quantity)

	test.AssertSameString(t, "2a", collection.Parts[1].Shape.Number)
	test.AssertSameInt(t, 1, collection.Parts[1].Quantity)

	test.AssertSameString(t, "25", collection.Parts[2].Shape.Number)
	test.AssertSameInt(t, 3, collection.Parts[2].Quantity)

	test.AssertSameString(t, "2412a", collection.Parts[3].Shape.Number)
	test.AssertSameInt(t, 5, collection.Parts[3].Quantity)

	test.AssertSameString(t, "2c", collection.Parts[4].Shape.Number)
	test.AssertSameInt(t, 5, collection.Parts[4].Quantity)

	test.AssertSameString(t, "25", collection.Parts[5].Shape.Number)
	test.AssertSameInt(t, 10, collection.Parts[5].Quantity)
}

func Test_Collection_SortByQuantityAndName_descending(t *testing.T) {
	collection := Load[Collection]("test_resources/testCollection1.parts")

	collection.SortByQuantityAndName(true)

	test.AssertSameString(t, "42005-1", collection.Sets[0].Number)
	test.AssertSameString(t, "42093-1", collection.Sets[1].Number)

	test.AssertSameString(t, "25", collection.Parts[0].Shape.Number)
	test.AssertSameInt(t, 10, collection.Parts[0].Quantity)

	test.AssertSameString(t, "2c", collection.Parts[1].Shape.Number)
	test.AssertSameInt(t, 5, collection.Parts[1].Quantity)

	test.AssertSameString(t, "2412a", collection.Parts[2].Shape.Number)
	test.AssertSameInt(t, 5, collection.Parts[2].Quantity)

	test.AssertSameString(t, "25", collection.Parts[3].Shape.Number)
	test.AssertSameInt(t, 3, collection.Parts[3].Quantity)

	test.AssertSameString(t, "2a", collection.Parts[4].Shape.Number)
	test.AssertSameInt(t, 1, collection.Parts[4].Quantity)

	test.AssertSameString(t, "2", collection.Parts[5].Shape.Number)
	test.AssertSameInt(t, 1, collection.Parts[5].Quantity)
}

func Test_Collection_FindByEquivalence(t *testing.T) {
	collection := Load[Collection]("test_resources/testCollection1.parts")

	equivalent := func(a, b string) bool {
		return a == "2c"
	}

	foundPart := collection.FindByEquivalence("2", 3, equivalent)

	test.AssertTrue(t, foundPart != nil)
	test.AssertSameString(t, "2c", foundPart.Shape.Number)
}

func Test_Collection_FindByEquivalence_none(t *testing.T) {
	collection := Load[Collection]("test_resources/testCollection1.parts")

	notEquivalent := func(a, b string) bool {
		return false
	}

	foundPart := collection.FindByEquivalence("2", 3, notEquivalent)

	test.AssertTrue(t, foundPart == nil)
}

func Test_Collection_Add(t *testing.T) {
	collection := Load[Collection]("test_resources/testCollection1.parts")
	collection2 := Load[Collection]("test_resources/testCollection2.parts")
	collection.Add(&collection2).SortByColorAndName(false)

	assertSize(t, &collection, 9)

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

func Test_Collection_Add_SpareParts(t *testing.T) {
	collection := Load[Collection]("test_resources/addSpareParts1.parts")
	collection2 := Load[Collection]("test_resources/addSpareParts2.parts")
	collection.Add(&collection2).SortByColorAndName(false)

	assertSize(t, &collection, 2)

	test.AssertSameString(t, "42", collection.Parts[0].Shape.Number)
	test.AssertSameInt(t, 6, collection.Parts[0].Quantity)

	test.AssertSameString(t, "99", collection.Parts[1].Shape.Number)
	test.AssertSameInt(t, 7, collection.Parts[1].Quantity)
}

func Test_Collection_Substract(t *testing.T) {
	collection := Load[Collection]("test_resources/testCollection1.parts")
	collection2 := Load[Collection]("test_resources/testCollection2.parts")
	collection.Subtract(&collection2).SortByColorAndName(false)

	assertSize(t, &collection, 9)

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

func Test_Collection_Max(t *testing.T) {
	collection := Load[Collection]("test_resources/testCollection1.parts")
	collection2 := Load[Collection]("test_resources/testCollection2.parts")
	collection.Max(&collection2).SortByColorAndName(false)

	assertSize(t, &collection, 9)

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

func Test_Collection_CountParts(t *testing.T) {
	collection := Load[Collection]("test_resources/testCollection1.parts")

	test.AssertSameInt(t, 25, collection.CountParts())
}

func Test_Collection_HasSpareParts(t *testing.T) {
	collection1 := Load[Collection]("test_resources/testCollection1.parts")
	collection2 := Load[Collection]("test_resources/testCollection2.parts")

	test.AssertTrue(t, collection1.HasSpareParts())
	test.AssertFalse(t, collection2.HasSpareParts())
}

func Test_Collection_HasMissingParts(t *testing.T) {
	collection := Load[Collection]("test_resources/testCollection1.parts")

	test.AssertFalse(t, collection.HasMissingParts())

	collection = *NewCollection()
	part := NewPart()
	part.Quantity = 10
	collection.Parts = append(collection.Parts, *part)
	part = NewPart()
	part.Quantity = -5
	collection.Parts = append(collection.Parts, *part)

	test.AssertTrue(t, collection.HasMissingParts())
}

func assertSize(t *testing.T, collection *Collection, expected int) {
	if len(collection.Parts) != expected {
		t.Errorf("Too many/less distinct part types in collection. Expected: %d, actual: %d", expected, len(collection.Parts))
	}
}
