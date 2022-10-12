package download

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wendehals/bricks/services"
)

const (
	UPDATE_OPT     = "update"
	UPDATE_SOPT    = "u"
	UPDATE_ARG     = "[-" + UPDATE_SOPT + "]"
	UPDATE_DEFAULT = false
	UPDATE_USAGE   = "update (overwrite) already downloaded files"
)

var imagesCmd = &cobra.Command{
	Use:   fmt.Sprintf("images %s", UPDATE_SOPT),
	Short: "Downloads part images",
	Long: `The images command downloads images for all available part/color combinations.
	 
Be aware that this download has an amount of more than 5GB.`,

	DisableFlagsInUseLine: true,

	Run: func(cmd *cobra.Command, args []string) {
		executeImagesDownload()
	},
}

func init() {
	imagesCmd.Flags().BoolVarP(&update, UPDATE_OPT, UPDATE_SOPT, UPDATE_DEFAULT, UPDATE_USAGE)
}

func executeImagesDownload() {
	services.DownloadPartImages(update)
}
