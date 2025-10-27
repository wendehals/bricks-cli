package utils

import (
	"testing"

	"github.com/wendehals/bricks-cli/test"
)

func Test_StringSet_Add(t *testing.T) {
	set := NewStringSet()

	test.AssertTrue(t, len(set.Values) == 0)

	set.Add("a")

	_, found := set.Values["a"]
	test.AssertTrue(t, found)
}

func Test_StringSet_Contains(t *testing.T) {
	set := NewStringSet()

	test.AssertFalse(t, set.Contains("a"))

	set.Add("a")

	test.AssertTrue(t, set.Contains("a"))
}

func Test_StringSet_RepresentativeFor(t *testing.T) {
	set := NewStringSet()

	set.Add("3709")
	set.Add("3020")
	set.Add("3020a")

	test.AssertSameString(t, "3020", set.Representative())
	test.AssertSameString(t, "", NewStringSet().Representative())
}
