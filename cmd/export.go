package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/wendehals/bricks/api"
	"github.com/wendehals/bricks/cmd/options"
	"github.com/wendehals/bricks/model"
	"github.com/wendehals/bricks/services"
	"github.com/wendehals/bricks/utils"
)

const (
	SORT_OPT   = "sort"
	SORT_SOPT  = "s"
	SORT_ARG   = "[-" + SORT_SOPT + "]"
	SORT_USAGE = "sort the parts in ascending order by color and name before exporting"
)

var (
	sort bool

	exportCmd = &cobra.Command{
		Use:   fmt.Sprintf("export %s %s %s PARTS_FILE", SORT_ARG, options.CREDENTIALS_ARG, options.OUTPUT_DIR_ARG),
		Short: "Exports the parts input file as HTML",
		Long:  "The command exports the parts input file as an HTML file to the given output directory.",

		DisableFlagsInUseLine: true,

		Args: cobra.ExactArgs(1),

		PreRunE: func(cmd *cobra.Command, args []string) error {
			var err error
			credentials, err = api.ImportCredentials(credentialsFile)
			return err
		},
		Run: func(cmd *cobra.Command, args []string) {
			executeExport(args)
		},
	}
)

func init() {
	exportCmd.Flags().BoolVarP(&sort, SORT_OPT, SORT_SOPT, false, SORT_USAGE)
	exportCmd.Flags().StringVarP(&credentialsFile, options.CREDENTIALS_OPT, options.CREDENTIALS_SOPT,
		utils.CredentialsDefaultPath(), options.CREDENTIALS_USAGE)
	exportCmd.Flags().StringVarP(&outputDir, options.OUTPUT_DIR_OPT, options.OUTPUT_DIR_SOPT,
		"", options.OUTPUT_DIR_USAGE)
}

func executeExport(args []string) {
	if outputDir == "" {
		outputDir = options.FileNameFromArgs(args, "")
	}

	log.Printf("Exporting '%s' to directory '%s'", args[0], outputDir)

	collection := model.Load(model.NewCollection(), args[0])
	if sort {
		collection.SortByColorAndName(false)
	}

	services.ExportCollectionToHTML(collection, outputDir, args[0])
}
