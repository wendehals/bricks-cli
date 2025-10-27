package model

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/wendehals/bricks-cli/test"
)

func Test_Alternates_Save(t *testing.T) {
	expectedAlternates := NewAlternates()

	expectedAlternates.Add("3709", "3020")
	expectedAlternates.Add("3020a", "3020")

	tmpFile := filepath.Join(os.TempDir(), "alternates.json")
	expectedAlternates.Save(tmpFile)
	defer os.Remove(tmpFile)

	actualAlternates := Load[Alternates](tmpFile)
	test.AssertTrue(t, cmp.Equal(expectedAlternates, &actualAlternates))
}
