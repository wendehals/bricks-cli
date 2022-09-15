package model

import (
	"testing"

	"github.com/wendehals/bricks/test"
)

var p = Load(&PartRelationships{}, "test_resources/partRelationships.json")

func TestIsAlternativeCompatible(t *testing.T) {
	test.AssertTrue(t, p.IsAlternative("3020", "3709"))
	test.AssertTrue(t, p.IsAlternative("3709", "3020"))
}

func TestIsMoldCompatible(t *testing.T) {
	test.AssertTrue(t, p.IsMold("4085a", "4085b"))
	test.AssertTrue(t, p.IsMold("4085b", "4085a"))
}

func TestIsPrintCompatible(t *testing.T) {
	test.AssertTrue(t, p.IsPrint("3010pr0918", "3010"))
	test.AssertTrue(t, p.IsPrint("3010", "3010pr0918"))
}
