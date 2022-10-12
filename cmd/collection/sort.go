package collection

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/wendehals/bricks/cmd/options"
	"github.com/wendehals/bricks/model"
)

const (
	DESCENDING_OPT   = "descending"
	DESCENDING_SOPT  = "d"
	DESCENDING_ARG   = "[-" + DESCENDING_SOPT + "]"
	DESCENDING_USAGE = "sort in descending order"

	QUANTITY_OPT   = "quantity"
	QUANTITY_SOPT  = "q"
	QUANTITY_ARG   = "[-" + QUANTITY_SOPT + "]"
	QUANTITY_USAGE = "sort by quantity and name"
)

var (
	descending     bool
	sortByQuantity bool

	sortCmd = &cobra.Command{
		Use:   fmt.Sprintf("sort PARTS_FILE %s %s %s", options.OUTPUT_FILE_ARG, QUANTITY_ARG, DESCENDING_ARG),
		Short: "Sorts the parts of a collection",
		Long: `The command sorts the parts of a collection by default in ascending order by their color and name.
It could also sort the parts by their quantity and name.`,

		DisableFlagsInUseLine: true,

		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			executeSort(args)
		},
	}
)

func init() {
	sortCmd.Flags().BoolVarP(&sortByQuantity, QUANTITY_OPT, QUANTITY_SOPT, false, QUANTITY_USAGE)
	sortCmd.Flags().BoolVarP(&descending, DESCENDING_OPT, DESCENDING_SOPT, false, DESCENDING_USAGE)
}

func executeSort(args []string) {
	order := "ascending"
	if descending {
		order = "descending"
	}
	criterion := "color"
	if sortByQuantity {
		criterion = "quantity"
	}

	log.Printf("Sorting the parts of the given collection in %s order by their %s and name.", order, criterion)

	collection := model.Load(&model.Collection{}, args[0])

	if outputFile == "" {
		outputFile = options.FileNameFromArgs(args, "_sorted.parts")
	}

	if sortByQuantity {
		collection.SortByQuantityAndName(descending)
	} else {
		collection.SortByColorAndName(descending)
	}

	model.Save(collection, outputFile)
}
