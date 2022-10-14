package services

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/wendehals/bricks/model"
	"github.com/wendehals/bricks/test"
)

func Test_Build_CAMP(t *testing.T) {
	neededCollection := model.Load(model.NewCollection(), "test_resources/886-1.parts")
	providedCollection := model.Load(model.NewCollection(), "test_resources/886-1_provided.parts")

	buildCollection := Build(neededCollection, providedCollection, 15) // mode = camp

	test.AssertTrue(t, cmp.Equal(buildCollection.Set, neededCollection.Sets[0]))

	test.AssertSameInt(t, 12, buildCollection.CountProvidedParts())
	test.AssertSameInt(t, 3, buildCollection.CountMissingParts())

	test.AssertSameInt(t, 2, buildCollection.Mapping[0].Quantity) // two parts missing
	test.AssertSameInt(t, 0, buildCollection.Mapping[1].Quantity)
	test.AssertSameInt(t, 0, buildCollection.Mapping[2].Quantity)
	test.AssertSameInt(t, 0, buildCollection.Mapping[3].Quantity)
	test.AssertSameInt(t, 0, buildCollection.Mapping[4].Quantity)
	test.AssertSameInt(t, 0, buildCollection.Mapping[5].Quantity)
	test.AssertSameInt(t, 0, buildCollection.Mapping[6].Quantity)
	test.AssertSameInt(t, 0, buildCollection.Mapping[7].Quantity)
	test.AssertSameInt(t, 0, buildCollection.Mapping[8].Quantity)
	test.AssertSameInt(t, 0, buildCollection.Mapping[9].Quantity)
	test.AssertSameInt(t, 1, buildCollection.Mapping[10].Quantity) // one part missing

	test.AssertSameString(t, "3957b", buildCollection.Mapping[1].Substitutes[0].Shape.Number) // Mold used

	test.AssertSameInt(t, 4, buildCollection.Mapping[2].Substitutes[0].Color.ID) // different Color used

	test.AssertSameString(t, "3709", buildCollection.Mapping[3].Substitutes[0].Shape.Number) // Alternate in different Color used
	test.AssertSameInt(t, 0, buildCollection.Mapping[3].Substitutes[0].Color.ID)             // different Color used

	test.AssertSameInt(t, 2, len(buildCollection.Mapping[6].Substitutes))        // two substitutes
	test.AssertSameInt(t, 0, buildCollection.Mapping[6].Substitutes[1].Color.ID) // different Color used for second substitute

	test.AssertSameInt(t, 1, buildCollection.Mapping[7].Substitutes[0].Color.ID)                   // different Color used
	test.AssertSameString(t, "3039pr0026", buildCollection.Mapping[7].Substitutes[0].Shape.Number) // Print used
}

func Test_Build_C(t *testing.T) {
	neededCollection := model.Load(model.NewCollection(), "test_resources/886-1.parts")
	providedCollection := model.Load(model.NewCollection(), "test_resources/886-1_provided.parts")

	buildCollection := Build(neededCollection, providedCollection, MODE_COLOR)

	test.AssertTrue(t, cmp.Equal(buildCollection.Set, neededCollection.Sets[0]))

	test.AssertSameInt(t, 9, buildCollection.CountProvidedParts())
	test.AssertSameInt(t, 6, buildCollection.CountMissingParts())

	test.AssertSameInt(t, 2, buildCollection.Mapping[0].Quantity) // two parts missing
	test.AssertSameInt(t, 1, buildCollection.Mapping[1].Quantity) // one part missing
	test.AssertSameInt(t, 0, buildCollection.Mapping[2].Quantity)
	test.AssertSameInt(t, 1, buildCollection.Mapping[3].Quantity) // one part missing
	test.AssertSameInt(t, 0, buildCollection.Mapping[4].Quantity)
	test.AssertSameInt(t, 0, buildCollection.Mapping[5].Quantity)
	test.AssertSameInt(t, 0, buildCollection.Mapping[6].Quantity)
	test.AssertSameInt(t, 1, buildCollection.Mapping[7].Quantity) // one part missing
	test.AssertSameInt(t, 0, buildCollection.Mapping[8].Quantity)
	test.AssertSameInt(t, 0, buildCollection.Mapping[9].Quantity)
	test.AssertSameInt(t, 1, buildCollection.Mapping[10].Quantity) // one part missing

	test.AssertSameInt(t, 4, buildCollection.Mapping[2].Substitutes[0].Color.ID) // different Color used

	test.AssertSameInt(t, 2, len(buildCollection.Mapping[6].Substitutes))        // two substitutes
	test.AssertSameInt(t, 0, buildCollection.Mapping[6].Substitutes[1].Color.ID) // different Color used for second substitute
}

