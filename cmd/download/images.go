package download

import (
	"github.com/spf13/cobra"
	"github.com/wendehals/bricks/model"
)

var imagesCmd = &cobra.Command{
	Use:   "images",
	Short: "Downloads part images",
	Long: `The images command downloads images for all available part/color combinations.
	 
Be aware that this download has an amount of more than 5GB.`,

	DisableFlagsInUseLine: true,

	Run: func(cmd *cobra.Command, args []string) {
		executeImagesDownload()
	},
}

func executeImagesDownload() {
	model.DownloadPartImages()
}
