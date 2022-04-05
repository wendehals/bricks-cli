package options

import (
	"path/filepath"
	"strings"
)

const (
	CREDENTIALS_OPT   = "credentials"
	CREDENTIALS_SOPT  = "c"
	CREDENTIALS_ARG   = "[-" + CREDENTIALS_SOPT + " CREDENTIAL_FILE]"
	CREDENTIALS_USAGE = "A JSON file containing the Rebrickable credentials"

	JSON_OUTPUT_OPT   = "json"
	JSON_OUTPUT_SOPT  = "o"
	JSON_OUTPUT_ARG   = "[-" + JSON_OUTPUT_SOPT + " JSON_FILE]"
	JSON_OUTPUT_USAGE = "A name for the JSON output file"

	HTML_OUTPUT_OPT   = "html"
	HTML_OUTPUT_SOPT  = "o"
	HTML_OUTPUT_ARG   = "[-" + HTML_OUTPUT_SOPT + " HTML_FILE]"
	HTML_OUTPUT_USAGE = "A name for the HTML output file"
)

func FileNameFromArgs(args []string, suffix string) string {
	var builder strings.Builder

	for i := 0; i < len(args) && i < 5; i++ {
		base := filepath.Base(args[i])
		ext := filepath.Ext(base)
		baseWoExt := base[:len(base)-len(ext)]

		builder.WriteString(baseWoExt[:strings.Index(baseWoExt, "_")])

		if i < len(args)-1 && i < 4 {
			builder.WriteString("_")
		}
	}

	builder.WriteString(suffix)

	return builder.String()
}