func Test_Build_CA(t *testing.T) {
	neededCollection := model.Load(model.NewCollection(), "test_resources/886-1.parts")
	providedCollection := model.Load(model.NewCollection(), "test_resources/886-1_provided.parts")

	buildCollection := Build(neededCollection, providedCollection, 3) // mode = ca

	test.AssertTrue(t, cmp.Equal(buildCollection.Set, neededCollection.Sets[0]))

	test.AssertSameInt(t, 10, buildCollection.CountProvidedParts())
	test.AssertSameInt(t, 5, buildCollection.CountMissingParts())

	test.AssertSameInt(t, 2, buildCollection.Mapping[0].Quantity) // two parts missing
	test.AssertSameInt(t, 1, buildCollection.Mapping[1].Quantity) // one part missing
	test.AssertSameInt(t, 0, buildCollection.Mapping[2].Quantity)
	test.AssertSameInt(t, 0, buildCollection.Mapping[3].Quantity)
	test.AssertSameInt(t, 0, buildCollection.Mapping[4].Quantity)
	test.AssertSameInt(t, 0, buildCollection.Mapping[5].Quantity)
	test.AssertSameInt(t, 0, buildCollection.Mapping[6].Quantity)
	test.AssertSameInt(t, 1, buildCollection.Mapping[7].Quantity) // one part missing
	test.AssertSameInt(t, 0, buildCollection.Mapping[8].Quantity)
	test.AssertSameInt(t, 0, buildCollection.Mapping[9].Quantity)
	test.AssertSameInt(t, 1, buildCollection.Mapping[10].Quantity) // one part missing

	test.AssertSameInt(t, 4, buildCollection.Mapping[2].Substitutes[0].Color.ID) // different Color used

	test.AssertSameString(t, "3709", buildCollection.Mapping[3].Substitutes[0].Shape.Number) // Alternate in different Color used
	test.AssertSameInt(t, 0, buildCollection.Mapping[3].Substitutes[0].Color.ID)             // different Color used

	test.AssertSameInt(t, 2, len(buildCollection.Mapping[6].Substitutes))        // two substitutes
	test.AssertSameInt(t, 0, buildCollection.Mapping[6].Substitutes[1].Color.ID) // different Color used for second substitute
}

func Test_Build_CAM(t *testing.T) {
	neededCollection := model.Load(model.NewCollection(), "test_resources/886-1.parts")
	providedCollection := model.Load(model.NewCollection(), "test_resources/886-1_provided.parts")

	buildCollection := Build(neededCollection, providedCollection, 7) // mode = camp

	test.AssertTrue(t, cmp.Equal(buildCollection.Set, neededCollection.Sets[0]))

	test.AssertSameInt(t, 11, buildCollection.CountProvidedParts())
	test.AssertSameInt(t, 4, buildCollection.CountMissingParts())

	test.AssertSameInt(t, 2, buildCollection.Mapping[0].Quantity) // two parts missing
	test.AssertSameInt(t, 0, buildCollection.Mapping[1].Quantity)
	test.AssertSameInt(t, 0, buildCollection.Mapping[2].Quantity)
	test.AssertSameInt(t, 0, buildCollection.Mapping[3].Quantity)
	test.AssertSameInt(t, 0, buildCollection.Mapping[4].Quantity)
	test.AssertSameInt(t, 0, buildCollection.Mapping[5].Quantity)
	test.AssertSameInt(t, 0, buildCollection.Mapping[6].Quantity)
	test.AssertSameInt(t, 1, buildCollection.Mapping[7].Quantity) // one part missing
	test.AssertSameInt(t, 0, buildCollection.Mapping[8].Quantity)
	test.AssertSameInt(t, 0, buildCollection.Mapping[9].Quantity)
	test.AssertSameInt(t, 1, buildCollection.Mapping[10].Quantity) // one part missing

	test.AssertSameString(t, "3957b", buildCollection.Mapping[1].Substitutes[0].Shape.Number) // Mold used

	test.AssertSameInt(t, 4, buildCollection.Mapping[2].Substitutes[0].Color.ID) // different Color used

	test.AssertSameString(t, "3709", buildCollection.Mapping[3].Substitutes[0].Shape.Number) // Alternate in different Color used
	test.AssertSameInt(t, 0, buildCollection.Mapping[3].Substitutes[0].Color.ID)             // different Color used

	test.AssertSameInt(t, 2, len(buildCollection.Mapping[6].Substitutes))        // two substitutes
	test.AssertSameInt(t, 0, buildCollection.Mapping[6].Substitutes[1].Color.ID) // different Color used for second substitute
}
