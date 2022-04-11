package options

import (
	"testing"

	"github.com/wendehals/bricks/test"
)

func TestReplaceIllegalCharsFromFileName(t *testing.T) {
	result := ReplaceIllegalCharsFromFileName("a<b>c:d\"e/f\\g|h?i*j")
	test.AssertSameString(t, "a_b_c_d_e_f_g_h_i_j", result)
}
