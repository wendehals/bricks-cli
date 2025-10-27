package collection

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/wendehals/bricks-cli/cmd/options"
	"github.com/wendehals/bricks-cli/model"
)

var subtractCmd = &cobra.Command{
	Use:   fmt.Sprintf("subtract PARTS_FILE_MINUEND PARTS_FILE_SUBTRAHEND %s", options.OUTPUT_FILE_ARG),
	Short: "Subtracts one collection of parts from another",
	Long: `
The command subtracts the second collection of parts given in the args from the
first collection.

If the quantity of a certain part is zero the lot of this part is removed from
the resulting collection. If the quantity is negative the lot will be added to
the list of missing parts.`,

	DisableFlagsInUseLine: true,

	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		executeSubtract(args)
	},
}

func executeSubtract(args []string) {
	log.Printf("Subtracting the parts of '%s' from the parts of '%s'", args[1], args[0])

	minuend := model.Load[model.Collection](args[0])
	subtrahend := model.Load[model.Collection](args[1])
	result := minuend.Subtract(&subtrahend)

	if outputFile == "" {
		outputFile = options.FileNameFromArgs(args, "_subtracted.parts")
	}

	result.Save(outputFile)
}
