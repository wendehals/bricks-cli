package collection

import (
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

const (
	json_output_opt   = "output"
	json_output_sopt  = "o"
	json_output_arg   = "[-" + json_output_sopt + " JSON_FILE]"
	json_output_usage = "A name for the JSON output file"
)

var (
	jsonFile string

	CollectionCmd = &cobra.Command{
		Use:   "collection",
		Short: "Groups all commands for working with bricks collections.",
		Long:  "The collection command groups all commands for working with bricks collections.",

		DisableFlagsInUseLine: true,
	}
)

func init() {
	CollectionCmd.AddCommand(exportCmd)
	CollectionCmd.AddCommand(maxCmd)
	CollectionCmd.AddCommand(mergeCmd)
	CollectionCmd.AddCommand(sortCmd)
	CollectionCmd.AddCommand(subtractCmd)
	CollectionCmd.AddCommand(sumCmd)
}

func fileNameFromArgs(args []string, suffix string) string {
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
