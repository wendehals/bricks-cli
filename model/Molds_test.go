package model

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/wendehals/bricks-cli/test"
)

func Test_Molds_Save(t *testing.T) {
	molds := NewMolds()

	molds.Add("3001", "3001a")
	molds.Add("123", "1234")

	tmpFile := filepath.Join(os.TempDir(), "molds.json")
	defer os.Remove(tmpFile)
	molds.Save(tmpFile)

	actualMolds := Load[Molds](tmpFile)
	test.AssertTrue(t, cmp.Equal(molds, &actualMolds))
}
