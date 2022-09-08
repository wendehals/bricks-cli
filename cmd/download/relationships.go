package download

import (
	"github.com/spf13/cobra"
	"github.com/wendehals/bricks/model"
)

var relationshipsCmd = &cobra.Command{
	Use:   "relationships",
	Short: "Downloads part relationships",
	Long:  "The relationships command downloads a file of part relationships concerning alternatives, molds, and prints.",

	DisableFlagsInUseLine: true,

	Run: func(cmd *cobra.Command, args []string) {
		executeRelationshipsDownload()
	},
}

func executeRelationshipsDownload() {
	partRelationships := model.ImportPartRelationships()
	model.Save(partRelationships, "partRelationships.json")
}
