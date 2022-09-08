package download

import (
	"github.com/spf13/cobra"
)

var DownloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Groups all download commands",
	Long:  "The download command groups all sub commands to download Rebrickable data for offline usage.",

	DisableFlagsInUseLine: true,
}

func init() {
	DownloadCmd.AddCommand(relationshipsCmd)
	DownloadCmd.AddCommand(imagesCmd)
}
