package options

import (
	"testing"

	"github.com/wendehals/bricks/test"
)

func TestFileNameFrom(t *testing.T) {
	_, err := FileNameFrom("CON")
	test.AssertError(t, err)

	result, err := FileNameFrom("a<b>c:d\"e/f\\g|h?i*j")
	test.AssertNoError(t, err)
	test.AssertSameString(t, "a_b_c_d_e_f_g_h_i_j", result)
}
