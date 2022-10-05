package download

import (
	"github.com/spf13/cobra"
)

var (
	update bool

	DownloadCmd = &cobra.Command{
		Use:   "download",
		Short: "Groups all download commands",
		Long:  "The download command groups all sub commands to download Rebrickable data for offline usage.",

		DisableFlagsInUseLine: true,
	}
)

func init() {
	DownloadCmd.AddCommand(updateCmd)
	DownloadCmd.AddCommand(imagesCmd)
}
