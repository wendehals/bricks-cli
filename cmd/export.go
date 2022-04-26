package cmd

import (
	"fmt"
	"net/http"
	"time"

	"github.com/spf13/cobra"
	"github.com/wendehals/bricks/api"
	"github.com/wendehals/bricks/cmd/options"
	"github.com/wendehals/bricks/model"
)

var (
	credentialsFile string
	htmlFile        string

	credentials *api.Credentials

	exportCmd = &cobra.Command{
		Use:   fmt.Sprintf("export %s [-o HTML_FILE] JSON_FILE", options.CREDENTIALS_ARG),
		Short: "Exports the JSON input as HTML.",
		Long:  "The command exports the JSON input file as an HTML file.",

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
	exportCmd.Flags().StringVarP(&htmlFile, options.HTML_OUTPUT_OPT, options.HTML_OUTPUT_SOPT,
		"", options.HTML_OUTPUT_USAGE)
}

func executeExport(args []string) {
	collection := model.ImportCollection(args[0])

	client := http.Client{
		Timeout: time.Second * 5,
	}

	bricksAPI := api.NewBricksAPI(&client, credentials.APIKey)
	bricksAPI.ReplaceImagesByMatchingColor(collection)

	if htmlFile == "" {
		htmlFile = options.FileNameFromArgs(args, ".html")
	}

	collection.ExportToHTML(htmlFile)
}
