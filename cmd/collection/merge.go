package collection

import (
	"fmt"
	"log"
	"regexp"

	"github.com/spf13/cobra"
	"github.com/wendehals/bricks/cmd/options"
	"github.com/wendehals/bricks/model"
	"github.com/wendehals/bricks/services"
)

var (
	mode string

	mergeCmd = &cobra.Command{
		Use:   fmt.Sprintf("merge PARTS_FILE [%s] %s", options.MODE_ARG, options.OUTPUT_FILE_ARG),
		Short: "Merges parts based on color, prints, molds, or alternates",
		Long:  "The command merges all parts by regards of their shape or color.",

		DisableFlagsInUseLine: true,

		Args: cobra.ExactArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return checkOptionsMerge()
		},
		Run: func(cmd *cobra.Command, args []string) {
			executeMerge(args)
		},
	}
)

func checkOptionsMerge() error {
	_, err := regexp.MatchString(`[camp]*`, mode)
	if err != nil {
		return fmt.Errorf("please provide only a combination of c(olor), a(lternatives), m(olds), and p(rints) for %s", options.MODE_ARG)
	}
	return nil
}

func init() {
	mergeCmd.Flags().StringVarP(&mode, options.MODE_OPT, options.MODE_SOPT, "c", options.MODE_USAGE)
}

func executeMerge(args []string) {
	log.Print("Merging the parts of the collection")

	collection := model.Load[model.Collection](args[0])

	services.Merge(&collection, services.ModeToUInt8(mode))

	if outputFile == "" {
		outputFile = options.FileNameFromArgs(args, "_merged.parts")
	}

	collection.Save(outputFile)
}
