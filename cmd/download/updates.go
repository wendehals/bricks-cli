package download

import (
	"github.com/spf13/cobra"
	"github.com/wendehals/bricks-cli/services"
)

var updateCmd = &cobra.Command{
	Use:   "updates",
	Short: "Updates locally cached data",
	Long: `The updates command downloads various files from the rebrickable.com server to update
locally cached data. This data comprises the part relationships based on
alternates, molds, and prints, as well as a colors and parts database.`,

	DisableFlagsInUseLine: true,

	Run: func(cmd *cobra.Command, args []string) {
		executeUpdatesDownload()
	},
}

func executeUpdatesDownload() {
	services.GetAlternates(true) // also updates Molds and Prints
	services.GetColors(true)
	services.GetShapes(true)
}
