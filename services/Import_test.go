package services

import (
	"testing"

	"github.com/wendehals/bricks-cli/test"
)

func Test_ImportCSVCollection(t *testing.T) {
	collection := ImportCSVCollection("test_resources/rebrickable_parts_test-list.csv")

	test.AssertTrue(t, collection != nil)

	test.AssertSameInt(t, 5, collection.Parts[0].Quantity)
	test.AssertSameString(t, "3001", collection.Parts[0].Shape.Number)
	test.AssertSameString(t, "Brick 2 x 4", collection.Parts[0].Shape.Name)
	test.AssertSameString(t, "https://rebrickable.com/parts/3001", collection.Parts[0].Shape.URL)
	test.AssertSameInt(t, 0, collection.Parts[0].Color.ID)
	test.AssertSameString(t, "Black", collection.Parts[0].Color.Name)

	test.AssertSameInt(t, 2, collection.Parts[1].Quantity)
	test.AssertSameString(t, "3008", collection.Parts[1].Shape.Number)
	test.AssertSameString(t, "Brick 1 x 8", collection.Parts[1].Shape.Name)
	test.AssertSameString(t, "https://rebrickable.com/parts/3008", collection.Parts[1].Shape.URL)
	test.AssertSameInt(t, 1, collection.Parts[1].Color.ID)
	test.AssertSameString(t, "Blue", collection.Parts[1].Color.Name)

	test.AssertSameInt(t, 10, collection.Parts[2].Quantity)
	test.AssertSameString(t, "3003", collection.Parts[2].Shape.Number)
	test.AssertSameString(t, "Brick 2 x 2", collection.Parts[2].Shape.Name)
	test.AssertSameString(t, "https://rebrickable.com/parts/3003", collection.Parts[2].Shape.URL)
	test.AssertSameInt(t, 4, collection.Parts[2].Color.ID)
	test.AssertSameString(t, "Red", collection.Parts[2].Color.Name)
}
