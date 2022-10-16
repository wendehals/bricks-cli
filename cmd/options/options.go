package options

import (
	"regexp"
	"strings"

	"github.com/wendehals/bricks/utils"
)

const (
	CREDENTIALS_OPT   = "credentials"
	CREDENTIALS_SOPT  = "c"
	CREDENTIALS_ARG   = "[-" + CREDENTIALS_SOPT + " CREDENTIAL_FILE]"
	CREDENTIALS_USAGE = "a JSON file containing the Rebrickable credentials"

	OUTPUT_FILE_OPT   = "output"
	OUTPUT_FILE_SOPT  = "o"
	OUTPUT_FILE_ARG   = "[-" + OUTPUT_FILE_SOPT + " OUTPUT_FILE]"
	OUTPUT_FILE_USAGE = "a name for the output file"

	OUTPUT_DIR_OPT   = "dir"
	OUTPUT_DIR_SOPT  = "d"
	OUTPUT_DIR_ARG   = "[-" + OUTPUT_DIR_SOPT + " OUTPUT_DIR]"
	OUTPUT_DIR_USAGE = "a name for the output directory"

	SET_NUM_OPT   = "set"
	SET_NUM_SOPT  = "s"
	SET_NUM_ARG   = "-" + SET_NUM_SOPT + " SET_NUMBER"
	SET_NUM_USAGE = "the set number"

	MODE_OPT   = "mode"
	MODE_SOPT  = "m"
	MODE_ARG   = "-" + MODE_SOPT + " MODES"
	MODE_USAGE = "a combination of c(olor), a(lternatives), m(olds), and p(rints) for calculating usable parts"
)

var Verbose bool

func FileNameFromArgs(args []string, suffix string) string {
	var builder strings.Builder

	for i := 0; i < len(args) && i < 5; i++ {
		base, _ := utils.SplitFileName(args[i])
		builder.WriteString(base)

		if i < len(args)-1 && i < 4 {
			builder.WriteString("_")
		}
	}

	builder.WriteString(suffix)

	return ReplaceIllegalCharsFromFileName(builder.String())
}

func ReplaceIllegalCharsFromFileName(s string) string {
	return regexp.MustCompile(`[<>:\"/\\|?*]`).ReplaceAllLiteralString(s, "_")
}
