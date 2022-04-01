package cmd

import (
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/wendehals/bricks/cmd"
)

const (
	json_output_opt   = "output"
	json_output_sopt  = "o"
	json_output_arg   = "[-" + json_output_sopt + " JSON_FILE]"
	json_output_usage = "A name for the JSON output file"
)

var (
	jsonFile string

	collectionCmd = &cobra.Command{
		Use:   "collection",
		Short: "Groups all commands for working with bricks collections.",
		Long:  "The collection command groups all commands for working with bricks collections.",

		DisableFlagsInUseLine: true,
	}
)

func init() {
	cmd.RootCmd.AddCommand(collectionCmd)
}

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
