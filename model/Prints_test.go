package model

import (
	"testing"

	"github.com/wendehals/bricks/test"
)

func Test_Prints_IsCompatible(t *testing.T) {
	prints := NewPrints()

	prints.Add("3001pr0001", "3001")

	test.AssertTrue(t, prints.IsCompatible("3001pr0001", "3001"))
	test.AssertTrue(t, prints.IsCompatible("3001", "3001pr0001"))

	test.AssertFalse(t, prints.IsCompatible("3002pr0001", "3001"))
}

func Test_Prints_RepresentativeFor(t *testing.T) {
	prints := NewPrints()

	prints.Add("3001pr0001", "3001")

	test.AssertSameString(t, "3001", prints.RepresentativeFor("3001pr0001"))
	test.AssertSameString(t, "3001", prints.RepresentativeFor("3001"))
}
