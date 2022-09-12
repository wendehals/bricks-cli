package model

import (
	"testing"

	"github.com/wendehals/bricks/test"
)

var p = ImportPartRelationships("test_resources/partRelationships.json")

func TestIsAlternativeCompatible(t *testing.T) {
	test.AssertTrue(t, p.IsAlternativeCompatible("3020", "3709"))
	test.AssertTrue(t, p.IsAlternativeCompatible("3709", "3020"))
}

func TestIsMoldCompatible(t *testing.T) {
	test.AssertTrue(t, p.IsMoldCompatible("4085a", "4085b"))
	test.AssertTrue(t, p.IsMoldCompatible("4085b", "4085a"))
}

func TestIsPrintCompatible(t *testing.T) {
	test.AssertTrue(t, p.IsPrintCompatible("3010pr0918", "3010"))
	test.AssertTrue(t, p.IsPrintCompatible("3010", "3010pr0918"))
}
