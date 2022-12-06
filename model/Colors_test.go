package model

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/wendehals/bricks/test"
)

func Test_Colors_Save(t *testing.T) {
	colors := NewColors()

	colors.Colors = append(colors.Colors, Color{ID: 1, Name: "black"})
	colors.Colors = append(colors.Colors, Color{ID: 2, Name: "red"})

	tmpFile := filepath.Join(os.TempDir(), "colors.json")
	defer os.Remove(tmpFile)
	colors.Save(tmpFile)

	actualColors := Load[Colors](tmpFile)
	test.AssertTrue(t, cmp.Equal(colors, &actualColors))
}
