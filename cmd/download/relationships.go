package download

import (
	"github.com/spf13/cobra"
	"github.com/wendehals/bricks/model"
	"github.com/wendehals/bricks/utils"
)

var relationshipsCmd = &cobra.Command{
	Use:   "relationships",
	Short: "Downloads part relationships",
	Long:  "The relationships command downloads a file of part relationships based on alternatives, molds, and prints.",

	DisableFlagsInUseLine: true,

	Run: func(cmd *cobra.Command, args []string) {
		executeRelationshipsDownload()
	},
}

func executeRelationshipsDownload() {
	partRelationshipsPath := utils.PartRelationshipsPath()
	if utils.FileExists(partRelationshipsPath) && !update {
		return
	}

	model.GetPartRelationships()
}
