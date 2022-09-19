package options

import (
	"path/filepath"
	"regexp"
	"strings"
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
)

var (
	Verbose bool
)

func FileNameFromArgs(args []string, suffix string) string {
	var builder strings.Builder

	for i := 0; i < len(args) && i < 5; i++ {
		base := filepath.Base(args[i])
		ext := filepath.Ext(base)
		baseWoExt := base[:len(base)-len(ext)]

		index := strings.Index(baseWoExt, "_")
		if index != -1 {
			builder.WriteString(baseWoExt[:index])
		} else {
			builder.WriteString(baseWoExt)
		}

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
