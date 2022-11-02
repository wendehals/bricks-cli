package services

import (
	"testing"

	"github.com/wendehals/bricks/test"
)

func Test_ColorMode(t *testing.T) {
	test.AssertTrue(t, ColorMode(ModeToUInt8("c")))
	test.AssertTrue(t, ColorMode(ModeToUInt8("C")))
	test.AssertFalse(t, ColorMode(ModeToUInt8("amp")))
}

func Test_PrintsMode(t *testing.T) {
	test.AssertTrue(t, PrintsMode(ModeToUInt8("camp")))
	test.AssertTrue(t, PrintsMode(ModeToUInt8("CAMP")))
	test.AssertFalse(t, PrintsMode(ModeToUInt8("cam")))
}

func Test_MoldsMode(t *testing.T) {
	test.AssertTrue(t, MoldsMode(ModeToUInt8("camp")))
	test.AssertTrue(t, MoldsMode(ModeToUInt8("CAMP")))
	test.AssertFalse(t, MoldsMode(ModeToUInt8("c")))
}

func Test_AlternatesMode(t *testing.T) {
	test.AssertTrue(t, AlternatesMode(ModeToUInt8("camp")))
	test.AssertTrue(t, AlternatesMode(ModeToUInt8("CAMP")))
	test.AssertFalse(t, AlternatesMode(ModeToUInt8("cmp")))
}

func Test_ModeToUInt8(t *testing.T) {
	test.AssertSameInt(t, 1, int(ModeToUInt8("c")))
	test.AssertSameInt(t, 1, int(ModeToUInt8("C")))
	test.AssertSameInt(t, 2, int(ModeToUInt8("a")))
	test.AssertSameInt(t, 2, int(ModeToUInt8("A")))
	test.AssertSameInt(t, 4, int(ModeToUInt8("m")))
	test.AssertSameInt(t, 4, int(ModeToUInt8("M")))
	test.AssertSameInt(t, 8, int(ModeToUInt8("p")))
	test.AssertSameInt(t, 8, int(ModeToUInt8("P")))
	test.AssertSameInt(t, 15, int(ModeToUInt8("camp")))
	test.AssertSameInt(t, 15, int(ModeToUInt8("CAMP")))
}
