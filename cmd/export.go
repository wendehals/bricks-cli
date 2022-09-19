package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/wendehals/bricks/api"
	"github.com/wendehals/bricks/cmd/options"
	"github.com/wendehals/bricks/model"
)

var (
	exportCmd = &cobra.Command{
		Use:   fmt.Sprintf("export %s %s PARTS_FILE", options.CREDENTIALS_ARG, options.OUTPUT_DIR_ARG),
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
	exportCmd.Flags().StringVarP(&credentialsFile, options.CREDENTIALS_OPT, options.CREDENTIALS_SOPT,
		"credentials.json", options.CREDENTIALS_USAGE)
	exportCmd.Flags().StringVarP(&outputDir, options.OUTPUT_DIR_OPT, options.OUTPUT_DIR_SOPT,
		"", options.OUTPUT_DIR_USAGE)
}

func executeExport(args []string) {
	if outputDir == "" {
		outputDir = options.FileNameFromArgs(args, "")
	}

	log.Printf("Exporting '%s' to directory '%s'", args[0], outputDir)

	model.Load(&model.Collection{}, args[0]).ExportToHTML(outputDir)
}
