package collection

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wendehals/bricks/cmd/options"
	"github.com/wendehals/bricks/services"
)

var (
	importCmd = &cobra.Command{
		Use:   fmt.Sprintf("import CSV_FILE %s", options.OUTPUT_FILE_ARG),
		Short: "Imports parts from a Rebrickable CSV file",
		Long:  "The command imports a collection of parts from a Rebrickable CSV file.",

		DisableFlagsInUseLine: true,

		Args: cobra.ExactArgs(1),

		Run: func(cmd *cobra.Command, args []string) {
			executeImport(args)
		},
	}
)

func executeImport(args []string) {
	if outputFile == "" {
		outputFile = options.FileNameFromArgs(args, ".parts")
	}

	collection := services.ImportCSVCollection(args[0])
	collection.Save(outputFile)
}
