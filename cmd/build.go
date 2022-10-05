package cmd

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/spf13/cobra"
	"github.com/wendehals/bricks/api"
	"github.com/wendehals/bricks/build"
	"github.com/wendehals/bricks/cmd/options"
	"github.com/wendehals/bricks/model"
)

const (
	NEEDED_PARTS_FILE_OPT   = "parts"
	NEEDED_PARTS_FILE_SOPT  = "p"
	NEEDED_PARTS_FILE_ARG   = "-" + NEEDED_PARTS_FILE_SOPT + " NEEDED_PARTS_FILE"
	NEEDED_PARTS_FILE_USAGE = "a parts file containing the parts needed for the build"
)

var (
	setNum          string
	neededPartsFile string

	buildCmd = &cobra.Command{
		Use:   fmt.Sprintf("build PROVIDED_PARTS_FILE { %s | %s } %s %s", options.SET_NUM_ARG, NEEDED_PARTS_FILE_ARG, options.CREDENTIALS_ARG, options.OUTPUT_DIR_ARG),
		Short: "Calculates needed parts from a given collection to build a set",
		Long:  "The command build calculates which parts from a given collection are needed to build the given set.",

		DisableFlagsInUseLine: true,

		Args: cobra.ExactArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			var err error
			credentials, err = api.ImportCredentials(credentialsFile)
			if err != nil {
				return err
			}
			err = checkOptionsBuild()
			return err
		},
		Run: func(cmd *cobra.Command, args []string) {
			executeBuild(args)
		},
	}
)

func checkOptionsBuild() error {
	if (setNum == "" && neededPartsFile == "") || (setNum != "" && neededPartsFile != "") {
		return fmt.Errorf("please provide either %s or %s", options.SET_NUM_ARG, NEEDED_PARTS_FILE_ARG)
	}
	return nil
}

func init() {
	buildCmd.Flags().StringVarP(&setNum, options.SET_NUM_OPT, options.SET_NUM_SOPT, "", options.SET_NUM_USAGE)
	buildCmd.Flags().StringVarP(&neededPartsFile, NEEDED_PARTS_FILE_OPT, NEEDED_PARTS_FILE_SOPT, "", NEEDED_PARTS_FILE_USAGE)
	buildCmd.Flags().StringVarP(&credentialsFile, options.CREDENTIALS_OPT, options.CREDENTIALS_SOPT, "credentials.json", options.CREDENTIALS_USAGE)
	buildCmd.Flags().StringVarP(&outputDir, options.OUTPUT_DIR_OPT, options.OUTPUT_DIR_SOPT, "", options.OUTPUT_DIR_USAGE)
}

func executeBuild(args []string) {
	client := &http.Client{
		Timeout: time.Second * 5,
	}
	bricksAPI := api.NewBricksAPI(client, credentials.APIKey, options.Verbose)

	var neededCollection *model.Collection
	if setNum != "" {
		log.Printf("Building set '%s' from provided parts '%s'", setNum, args[0])
		neededCollection = api.RetrieveSetParts(bricksAPI, setNum, false)
		if outputDir == "" {
			outputDir = setNum
		}
	} else {
		log.Printf("Building '%s' from provided parts '%s'", neededPartsFile, args[0])
		neededCollection = model.Load(&model.Collection{}, neededPartsFile)
		if outputDir == "" {
			arg := []string{}
			arg = append(arg, neededPartsFile)
			outputDir = options.FileNameFromArgs(arg, "")
		}
	}

	providedCollection := model.Load(&model.Collection{}, args[0])

	build.Build(neededCollection, providedCollection, outputDir, options.Verbose)
}
