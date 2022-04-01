package cmd

import (
	"path/filepath"
	"strings"

	"github.com/wendehals/bricks/api"
)

const (
	credentials_opt   = "credentials"
	credentials_sopt  = "c"
	credentials_arg   = "[-" + credentials_sopt + " CREDENTIAL_FILE]"
	credentials_usage = "A JSON file containing the Rebrickable credentials"

	json_output_opt   = "output"
	json_output_sopt  = "o"
	json_output_arg   = "[-" + json_output_sopt + " JSON_FILE]"
	json_output_usage = "A name for the JSON output file"
)

var (
	credentialsFile string
	credentials     *api.Credentials

	jsonFile string
)

func fileNameFromArgs(args []string, suffix string) string {
	var builder strings.Builder

	for i := 0; i < len(args) && i < 5; i++ {
		base := filepath.Base(args[i])
		ext := filepath.Ext(base)
		builder.WriteString(base[:len(base)-len(ext)])
	}

	builder.WriteString(suffix)

	return builder.String()
}
