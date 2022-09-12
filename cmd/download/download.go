package download

import (
	"fmt"

	"github.com/spf13/cobra"
)

const (
	UPDATE_OPT     = "update"
	UPDATE_SOPT    = "u"
	UPDATE_ARG     = "[-" + UPDATE_SOPT + "]"
	UPDATE_DEFAULT = false
	UPDATE_USAGE   = "update (overwrite) already downloaded files"
)

var (
	update bool

	DownloadCmd = &cobra.Command{
		Use:   fmt.Sprintf("download %s", UPDATE_SOPT),
		Short: "Groups all download commands",
		Long:  "The download command groups all sub commands to download Rebrickable data for offline usage.",

		DisableFlagsInUseLine: true,
	}
)

func init() {
	DownloadCmd.AddCommand(relationshipsCmd)
	DownloadCmd.AddCommand(imagesCmd)

	DownloadCmd.PersistentFlags().BoolVarP(&update, UPDATE_OPT, UPDATE_SOPT, UPDATE_DEFAULT, UPDATE_USAGE)
}
