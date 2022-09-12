package download

import (
	"os"

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
	partRelationshipsPath := model.GetPartRelationshipsPath()
	if utils.FileExists(partRelationshipsPath) && !update {
		return
	}

	csvFile := model.DownloadPartRelationships()
	partRelationships := model.ConvertPartRelationships(csvFile)
	model.Save(partRelationships, partRelationshipsPath)
	os.Remove(csvFile)
}
