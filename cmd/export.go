package cmd

import (
	"net/http"
	"time"

	"github.com/spf13/cobra"
	"github.com/wendehals/bricks/api"
	"github.com/wendehals/bricks/model"
)

var (
	credentialsFile string
	htmlFile        string
	jsonFile        string

	credentials *api.Credentials

	exportCmd = &cobra.Command{
		Use:   "export [-c CREDENTIAL_FILE] [-html HTML_FILE] JSON_FILE",
		Short: "Exports the JSON input as HTML.",
		Long:  "The command exports the JSON input file as an HTML file.",

		DisableFlagsInUseLine: true,

		Args: cobra.ExactArgs(1),

		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			var err error
			credentials, err = api.ImportCredentials(credentialsFile)
			return err
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return executeExport(args)
		},
	}
)

func init() {
	exportCmd.Flags().StringVar(&htmlFile, "html", "", "A name for the HTML output file (default is name of JSON input file)")
}

func executeExport(args []string) error {
	collection, err := model.ImportCollection(args[0])
	if err != nil {
		return err
	}

	client := http.Client{
		Timeout: time.Second * 5,
	}

	bricksAPI := api.NewBricksAPI(&client, credentials.APIKey)
	err = bricksAPI.ReplaceImagesByMatchingColor(collection)
	if err != nil {
		return err
	}

	if htmlFile == "" {
		htmlFile = "export.html"
	}

	err = collection.ExportToHTML(htmlFile)
	if err != nil {
		return err
	}

	return nil
}
