package options

import (
	"fmt"
	"path/filepath"
	"regexp"
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

func FileNameFrom(s string) (string, error) {
	var illegalFileNames = [...]string{"CON", "PRN", "AUX", "NUL", "COM1", "COM2", "COM3", "COM4",
		"COM5", "COM6", "COM7", "COM8", "COM9", "LPT1", "LPT2", "LPT3", "LPT4", "LPT5",
		"LPT6", "LPT7", "LPT8", "LPT9"}

	for i := range illegalFileNames {
		if s == illegalFileNames[i] {
			return "", fmt.Errorf("the string %s is an illegal file name", s)
		}
	}

	illegalChars := regexp.MustCompile(`[<>:\"/\\|?*]`)
	return illegalChars.ReplaceAllLiteralString(s, "_"), nil
}
