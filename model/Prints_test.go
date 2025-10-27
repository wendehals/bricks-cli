package model

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/wendehals/bricks-cli/test"
)

func Test_Prints_Save(t *testing.T) {
	prints := NewPrints()

	prints.Add("3001pr0001", "3001a")
	prints.Add("123pr0001", "123")

	tmpFile := filepath.Join(os.TempDir(), "prints.json")
	defer os.Remove(tmpFile)
	prints.Save(tmpFile)

	actualPrints := Load[Prints](tmpFile)
	test.AssertTrue(t, cmp.Equal(prints, &actualPrints))
}

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
