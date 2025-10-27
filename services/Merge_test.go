package services

import (
	"testing"

	"github.com/wendehals/bricks-cli/model"
	"github.com/wendehals/bricks-cli/test"
)

func Test_MergeAllCollections(t *testing.T) {
	collections := []model.Collection{
		model.Load[model.Collection]("../model/test_resources/testCollection1.parts"),
		model.Load[model.Collection]("../model/test_resources/testCollection2.parts"),
	}

	collection := MergeAllCollections(collections)

	test.AssertSameInt(t, 9, len(collection.Parts))

	test.AssertSameString(t, "42093-1", collection.Sets[0].Number)
	test.AssertSameString(t, "42005-1", collection.Sets[1].Number)

	test.AssertSameInt(t, 148, collection.CountParts())
}

func Test_Merge_C(t *testing.T) {
	collection := model.Load[model.Collection]("../model/test_resources/testCollection1.parts")
	Merge(&collection, ModeToUInt8("c"))
	collection.SortByColorAndName(false)

	assertMergeByColor(t, &collection)
}

func Test_MergeByColor(t *testing.T) {
	collection := model.Load[model.Collection]("../model/test_resources/testCollection1.parts")
	MergeByColor(&collection)
	collection.SortByColorAndName(false)

	assertMergeByColor(t, &collection)
}

func Test_Merge_P(t *testing.T) {
	collection := model.Load[model.Collection]("test_resources/mergeByPrint.parts")
	Merge(&collection, ModeToUInt8("p"))
	collection.SortByColorAndName(false)

	assertMergeByPrint(t, &collection)
}

func Test_MergeByPrint(t *testing.T) {
	collection := model.Load[model.Collection]("test_resources/mergeByPrint.parts")
	MergeByPrint(&collection)
	collection.SortByColorAndName(false)

	assertMergeByPrint(t, &collection)
}

func Test_Merge_M(t *testing.T) {
	collection := model.Load[model.Collection]("test_resources/mergeByMold.parts")
	Merge(&collection, ModeToUInt8("m"))
	collection.SortByColorAndName(false)

	assertMergeByMold(t, &collection)
}

func Test_MergeByMold(t *testing.T) {
	collection := model.Load[model.Collection]("test_resources/mergeByMold.parts")
	MergeByMold(&collection)
	collection.SortByColorAndName(false)

	assertMergeByMold(t, &collection)
}

func Test_Merge_A(t *testing.T) {
	collection := model.Load[model.Collection]("test_resources/mergeByAlternate.parts")
	Merge(&collection, ModeToUInt8("a"))
	collection.SortByColorAndName(false)

	assertMergeByAlternate(t, &collection)
}

func Test_MergeByAlternate(t *testing.T) {
	collection := model.Load[model.Collection]("test_resources/mergeByAlternate.parts")
	MergeByAlternate(&collection)
	collection.SortByColorAndName(false)

	assertMergeByAlternate(t, &collection)
}

func assertMergeByColor(t *testing.T, collection *model.Collection) {
	test.AssertSameInt(t, 5, len(collection.Parts))

	test.AssertSameString(t, "25", collection.Parts[2].Shape.Number)
	test.AssertSameInt(t, 13, collection.Parts[2].Quantity)
	test.AssertSameInt(t, -1, collection.Parts[2].Color.ID)
}

func assertMergeByPrint(t *testing.T, collection *model.Collection) {
	test.AssertSameInt(t, 3, len(collection.Parts))
	test.AssertSameInt(t, 6, collection.CountParts())

	// print is replaced by alternate
	test.AssertSameInt(t, 1, collection.Parts[0].Quantity)
	test.AssertSameString(t, "3001", collection.Parts[0].Shape.Number)
	test.AssertSameString(t, "Brick 2 x 4", collection.Parts[0].Shape.Name)

	// prints are replaced by representant
	test.AssertSameInt(t, 4, collection.Parts[1].Quantity)
	test.AssertSameString(t, "3039", collection.Parts[1].Shape.Number)
	test.AssertSameString(t, "Slope 45° 2 x 2", collection.Parts[1].Shape.Name)

	// alternate keeps the same
	test.AssertSameInt(t, 1, collection.Parts[2].Quantity)
	test.AssertSameString(t, "3709", collection.Parts[2].Shape.Number)
	test.AssertSameString(t, "Technic Plate 2 x 4 [3 Holes]", collection.Parts[2].Shape.Name)
}

func assertMergeByMold(t *testing.T, collection *model.Collection) {
	test.AssertSameInt(t, 3, len(collection.Parts))
	test.AssertSameInt(t, 4, collection.CountParts())

	// molds are replaced by representant
	test.AssertSameInt(t, 2, collection.Parts[0].Quantity)
	test.AssertSameString(t, "3957a", collection.Parts[0].Shape.Number)
	test.AssertSameString(t, "Antenna 1 x 4 with Rounded Top", collection.Parts[0].Shape.Name)

	// mold is replaced by representant
	test.AssertSameInt(t, 1, collection.Parts[1].Quantity)
	test.AssertSameString(t, "3001", collection.Parts[1].Shape.Number)
	test.AssertSameString(t, "Brick 2 x 4", collection.Parts[1].Shape.Name)

	// alternate keeps the same
	test.AssertSameInt(t, 1, collection.Parts[2].Quantity)
	test.AssertSameString(t, "3709", collection.Parts[2].Shape.Number)
	test.AssertSameString(t, "Technic Plate 2 x 4 [3 Holes]", collection.Parts[2].Shape.Name)
}

func assertMergeByAlternate(t *testing.T, collection *model.Collection) {
	test.AssertSameInt(t, 3, len(collection.Parts))
	test.AssertSameInt(t, 3, collection.CountParts())

	// mold keeps the same
	test.AssertSameInt(t, 1, collection.Parts[0].Quantity)
	test.AssertSameString(t, "3957b", collection.Parts[0].Shape.Number)
	test.AssertSameString(t, "Antenna 1 x 4 with Flat Top", collection.Parts[0].Shape.Name)

	// alternate replaced by representant
	test.AssertSameInt(t, 1, collection.Parts[1].Quantity)
	test.AssertSameString(t, "3020", collection.Parts[1].Shape.Number)
	test.AssertSameString(t, "Plate 2 x 4", collection.Parts[1].Shape.Name)

	// print keeps the same
	test.AssertSameInt(t, 1, collection.Parts[2].Quantity)
	test.AssertSameString(t, "3039pr0026", collection.Parts[2].Shape.Number)
	test.AssertSameString(t, "Slope 45° 2 x 2 with Computer Screen Print", collection.Parts[2].Shape.Name)
}
