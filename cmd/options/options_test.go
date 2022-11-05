package options

import (
	"testing"

	"github.com/wendehals/bricks/test"
)

func Test_FileNameFromArgs(t *testing.T) {
	result := FileNameFromArgs([]string{"arg1", "arg2", "arg3"}, ".test")
	test.AssertSameString(t, "arg1_arg2_arg3.test", result)

	result = FileNameFromArgs([]string{"ar<g1", "arg2", "arg3", "arg4", "arg5", "arg6"}, ".test")
	test.AssertSameString(t, "ar_g1_arg2_arg3_arg4_arg5.test", result)
}

func Test_ReplaceIllegalCharsFromFileName(t *testing.T) {
	result := ReplaceIllegalCharsFromFileName("a<b>c:d\"e/f\\g|h?i*j")
	test.AssertSameString(t, "a_b_c_d_e_f_g_h_i_j", result)
}
