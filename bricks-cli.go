package main

import (
	"os"

	"github.com/wendehals/bricks-cli/cmd"
)

func main() {
	err := cmd.RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
