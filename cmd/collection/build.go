package collection

import (
	"log"
	"net/http"
	"time"

	"github.com/spf13/cobra"
	"github.com/wendehals/bricks/api"
	"github.com/wendehals/bricks/cmd/options"
	"github.com/wendehals/bricks/model"
	"github.com/wendehals/bricks/utils"
)

var (
	credentialsFile string

	credentials *api.Credentials

	buildCmd = &cobra.Command{
		Use:   "build SET_NUM PARTS_FILE",
		Short: "Calculates needed parts from a given collection to build a set",
		Long:  "The command build calculates which parts from a given collection are needed to build the given set.",

		DisableFlagsInUseLine: true,

		Args: cobra.ExactArgs(2),
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			var err error
			credentials, err = api.ImportCredentials(credentialsFile)
			return err
		},
		Run: func(cmd *cobra.Command, args []string) {
			executeBuild(args)
		},
	}
)

func init() {
	buildCmd.Flags().StringVarP(&credentialsFile, options.CREDENTIALS_OPT, options.CREDENTIALS_SOPT, "credentials.json", options.CREDENTIALS_USAGE)
}

func executeBuild(args []string) {
	log.Printf("Building set '%s' from collection '%s'", args[0], args[1])

	client := &http.Client{
		Timeout: time.Second * 5,
	}
	bricksAPI := api.NewBricksAPI(client, credentials.APIKey, options.Verbose)

	neededCollection := api.RetrieveSetParts(bricksAPI, args[0], false)
	providedCollection := model.Load(&model.Collection{}, args[1])
	colors := bricksAPI.GetColors()
	partRelationships := model.Load(&model.PartRelationships{}, utils.GetPartRelationshipsPath())

	buildCollection := neededCollection.Build(providedCollection, &colors, partRelationships)

	neededCollection.RemoveQuantityZero()
	model.Save(neededCollection, "needed.parts")

	neededCollection.RemoveQuantityZero()
	model.Save(providedCollection, "remaining.parts")

	partEntries := []model.PartEntry{}
	for _, partEntry := range buildCollection.Parts {
		partEntries = append(partEntries, partEntry...)
	}

	resultCollection := model.Collection{}
	resultCollection.Parts = partEntries

	if outputFile == "" {
		outputFile = args[0]
	}
	resultCollection.ExportToHTML(outputFile)
}
