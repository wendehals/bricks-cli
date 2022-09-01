package cmd

import (
	"github.com/spf13/cobra"
	"github.com/wendehals/bricks/model"
)

var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Downloads Rebrickable data for offline usage",
	Long:  "Downloads Rebrickable data for offline usage.",

	DisableFlagsInUseLine: true,

	Run: func(cmd *cobra.Command, args []string) {
		executeDownload()
	},
}

func executeDownload() {
	partRelationships := model.ImportPartRelationships()
	model.Save(partRelationships, "partRelationships.json")
}
