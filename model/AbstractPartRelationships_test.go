package model

import (
	"testing"

	"github.com/wendehals/bricks/test"
)

func Test_AbstractRelationships_RepresentativeFor(t *testing.T) {
	alternates := NewAlternates()

	alternates.Add("3709", "3020")
	alternates.Add("3020a", "3020")

	alternates.TransitiveClosure()

	test.AssertSameString(t, "3020", alternates.RepresentativeFor("3709"))
	test.AssertSameString(t, "3020", alternates.RepresentativeFor("3020"))
	test.AssertSameString(t, "3020", alternates.RepresentativeFor("3020a"))
	test.AssertSameString(t, "3001", alternates.RepresentativeFor("3001"))
}

func Test_AbstractPartRelationships_TransitiveClosure(t *testing.T) {
	alternates := NewAlternates()

	alternates.Add("3709", "3020")
	alternates.Add("3020a", "3020")

	alternates.TransitiveClosure()

	test.AssertTrue(t, alternates.IsCompatible("3020", "3709"))
	test.AssertTrue(t, alternates.IsCompatible("3020", "3020a"))

	test.AssertTrue(t, alternates.IsCompatible("3020a", "3020"))
	test.AssertTrue(t, alternates.IsCompatible("3020a", "3709"))

	test.AssertTrue(t, alternates.IsCompatible("3709", "3020"))
	test.AssertTrue(t, alternates.IsCompatible("3709", "3020a"))
}
