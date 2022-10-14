package utils

import (
	"testing"

	"github.com/wendehals/bricks/test"
)

func Test_Identity(t *testing.T) {
	test.AssertSameString(t, "x", Identity("x"))
}

func Test_Add(t *testing.T) {
	test.AssertSameInt(t, 2, Add(1, 1))
}

func Test_Subtract(t *testing.T) {
	test.AssertSameInt(t, 1, Subtract(2, 1))
}

func Test_Max(t *testing.T) {
	test.AssertSameInt(t, 5, Max(1, 5))
	test.AssertSameInt(t, 10, Max(10, 5))
}

func Test_Equals(t *testing.T) {
	test.AssertTrue(t, Equals("a", "a"))
	test.AssertFalse(t, Equals("a", "b"))
}
