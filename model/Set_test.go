package model

import (
	"testing"

	"github.com/wendehals/bricks-cli/test"
)

func Test_Set_Compare(t *testing.T) {
	set1 := Set{
		Number: "886-1",
	}
	set2 := Set{
		Number: "889-1",
	}
	set3 := Set{
		Number: "889-10",
	}

	test.AssertSameInt(t, -1, set1.Compare(&set2))
	test.AssertSameInt(t, 1, set2.Compare(&set1))
	test.AssertSameInt(t, 0, set1.Compare(&set1))
	test.AssertSameInt(t, -1, set2.Compare(&set3))
	test.AssertSameInt(t, 1, set3.Compare(&set2))
}
