package collection

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/wendehals/bricks/cmd/options"
	"github.com/wendehals/bricks/model"
)

var (
	mergeCmd = &cobra.Command{
		Use:   fmt.Sprintf("merge PARTS_FILE %s", options.OUTPUT_FILE_ARG),
		Short: "Merges the parts of the same shape regardless of their color",
		Long:  "The command merges all parts of the same shape regardless of their color.",

		DisableFlagsInUseLine: true,

		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			executeMerge(args)
		},
	}
)

func executeMerge(args []string) {
	log.Print("Merging the parts of the same shape regardless of their color")

	collection := model.Load(&model.Collection{}, args[0])

	merged := collection.MergeByColor()

	if outputFile == "" {
		outputFile = options.FileNameFromArgs(args, "_merged.parts")
	}

	model.Save(merged, outputFile)
}
