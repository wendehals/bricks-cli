package services

import (
	"testing"

	"github.com/wendehals/bricks-cli/test"
)

func Test_GetColors(t *testing.T) {
	localColors := GetColors(true)

	test.AssertTrue(t, localColors != nil)

	test.AssertSameInt(t, -1, localColors.Colors[0].ID)
	test.AssertSameString(t, "[Unknown]", localColors.Colors[0].Name)

	test.AssertSameInt(t, 0, localColors.Colors[1].ID)
	test.AssertSameString(t, "Black", localColors.Colors[1].Name)

	localColors = GetColors(false) // should used cached colors

	test.AssertTrue(t, localColors != nil)

	test.AssertSameInt(t, -1, localColors.Colors[0].ID)
	test.AssertSameString(t, "[Unknown]", localColors.Colors[0].Name)

	test.AssertSameInt(t, 0, localColors.Colors[1].ID)
	test.AssertSameString(t, "Black", localColors.Colors[1].Name)

	colors = nil                   // remove singleton instance
	localColors = GetColors(false) // should reload colors file

	test.AssertTrue(t, localColors != nil)

	test.AssertSameInt(t, -1, localColors.Colors[0].ID)
	test.AssertSameString(t, "[Unknown]", localColors.Colors[0].Name)

	test.AssertSameInt(t, 0, localColors.Colors[1].ID)
	test.AssertSameString(t, "Black", localColors.Colors[1].Name)
}

func Test_GetShapes(t *testing.T) {
	localShapes := GetShapes(true)

	test.AssertTrue(t, localShapes != nil)
	test.AssertTrue(t, len(localShapes.Shapes) > 0)

	localShapes = GetShapes(false) // should used cached shapes

	test.AssertTrue(t, localShapes != nil)
	test.AssertTrue(t, len(localShapes.Shapes) > 0)

	shapes = nil                   // remove singleton instance
	localShapes = GetShapes(false) // should reload shapes file

	test.AssertTrue(t, localShapes != nil)
	test.AssertTrue(t, len(localShapes.Shapes) > 0)
}

func Test_GetAlternates(t *testing.T) {
	localAlternates := GetAlternates(true)

	test.AssertTrue(t, localAlternates != nil)
	test.AssertTrue(t, len(localAlternates.Map) > 0)

	localAlternates = GetAlternates(false) // should used cached alternates

	test.AssertTrue(t, localAlternates != nil)
	test.AssertTrue(t, len(localAlternates.Map) > 0)

	alternates = nil                       // remove singleton instance
	localAlternates = GetAlternates(false) // should reload alternates file

	test.AssertTrue(t, localAlternates != nil)
	test.AssertTrue(t, len(localAlternates.Map) > 0)
}

func Test_GetMolds(t *testing.T) {
	localMolds := GetMolds(true)

	test.AssertTrue(t, localMolds != nil)
	test.AssertTrue(t, len(localMolds.Map) > 0)

	localMolds = GetMolds(false) // should used cached molds

	test.AssertTrue(t, localMolds != nil)
	test.AssertTrue(t, len(localMolds.Map) > 0)

	molds = nil                  // remove singleton instance
	localMolds = GetMolds(false) // should reload molds file

	test.AssertTrue(t, localMolds != nil)
	test.AssertTrue(t, len(localMolds.Map) > 0)
}

func Test_GetPrints(t *testing.T) {
	localPrints := GetPrints(true)

	test.AssertTrue(t, localPrints != nil)
	test.AssertTrue(t, len(localPrints.Map) > 0)

	localPrints = GetPrints(false) // should used cached prints

	test.AssertTrue(t, localPrints != nil)
	test.AssertTrue(t, len(localPrints.Map) > 0)

	prints = nil                   // remove singleton instance
	localPrints = GetPrints(false) // should reload prints file

	test.AssertTrue(t, localPrints != nil)
	test.AssertTrue(t, len(localPrints.Map) > 0)
}
