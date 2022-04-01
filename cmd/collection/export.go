package collection

import (
	"github.com/spf13/cobra"
	"github.com/wendehals/bricks/model"
)

var (
	htmlFile string

	exportCmd = &cobra.Command{
		Use:   "export [-html HTML_FILE] JSON_FILE",
		Short: "Exports the JSON input as HTML.",
		Long:  "The command exports the JSON input file as an HTML file.",

		DisableFlagsInUseLine: true,

		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return executeExport(args)
		},
	}
)

func init() {
	exportCmd.Flags().StringVar(&htmlFile, "html", "", "A name for the HTML output file (default is name of JSON input file)")
}

func executeExport(args []string) error {
	collection, err := model.Import(args[0])
	if err != nil {
		return err
	}

	if htmlFile == "" {
		htmlFile = fileNameFromArgs(args, ".html")
	}

	err = collection.ExportToHTML(htmlFile)
	if err != nil {
		return err
	}

	return nil
}
