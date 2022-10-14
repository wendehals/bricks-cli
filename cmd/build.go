package cmd

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/wendehals/bricks/api"
	"github.com/wendehals/bricks/cmd/options"
	"github.com/wendehals/bricks/model"
	"github.com/wendehals/bricks/services"
)

const (
	NEEDED_PARTS_FILE_OPT   = "parts"
	NEEDED_PARTS_FILE_SOPT  = "p"
	NEEDED_PARTS_FILE_ARG   = "-" + NEEDED_PARTS_FILE_SOPT + " NEEDED_PARTS_FILE"
	NEEDED_PARTS_FILE_USAGE = "a parts file containing the parts needed for the build"

	BUILD_MODE_OPT   = "mode"
	BUILD_MODE_SOPT  = "m"
	BUILD_MODE_ARG   = "-" + BUILD_MODE_SOPT + " MODES"
	BUILD_MODE_USAGE = "a combination of c(olor), a(lternatives), m(olds), and p(rints) for calculating usable parts"
)

var (
	setNum          string
	buildMode       string
	neededPartsFile string

	buildCmd = &cobra.Command{
		Use: fmt.Sprintf("build PROVIDED_PARTS_FILE { %s | %s } %s %s %s", options.SET_NUM_ARG, NEEDED_PARTS_FILE_ARG,
			BUILD_MODE_ARG, options.CREDENTIALS_ARG, options.OUTPUT_DIR_ARG),
		Short: "Calculates needed parts from a given collection to build a set",
		Long: fmt.Sprintf(`The command build calculates which parts from a given collection are needed to build the given set.

By default, it considers only parts of the same type and color. If also parts with a different color, alternative parts,
molds, or prints should be considered, use the '%s' option to provide any combination of those.`, BUILD_MODE_ARG),

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

func init() {
	buildCmd.Flags().StringVarP(&setNum, options.SET_NUM_OPT, options.SET_NUM_SOPT, "", options.SET_NUM_USAGE)
	buildCmd.Flags().StringVarP(&neededPartsFile, NEEDED_PARTS_FILE_OPT, NEEDED_PARTS_FILE_SOPT, "", NEEDED_PARTS_FILE_USAGE)
	buildCmd.Flags().StringVarP(&buildMode, BUILD_MODE_OPT, BUILD_MODE_SOPT, "", BUILD_MODE_USAGE)
	buildCmd.Flags().StringVarP(&credentialsFile, options.CREDENTIALS_OPT, options.CREDENTIALS_SOPT, "credentials.json", options.CREDENTIALS_USAGE)
	buildCmd.Flags().StringVarP(&outputDir, options.OUTPUT_DIR_OPT, options.OUTPUT_DIR_SOPT, "", options.OUTPUT_DIR_USAGE)
}

func checkOptionsBuild() error {
	if (setNum == "" && neededPartsFile == "") || (setNum != "" && neededPartsFile != "") {
		return fmt.Errorf("please provide either %s or %s", options.SET_NUM_ARG, NEEDED_PARTS_FILE_ARG)
	}
	_, err := regexp.MatchString(`[camp]*`, buildMode)
	if err != nil {
		return fmt.Errorf("please provide only a combination of c(olor), a(lternatives), m(olds), and p(rints) for %s", BUILD_MODE_ARG)
	}
	return nil
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

	var mode uint8 = 0
	if strings.Contains(buildMode, "c") {
		mode = mode ^ services.MODE_COLOR
	}
	if strings.Contains(buildMode, "a") {
		mode = mode ^ services.MODE_ALTERNATES
	}
	if strings.Contains(buildMode, "m") {
		mode = mode ^ services.MODE_MOLDS
	}
	if strings.Contains(buildMode, "p") {
		mode = mode ^ services.MODE_PRINTS
	}

	buildCollection := services.Build(neededCollection, providedCollection, mode)
	model.Save(buildCollection, fmt.Sprintf("%s/result.build", outputDir))
	services.ExportBuildCollectionToHTML(buildCollection, outputDir, "build")

	model.Save(providedCollection, fmt.Sprintf("%s/remaining.parts", outputDir))
	services.ExportCollectionToHTML(providedCollection, outputDir, "remaining")
}
