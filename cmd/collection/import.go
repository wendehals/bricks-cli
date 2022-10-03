package collection

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/wendehals/bricks/cmd/options"
	"github.com/wendehals/bricks/model"
)

var (
	importCmd = &cobra.Command{
		Use:   fmt.Sprintf("import CSV_FILE %s", options.OUTPUT_FILE_ARG),
		Short: "Import parts from a CSV file",
		Long:  "The command imports a collection of parts from a CSV file.",

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

	log.Printf("Importing parts from '%s'", args[0])

	collection := model.ImportCSVCollection(args[0])

	model.Save(collection, outputFile)
}
