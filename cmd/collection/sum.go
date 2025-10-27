package collection

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/wendehals/bricks-cli/cmd/options"
	"github.com/wendehals/bricks-cli/model"
)

var sumCmd = &cobra.Command{
	Use:   fmt.Sprintf("sum %s PARTS_FILE1 PARTS_FILE2 ...", options.OUTPUT_FILE_ARG),
	Short: "Sums up the parts of multiple collections",
	Long: `
The command sums up all parts of multiple collections to a new collection by
merging identical parts to single lots.`,

	DisableFlagsInUseLine: true,

	Args: cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		executeSum(args)
	},
}

func executeSum(args []string) {
	log.Print("Summing up the parts of the given collections")

	var collections []model.Collection
	for _, filename := range args {
		collections = append(collections, model.Load[model.Collection](filename))
	}

	sum := model.NewCollection()
	for _, collection := range collections {
		sum.Add(&collection)
	}

	if outputFile == "" {
		outputFile = options.FileNameFromArgs(args, "_sum.parts")
	}

	sum.Save(outputFile)
}
