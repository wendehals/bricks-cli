package model

import (
	"testing"

	"github.com/wendehals/bricks/test"
)

func TestTransitiveClosure(t *testing.T) {
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
